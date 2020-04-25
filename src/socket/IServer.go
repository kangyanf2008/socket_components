package socket

type IServer interface {
	//启动服务
	Start()
	//停止服务
	Stop()
	//运行服务
	Serve()
	//添加路由功能
	AddRouter(router IRouter)

}

