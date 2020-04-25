package impl

//消息内容
type Message struct {
	Id      uint32 //消息ID
	DateLen uint32 //消息长度
	EventId uint32 //事件ID
	Data    []byte //消息内容
}
//消息ID
func (m *Message) GetMsgId() uint32 {
	return m.Id
}
//获取消息长度
func (m *Message) GetMsgLen() uint32 {
	return m.DateLen
}
//事件ID
func (m *Message) GetEventId() uint32 {
	return m.EventId
}
//获取消息内容
func (m *Message) GetData() []byte {
	return m.Data
}
//设置消息ID
func (m *Message) SetMsgId(msgId uint32) {
	m.Id = msgId
}
//设置事件
func (m *Message) SetEventId(eventId uint32) {
	m.EventId = eventId
}
//消息数据
func (m *Message) SetData(data []byte) {
	m.Data = data
}
//设置数据长度
func (m *Message) SetDataLen(dataLen uint32) {
	m.DateLen = dataLen
}