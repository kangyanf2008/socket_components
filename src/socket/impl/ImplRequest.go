package impl

import "socket"

type Request struct {
	//客户端已经建立连接
	conn socket.IConnection

	//客户端请求数据
	data []byte
}

//得到当前连接
func ( r *Request) GetConnection() socket.IConnection {
	return 	r.conn
}

//得到当前请求数据
func ( r *Request) GetData() []byte {
	return r.data
}