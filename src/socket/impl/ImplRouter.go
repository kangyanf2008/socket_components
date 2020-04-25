package impl

import "socket"

//实现router时，先
type BaseRouter struct {}

//预期处理方法
func (br *BaseRouter) PreHandle(request socket.IRequest){}
//处理业务方法
func (br *BaseRouter) Handle(request socket.IRequest){}
//处理完成后钩子函数
func (br *BaseRouter) PostHandle(request socket.IRequest) {}