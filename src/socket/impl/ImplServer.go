package impl

import (
	"fmt"
	"net"
	"socket"
	"utils"
)

type Server struct {
	ServerName string
	IpVersion string
	Address string
	Port int
	Router socket.IRouter
}

//初始化服务
func NewServer(serverName, address string, port int) socket.IServer {
	s := &Server{
		ServerName:serverName,
		Address:address,
		Port:port,
		IpVersion:"tcp4",
	}
	return s
}

//启动服务
func (s *Server) Start() {
	go func() {
		addr, err :=net.ResolveTCPAddr(s.IpVersion, fmt.Sprintf("%s:%d",s.Address, s.Port))
		if err != nil {
			fmt.Println("resolve tcp address error:", err)
			return
		}
		//进行监听
		listener, err := net.ListenTCP(s.IpVersion, addr)
		if err != nil {
			fmt.Println("listen ",s.IpVersion, "err ", err)
		}
		fmt.Println("start socket success", s.ServerName, "sucess, Listenning....")
		//监听客户端监听连接
		var cid uint32 = 0
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}
			dealConn := NewConnection(conn, cid, s.Router, utils.GlobalConfig.RouterQueueSize)
			cid++
			dealConn.Start()
		}
	}() //异常启动

}

func (s *Server) Stop() {

}

//启动服务
func (s *Server) Serve() {
	//启动服务
	s.Start()
	select {}
}

//添加路由方法
func (s *Server) AddRouter(router socket.IRouter) {
	s.Router = router
	fmt.Println("Add Router Success !")
}
