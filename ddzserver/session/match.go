package session

import (
	"ddzserver/message"
	"sync"
	"github.com/alecthomas/log4go"
)

var roomMutex sync.Mutex = sync.Mutex{}

type RoomInfo message.MatchReply

var GameRooms map[int32]*RoomInfo = make(map[int32]*RoomInfo)

func HandleCancelMatchFunc(sess *PlayerSession, request *message.CancelMatchRequest) {
	ret,roomInfo := QuitRoom(request.Roomid, sess)

	reply := &message.CancelMatchReply{Result:ret ,RoomInfo: (*message.MatchReply)(roomInfo)}
	sess.SendPackageMessage(message.REPLY_TYPE_CancelMatch_Result, reply)
	for _, v := range roomInfo.MatchedUser {
		GetSession(v.Userid).SendPackageMessage(message.REPLY_TYPE_CancelMatch_Result, reply)//给每一位在房间中的玩家发送一个信息
	}
}
func HandleMatchFunc(sess *PlayerSession, request *message.MatchRequest) {
	roomInfo := JoinRoom(request.Roomid, sess)
	var MatchedUser []*message.UserBriefInfo
	if roomInfo==nil{
		MatchedUser=nil
	}else{
		MatchedUser=roomInfo.MatchedUser
	}
	reply := &message.MatchReply{Roomid: request.Roomid, MatchedUser:MatchedUser}
	for _, v := range roomInfo.MatchedUser {
		GetSession(v.Userid).SendPackageMessage(message.REPLY_TYPE_Match_Result, reply)//给每一位在房间中的玩家发送一个信息
	}
}

func QuitRoom(roomId int32, sess *PlayerSession) (message.CANCEL_MATCH_RESULT,*RoomInfo) {
	userId := sess.GetId()
	if userId<0{
		return message.CANCEL_MATCH_RESULT_CANCEL_MATCH_NOT_LOGIN, nil
	}
	isExist, hasJoin, room := IsAlreadyJoinRoom(roomId, userId)
	if hasJoin&&isExist {
		DeleteRoomUser(room,userId)
		return  message.CANCEL_MATCH_RESULT_CANCEL_MATCH_SUCCESS,room
	}else{
		if !isExist{
			return message.CANCEL_MATCH_RESULT_CANCEL_MATCH_NO_ROOM,room
		}else{
			log4go.Error("不存在的房间或者用户尝试退出房间", userId)
			return message.CANCEL_MATCH_RESULT_CANCEL_MATCH_NO_USER,room
		}
	}
}

func JoinRoom(roomId int32, sess *PlayerSession) *RoomInfo {
	userId := sess.GetId()
	if userId<0{
		return nil
	}
	isExist, hasJoin, room := IsAlreadyJoinRoom(roomId, userId)
	if hasJoin {
		log4go.Error("用户尝试多次加入房间", userId)
		return room
	}

	if isExist {
		room.MatchedUser = append(room.MatchedUser, sess.Info.BasicInfo)
	} else {
		room = NewRoom(roomId, sess.Info.BasicInfo)
	}
	return room
}
func DeleteRoomUser(room *RoomInfo,userId int64){
	delIndex:=-1
	for i,v:=range room.MatchedUser{
		if v.Userid==userId{
			delIndex=i
		}
	}
	l:=len(room.MatchedUser)
	if delIndex<0{
		return
	}
	if delIndex==0{
		room.MatchedUser=room.MatchedUser[1:]
	}else if delIndex==l{
		room.MatchedUser=room.MatchedUser[:l-1]
	}else {
		room.MatchedUser=append(room.MatchedUser[0:delIndex],room.MatchedUser[delIndex+1:]...)
	}
}
func IsAlreadyJoinRoom(roomId int32, userId int64) (bool, bool, *RoomInfo) {
	isExist, room := IsRoomExist(roomId)
	if !isExist {
		return isExist, false, room
	}
	for _, v := range room.MatchedUser {
		if v.Userid == userId {
			return isExist, true, room
		}
	}
	return isExist, false, room
}
func IsRoomExist(roomId int32) (bool, *RoomInfo) {
	value, ok := GameRooms[roomId]
	return ok, value
}
func NewRoom(roomId int32, userInfo *message.UserBriefInfo) *RoomInfo {
	room := &RoomInfo{Roomid: roomId, MatchedUser: []*message.UserBriefInfo{userInfo}}
	GameRooms[roomId] = room
	return room
}
