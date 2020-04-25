package socket

/**
请求接口
 */

type IRequest interface {
	//得到当前连接
	GetConnection() IConnection

	//得到当前请求数据
	GetData() []byte

}