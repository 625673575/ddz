package session

import (
	"ddzserver/message"
	"ddzserver/login/logindb"
	"log"
	"github.com/alecthomas/log4go"
	"strings"
)

func HandleLoginFunc(sess *PlayerSession, request *message.LoginRequest) {
	hasUser, user := logindb.QueryLoginUser(request.Username)
	log.Println(hasUser, user)
	var stateCode message.LOGIN_RESULT
	if hasUser {
		if user.Password == request.Password {
			stateCode = message.LOGIN_RESULT_LOGIN_SUCCESS
			sess.HasLogin = true
			hasUser, sess.Info = logindb.QueryUserInfo(user.Userid)
			if (!hasUser) {
				log4go.Error("找到了登陆用户但没有对应的用户信息", user)
			}
			SessionMap[sess.GetId()]=sess
		} else {
			stateCode = message.LOGIN_RESULT_LOGIN_WRONG_PASSWORD
		}
	} else {
		stateCode = message.LOGIN_RESULT_LOGIN_NO_USER
	}
	reply := &message.LoginReply{StateCode: stateCode, UserInfo: &sess.Info}
	sess.SendPackageMessage(message.REPLY_TYPE_Login_Result, reply)
}
func HandleRegisterFunc(sess *PlayerSession, request *message.RegisterRequest) {
	hasUser, _ := logindb.QueryLoginUser(request.Username)
	reply:=&message.RegisterReply{}
	if hasUser{
		reply.StateCode=message.REGISTER_RESULT_REGISTER_ALREADY_EXIST
	}else{
		if !isValidPassword(request.Password){
reply.StateCode=message.REGISTER_RESULT_REGISTER_INVALID_PASSWORD
		}
		if !isValidUsername(request.Username){
			reply.StateCode=message.REGISTER_RESULT_REGISTER_INVALID_NAME
		}
	}
	if reply.StateCode==message.REGISTER_RESULT_REGISTER_SUCCESS{
		logindb.InsertLoginUser(request.Username,request.Password)
		_,user:=logindb.QueryLoginUser(request.Username)
		userid:=user.Userid
		logindb.InsertUserInfo(userid,request.Username,0)
		reply.UserInfo=new(message.UserInfo)
		reply.UserInfo.BasicInfo=&message.UserBriefInfo{Userid:userid,UserName:request.Username,UserIcon:0}
		reply.UserInfo.Shit="shit"
	}
	sess.SendPackageMessage(message.REPLY_TYPE_Register_Result, reply)
}

func isValidPassword(pwd string)bool{
	if len(pwd)<4 || strings.Contains(pwd," "){
		return false
	}
	return true
}
func isValidUsername(pwd string)bool{
	if len(pwd)<4 || strings.Contains(pwd," "){
		return false
	}
	return true
}