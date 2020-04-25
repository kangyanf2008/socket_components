package impl

import (
	"fmt"
	"io"
	"net"
	"socket"
	"sync"
	"time"
)

type Connection struct {
	//当前连接socket tcp套接字
	Conn *net.TCPConn

	//连接ID
	ConnId uint32

	//当前连接状态
	isClose bool

	//停止channel
	ExitChan chan bool

	//处理方法
	Router socket.IRouter

	RWLock sync.RWMutex

	//任务队列
	TaskQueue chan *Request
	//队列大小
	QueueSize int
}

//创建连接对象
func NewConnection(conn *net.TCPConn, connID uint32, router socket.IRouter, queueSize int) socket.IConnection {
	if queueSize <= 0 {
		queueSize = 100
	}
	c := &Connection{
		Conn:     conn,
		ConnId:   connID,
		isClose:  false,
		ExitChan: make(chan bool, 1),
		Router:   router,
		TaskQueue:make(chan *Request, queueSize),
	}
	return c
}

//读取连接数据
func (c *Connection) StartReader() {
	fmt.Printf("reader Goroutine is running...\n")
	defer fmt.Printf("Connid=%d exit \n", c.ConnId)
	defer c.Stop()
	//处理客户端请求任务
	go c.HandleTask()

	for ;!c.isClose;{
		//读取客户端heade信息
		dp := NewDataPack()
		headDataByte := make([]byte, dp.GetHeadLen())
		_, err :=io.ReadFull(c.Conn, headDataByte)
		if err != nil {
			fmt.Println("read head err: ",err)
			return
		}
		//解析头信息
		msgHead, err := dp.Unpack(headDataByte)
		if err != nil {
			fmt.Println("server unpack err: ",err)
			return
		}
		//读取客户body信息
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*Message)
			data := make([]byte, msg.GetMsgLen())
			msg.SetData(data)
			_, err := io.ReadFull(c.Conn, data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
		}
		req := Request{
			conn: c,
			data: msgHead.GetData(),
		}
		//接收消息放到消息队列
		c.TaskQueue <- &req
	}
}

//启动连接
func (c *Connection) Start() {
	fmt.Printf("conn start() ConnID=%d \n", c.ConnId)
	go c.StartReader()
}

//停止连接
func (c *Connection) Stop() {
	if c.isClose {
		return
	}
	c.ExitChan <- true //退出通知
	c.isClose = true   //修改关闭状态
	close(c.TaskQueue) //关闭任务队列
	//c.Conn.Close()     //关闭连接
	//close(c.ExitChan)  //关闭管道
}

//获以当前链接绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

//获取远程客户端TCP状态和IP port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

//发送数据
func (c *Connection) Send(data []byte) error {
	return nil
}

//发送数据
func (c *Connection) HandleTask()  {
	endCond := false
	defer c.Conn.Close()
	for ; !endCond || !c.isClose; {
		select {
		case exit := <-c.ExitChan:
			if exit {
				close(c.TaskQueue)
			}
		case t:= <-c.TaskQueue:
			endCond = false
			c.Router.PreHandle(t)
			c.Router.Handle(t)
			c.Router.PostHandle(t)
		//超过，并且连接已经关闭
		case <-time.After(time.Millisecond*20):
			println("end with timeout")
			endCond = true
		}
	}

}

