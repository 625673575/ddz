package session

import (
	"github.com/golang/protobuf/proto"
	"ddzserver/message"
	"ddzserver/login/logindb"
	"log"
	"github.com/alecthomas/log4go"
)

func HandleLoginFunc(sess *PlayerSession, request *message.LoginRequest) {
	hasUser, user := logindb.QueryLoginUser(request.Username)
	log.Println(hasUser, user)
	var stateCode message.LOGIN_RESULT
	if hasUser {
		if user.Password == request.Password {
			stateCode = message.LOGIN_RESULT_SUCCESS
			sess.HasLogin = true
			hasUser, sess.Info = logindb.QueryUserInfo(user.Userid)
			if (!hasUser) {
				log4go.Error("找到了登陆用户但没有对应的用户信息", user)
			}
			SessionMap[sess.GetId()]=sess
		} else {
			stateCode = message.LOGIN_RESULT_WRONG_PASSWORD
		}
	} else {
		stateCode = message.LOGIN_RESULT_NO_USER
	}
	reply := &message.LoginReply{StateCode: stateCode, UserInfo: &sess.Info}
	message.SendPackageMessage(sess.Connect, message.REPLY_TYPE_Login_Result, reply)
}

func GetLoginRequest(data []byte) *message.LoginRequest {
	x := new(message.LoginRequest)
	proto.Unmarshal(data, x)
	return x
}
