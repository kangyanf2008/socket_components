package main

import (
	"socket/impl"
	"utils"
)

func main() {
	server := impl.NewServer(utils.GlobalConfig.Name,utils.GlobalConfig.Host, utils.GlobalConfig.TcpPort)
	//添加router
	server.AddRouter(new(impl.PingRouter))
	server.Serve()
}
