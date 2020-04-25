package impl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"socket"
	"utils"
)

type DataPack struct{}

func NewDataPack() socket.IDataPack {
	return &DataPack{}
}

//读取包头的长度
func (d DataPack) GetHeadLen() uint32 {
	//DataLen uint32(4字节) - eventId uint32(4字节)  - Id uint32(4字节)
	return 12
}

//封包
func (d DataPack) Pack(msg socket.IMessage) ([]byte, error) {
	//定义消息buffer
	dataBuff := bytes.NewBuffer([]byte{})
	//先写数据长度
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	//事件ID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetEventId()); err != nil {
		return nil, err
	}
	//消息ID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}
	//消息
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil
}

//解包,只需要解析head
func (d DataPack) Unpack(byteData []byte) (socket.IMessage, error) {
	dataBuf := bytes.NewReader(byteData)
	//读取数据长度
	msg := &Message{}
	//解析数据长度
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.DateLen); err != nil {
		return nil, err
	}
	//解析eventID
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.EventId); err != nil {
		return nil, err
	}
	//解析消息ID
	if err := binary.Read(dataBuf, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}
	if utils.GlobalConfig.MaxPackageSize < msg.DateLen {
		return nil, errors.New("too large msg data recv")
	}
	return msg, nil
}
