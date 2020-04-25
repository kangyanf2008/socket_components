package impl

import (
	"errors"
	"fmt"
	"net"
)

func CallBackToClient(conn *net.TCPConn, data []byte,  cnt int) error {
	fmt.Printf("[Conn Handle] CallBackToClient \n")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Printf("write buf to client err=%s \n", err)
		return errors.New("CallBackToClient error="+err.Error())
	}
	return nil
}
