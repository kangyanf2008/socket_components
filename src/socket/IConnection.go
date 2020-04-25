package socket

import "net"

type IConnection interface {
	//启动连接
	Start()

	//停止连接
	Stop()

	//获以当前链接绑定socket conn
	GetTCPConnection() *net.TCPConn

	//获取连接ID
	GetConnID() uint32

	//获取远程客户端TCP状态和IP port
	RemoteAddr() net.Addr

	//发送数据
	Send(data []byte) error
}
//定义一个处理业务方法
type HandleFun func(*net.TCPConn, []byte, int) error
