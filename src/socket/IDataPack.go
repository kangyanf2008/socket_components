package socket

type IDataPack interface {
	//读取包头的长度
	GetHeadLen() uint32
	//封包
	Pack(msg IMessage) ([]byte, error)
	//解包
	Unpack([]byte) (IMessage, error)
}