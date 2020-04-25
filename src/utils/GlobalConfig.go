package utils

type GlobalConfigStruct struct {
	TCPServerName   string
	Host            string //主机地址
	TcpPort         int    //端口号
	Name            string
	Version         string //版本大
	MaxConn         int    //最大连接数
	MaxPackageSize  uint32 //最大读取包大小 单位字节
	WorkerNum       int    //工作线程大小
	RouterQueueSize int    //路由队列大小
}

func (g GlobalConfigStruct) Reload() {
	GlobalConfig.Name = "myServer"
	GlobalConfig.Version = "0.1"
	GlobalConfig.TcpPort = 8888
	GlobalConfig.Host = "0.0.0.0"
	GlobalConfig.MaxConn = 1000
	GlobalConfig.MaxPackageSize = 512
	GlobalConfig.RouterQueueSize = 100
	GlobalConfig.WorkerNum = 5
}

var GlobalConfig *GlobalConfigStruct

func init() {
	GlobalConfig = &GlobalConfigStruct{
		Name:            "myServer",
		Version:         "0.1",
		TcpPort:         8888,
		Host:            "0.0.0.0",
		MaxConn:         1000,
		MaxPackageSize:  512,
		RouterQueueSize: 100,
		WorkerNum:       5,
	}
}
