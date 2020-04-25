package impl

import (
	"fmt"
	"socket"
)

type PingRouter struct {
	socket.IRouter
}

//预期处理方法
func (pr *PingRouter) PreHandle(request socket.IRequest){
	fmt.Println("Call PingRouter PreHandle...")
	_, err:= request.GetConnection().GetTCPConnection().Write([]byte("before ping... \n"))
	if err != nil {
		fmt.Printf("Call PingRouter PreHandle call back ping err=%s \n", err)
	}
}
//处理业务方法
func (pr *PingRouter) Handle(request socket.IRequest){
	fmt.Println("Call PingRouter Handle...")
	_, err:= request.GetConnection().GetTCPConnection().Write([]byte("ping... \n"))
	if err != nil {
		fmt.Printf("Call PingRouter Handle call back ping err=%s \n", err)
	}
}
//处理完成后钩子函数
func (pr *PingRouter) PostHandle(request socket.IRequest) {
	fmt.Println("Call PingRouter PostHandle...")
	_, err:= request.GetConnection().GetTCPConnection().Write([]byte("after ping... \n"))
	if err != nil {
		fmt.Printf("Call PingRouter PostHandle call back ping err=%s \n", err)
	}
}