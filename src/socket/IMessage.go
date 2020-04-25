package socket

/**
请求消息封闭message
*/

type IMessage interface {
	//消息ID
	GetMsgId() uint32
	//事件ID
	GetEventId() uint32
	//获取消息长度
	GetMsgLen() uint32
	//获取消息内容
	GetData() []byte
	//设置消息ID
	SetMsgId(uint32)
	//设置事件
	SetEventId(uint32)
	//消息数据
	SetData([]byte)
	//设置数据长度
	SetDataLen(uint32)
}
