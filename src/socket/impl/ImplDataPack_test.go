package impl

import (
	"fmt"
	"io"
	"net"
	"testing"
	"time"
)

func TestDataPack( t *testing.T) {

	//服务器连接
	listener, err :=net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("server listener err: ", err)
		return
	}
	go func() {
		for {
			conn, err := listener.Accept()
			fmt.Println("accept client connect")
			if err != nil {
				fmt.Println("server accept error:", err)
			}
			go func(conn net.Conn) {
				dp := NewDataPack()
				for {
					headDataByte := make([]byte, dp.GetHeadLen())
					//conn.Read(headDataByte)
					_, err :=io.ReadFull(conn, headDataByte)
					if err != nil {
						fmt.Println("read head err: ",err)
						return
					}
					msgHead, err := dp.Unpack(headDataByte)
					if err != nil {
						fmt.Println("server unpack err: ",err)
						return
					}
					if msgHead.GetMsgLen() > 0 {
						msg := msgHead.(*Message)
						data := make([]byte, msg.GetMsgLen())
						msg.Data = data
						_, err := io.ReadFull(conn, data)
						if err != nil {
							fmt.Println("server unpack data err:", err)
							return
						}
						fmt.Println("-->Recv msgID=", msg.Id, ", datalen=", msg.DateLen, ", eventId=", msg.EventId, ",data=", string(msg.Data))
					}
				}//end for

			}(conn)
		}
	}()

	time.Sleep(time.Second*5)
	//客户端连接
	conn, err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err:", err)
	}

	//组装请求数据
	dp := NewDataPack()
	data := []byte("abcdefghijklmnopqrstuvwxyz")
	msg1 := &Message{
		Id:1,
		DateLen:uint32(len(data)),
		EventId:2,
		Data:data,
	}
fmt.Println(msg1)
	//封装包1
	send1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("client pack msg1 error: ", err)
		return
	}

	//封装包2
	data2 := []byte("1234567890abcdefghijklmnopqrstuvwxyz")
	msg2 := &Message{
		Id:2,
		DateLen:uint32(len(data2)),
		EventId:3,
		Data:data2,
	}
fmt.Println(msg2)
	send2, err2 := dp.Pack(msg2)
	if err2 != nil {
		fmt.Println("client pack msg2 error: ", err2)
		return
	}
	send1 = append(send1, send2...)
	fmt.Println(send1)
	conn.Write(send1)
	select {}
}


