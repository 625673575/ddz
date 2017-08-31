package message

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"net"
	"github.com/alecthomas/log4go"
)
func SendAck(conn *net.Conn, ack proto.Message) {
	content, err := proto.Marshal(ack)
	if err == nil {
		(*conn).Write(content)
	} else {
		log4go.Error(err.Error())
	}
}
func PackageMessage(messageType REPLY_TYPE, message proto.Message) proto.Message {
	data, _ := proto.Marshal(message)
	mess := &ProtoMessage{MessageType: int32(messageType), Data: data}
	return mess
}

func SendPackageMessage(conn *net.Conn,messageType REPLY_TYPE, message proto.Message){
	SendAck(conn,PackageMessage(messageType,message))
}
type MessageConverter interface {
	Marshal() ([]byte, error)
	UnMarshal()
}

type JsonMessage struct {
	MessageType int32
	Data        interface{}
}

func (x *JsonMessage) Marshal() ([]byte, error) {
	return json.Marshal(*x)
}
func (x *JsonMessage) UnMarshal(data []byte) {
	json.Unmarshal(data, x)
}
