package session

import (
	"github.com/golang/protobuf/proto"
	"ddzserver/message"
	"sync"
	"github.com/alecthomas/log4go"
)

var roomMutex sync.Mutex = sync.Mutex{}

type RoomInfo message.MatchReply

var GameRooms map[int32]RoomInfo = make(map[int32]RoomInfo)

func HandleMatchFunc(sess *PlayerSession, request *message.MatchRequest) {
	roomInfo := JoinRoom(request.Roomid, sess)
	reply := &message.MatchReply{Roomid: request.Roomid, MatchedUser: roomInfo.MatchedUser}
	for _,v:=range roomInfo.MatchedUser{
GetSession(v.Userid) .SendPackageMessage(message.REPLY_TYPE_Match_Result, reply)
	}
}
func GetMatchRequest(data []byte) *message.MatchRequest {
	x := new(message.MatchRequest)
	proto.Unmarshal(data, x)
	return x
}

func JoinRoom(roomId int32, sess *PlayerSession) RoomInfo {
	userId := sess.GetId()
	isExist, hasJoin, room := IsAlreadyJoinRoom(roomId, userId)
	if hasJoin {
		log4go.Error("用户尝试多次加入房间", sess.Info.BasicInfo)
		return room
	}

	if isExist {
		room.MatchedUser = append(room.MatchedUser, sess.Info.BasicInfo)
	} else {
		room = NewRoom(roomId, sess.Info.BasicInfo)
	}
	return room
}

func IsAlreadyJoinRoom(roomId int32, userId int64) (bool, bool, RoomInfo) {
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
func IsRoomExist(roomId int32) (bool, RoomInfo) {
	value, ok := GameRooms[roomId]
	return ok, value
}
func NewRoom(roomId int32, userInfo *message.UserBriefInfo) RoomInfo {
	room := RoomInfo{Roomid: roomId, MatchedUser: []*message.UserBriefInfo{userInfo}}
	GameRooms[roomId] = room
	return room
}
