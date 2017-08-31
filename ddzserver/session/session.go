package session

import (
	"net"
	"github.com/golang/protobuf/proto"
	"ddzserver/message"
	"github.com/alecthomas/log4go"
	"fmt"
)

const STATE_LOGIN_SUCCESS=0
const STATE_LOGIN_WRONG_PASSWORD=1
const STATE_LOGIN_NO_USER=2
const STATE_LOGIN_UNKNOW_ERROR=-1

var SessionMap map[int64]*PlayerSession = make(map[int64]*PlayerSession)

type PlayerSession struct {
	IP string
	Connect *net.Conn
	HasLogin bool
	Info message.UserInfo
	IsConnect bool
}
func GetSession(userId int64)*PlayerSession{
	return SessionMap[userId]
}
func NewSession(conn *net.Conn)  {
	sess := new(PlayerSession)
	sess.IP=(*conn).RemoteAddr().String()
	sess.Connect = conn
	go sess.handleClient()
}
func (sess *PlayerSession)handleClient() {
	defer (*sess.Connect).Close()
	var buf [512]byte
	go func() {
		for {
			n, err := (*sess.Connect).Read(buf[0:])
			if err != nil {
				log4go.Debug(err)
				return
			}
			content := buf[0:n]
			HandleMessage(sess, content)
		}
	}()
	for {
		var input string
		fmt.Scanln(&input)
		if input != "" {
			(*sess.Connect).Write([]byte(input))
		}
	}
}
func (sess *PlayerSession)SendPackageMessage(messageType message.REPLY_TYPE, mess proto.Message){
	message.SendPackageMessage(sess.Connect,messageType,mess)
}
//关闭一个连接并且把Map中的数据删除
func (sess *PlayerSession)Close(){
	(*sess.Connect).Close()
	defer delete(SessionMap,sess.GetId())

}
func (sess *PlayerSession)GetId()int64{
	return sess.Info.BasicInfo.Userid
}
func HandleMessage(sess *PlayerSession,data []byte){
	sess.IsConnect=true
	x:=new(message.ProtoMessage)
	proto.Unmarshal(data,x)
	switch message.REQUEST_TYPE(x.MessageType){
	case message.REQUEST_TYPE_Login:
		point:= new(message.LoginRequest)
		proto.Unmarshal(data,point)
		HandleLoginFunc(sess,point)
	case message.REQUEST_TYPE_Match:
		point:= new(message.MatchRequest)
		proto.Unmarshal(data,point)
		HandleMatchFunc(sess,point)
	default :
	}
}