package socket

/**
路由抽像接口
 */
type IRouter interface {
	//预期处理方法
	PreHandle(request IRequest)
	//处理业务方法
	Handle(request IRequest)
	//处理完成后钩子函数
	PostHandle(request IRequest)
}