package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err :=net.Dial("tcp","127.0.0.1:8888")
	if err != nil {

	}
	for {
		_,err := conn.Write([]byte("hello 你好  ........"))
		if err != nil {
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error", err)
		}
		fmt.Printf("socket call back %s, cnt = %d \n", buf, cnt)
		time.Sleep(1)
	}
}
