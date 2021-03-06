// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	ProtoMessage
	LoginRequest
	LoginReply
	RegisterRequest
	RegisterReply
	MatchRequest
	CancelMatchRequest
	MatchReply
	CancelMatchReply
	CardInfo
	PlayCardInfo
	UserBriefInfo
	UserInfo
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type REQUEST_TYPE int32

const (
	REQUEST_TYPE_No_Request  REQUEST_TYPE = 0
	REQUEST_TYPE_Login       REQUEST_TYPE = 1
	REQUEST_TYPE_Register    REQUEST_TYPE = 2
	REQUEST_TYPE_Match       REQUEST_TYPE = 10
	REQUEST_TYPE_CancelMatch REQUEST_TYPE = 11
)

var REQUEST_TYPE_name = map[int32]string{
	0:  "No_Request",
	1:  "Login",
	2:  "Register",
	10: "Match",
	11: "CancelMatch",
}
var REQUEST_TYPE_value = map[string]int32{
	"No_Request":  0,
	"Login":       1,
	"Register":    2,
	"Match":       10,
	"CancelMatch": 11,
}

func (x REQUEST_TYPE) String() string {
	return proto.EnumName(REQUEST_TYPE_name, int32(x))
}
func (REQUEST_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type REPLY_TYPE int32

const (
	REPLY_TYPE_No_Result          REPLY_TYPE = 0
	REPLY_TYPE_Login_Result       REPLY_TYPE = 1
	REPLY_TYPE_Register_Result    REPLY_TYPE = 2
	REPLY_TYPE_Match_Result       REPLY_TYPE = 10
	REPLY_TYPE_CancelMatch_Result REPLY_TYPE = 11
)

var REPLY_TYPE_name = map[int32]string{
	0:  "No_Result",
	1:  "Login_Result",
	2:  "Register_Result",
	10: "Match_Result",
	11: "CancelMatch_Result",
}
var REPLY_TYPE_value = map[string]int32{
	"No_Result":          0,
	"Login_Result":       1,
	"Register_Result":    2,
	"Match_Result":       10,
	"CancelMatch_Result": 11,
}

func (x REPLY_TYPE) String() string {
	return proto.EnumName(REPLY_TYPE_name, int32(x))
}
func (REPLY_TYPE) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// 花色
type Suits int32

const (
	Suits_Club      Suits = 0
	Suits_Diamond   Suits = 1
	Suits_Heart     Suits = 2
	Suits_Spade     Suits = 3
	Suits_NoneSuits Suits = 4
)

var Suits_name = map[int32]string{
	0: "Club",
	1: "Diamond",
	2: "Heart",
	3: "Spade",
	4: "NoneSuits",
}
var Suits_value = map[string]int32{
	"Club":      0,
	"Diamond":   1,
	"Heart":     2,
	"Spade":     3,
	"NoneSuits": 4,
}

func (x Suits) String() string {
	return proto.EnumName(Suits_name, int32(x))
}
func (Suits) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// / 卡牌权值
type Weight int32

const (
	Weight_Three  Weight = 0
	Weight_Four   Weight = 1
	Weight_Five   Weight = 2
	Weight_Six    Weight = 3
	Weight_Seven  Weight = 4
	Weight_Eight  Weight = 5
	Weight_Nine   Weight = 6
	Weight_Ten    Weight = 7
	Weight_Jack   Weight = 8
	Weight_Queen  Weight = 9
	Weight_King   Weight = 10
	Weight_One    Weight = 11
	Weight_Two    Weight = 12
	Weight_SJoker Weight = 13
	Weight_LJoker Weight = 14
)

var Weight_name = map[int32]string{
	0:  "Three",
	1:  "Four",
	2:  "Five",
	3:  "Six",
	4:  "Seven",
	5:  "Eight",
	6:  "Nine",
	7:  "Ten",
	8:  "Jack",
	9:  "Queen",
	10: "King",
	11: "One",
	12: "Two",
	13: "SJoker",
	14: "LJoker",
}
var Weight_value = map[string]int32{
	"Three":  0,
	"Four":   1,
	"Five":   2,
	"Six":    3,
	"Seven":  4,
	"Eight":  5,
	"Nine":   6,
	"Ten":    7,
	"Jack":   8,
	"Queen":  9,
	"King":   10,
	"One":    11,
	"Two":    12,
	"SJoker": 13,
	"LJoker": 14,
}

func (x Weight) String() string {
	return proto.EnumName(Weight_name, int32(x))
}
func (Weight) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// / 身份
type Identity int32

const (
	Identity_Farmer   Identity = 0
	Identity_Landlord Identity = 1
)

var Identity_name = map[int32]string{
	0: "Farmer",
	1: "Landlord",
}
var Identity_value = map[string]int32{
	"Farmer":   0,
	"Landlord": 1,
}

func (x Identity) String() string {
	return proto.EnumName(Identity_name, int32(x))
}
func (Identity) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// / 出牌类型
type CardsType int32

const (
	// 未知类型
	CardsType_UnknowCard CardsType = 0
	// 王炸
	CardsType_JokerBoom CardsType = 1
	// 炸弹
	CardsType_Boom CardsType = 2
	// 三个不带
	CardsType_OnlyThree CardsType = 3
	// 三个带一
	CardsType_ThreeAndOne CardsType = 4
	// 三个带二
	CardsType_ThreeAndTwo CardsType = 5
	// 顺子 五张或更多的连续单牌
	CardsType_Straight CardsType = 6
	// 双顺 三对或更多的连续对牌
	CardsType_DoubleStraight CardsType = 7
	// 三顺 二个或更多的连续三张牌
	CardsType_TripleStraight CardsType = 8
	// 对子
	CardsType_Double CardsType = 9
	// 单个
	CardsType_Single CardsType = 10
)

var CardsType_name = map[int32]string{
	0:  "UnknowCard",
	1:  "JokerBoom",
	2:  "Boom",
	3:  "OnlyThree",
	4:  "ThreeAndOne",
	5:  "ThreeAndTwo",
	6:  "Straight",
	7:  "DoubleStraight",
	8:  "TripleStraight",
	9:  "Double",
	10: "Single",
}
var CardsType_value = map[string]int32{
	"UnknowCard":     0,
	"JokerBoom":      1,
	"Boom":           2,
	"OnlyThree":      3,
	"ThreeAndOne":    4,
	"ThreeAndTwo":    5,
	"Straight":       6,
	"DoubleStraight": 7,
	"TripleStraight": 8,
	"Double":         9,
	"Single":         10,
}

func (x CardsType) String() string {
	return proto.EnumName(CardsType_name, int32(x))
}
func (CardsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// 角色类型
type PlayerRole int32

const (
	PlayerRole_Library     PlayerRole = 0
	PlayerRole_PlayerOne   PlayerRole = 1
	PlayerRole_PlayerTwo   PlayerRole = 2
	PlayerRole_PlayerThree PlayerRole = 3
	PlayerRole_PlayerFour  PlayerRole = 4
	PlayerRole_PlayerFive  PlayerRole = 5
	PlayerRole_PlayerSix   PlayerRole = 6
)

var PlayerRole_name = map[int32]string{
	0: "Library",
	1: "PlayerOne",
	2: "PlayerTwo",
	3: "PlayerThree",
	4: "PlayerFour",
	5: "PlayerFive",
	6: "PlayerSix",
}
var PlayerRole_value = map[string]int32{
	"Library":     0,
	"PlayerOne":   1,
	"PlayerTwo":   2,
	"PlayerThree": 3,
	"PlayerFour":  4,
	"PlayerFive":  5,
	"PlayerSix":   6,
}

func (x PlayerRole) String() string {
	return proto.EnumName(PlayerRole_name, int32(x))
}
func (PlayerRole) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type LOGIN_RESULT int32

const (
	LOGIN_RESULT_LOGIN_SUCCESS        LOGIN_RESULT = 0
	LOGIN_RESULT_LOGIN_WRONG_PASSWORD LOGIN_RESULT = 1
	LOGIN_RESULT_LOGIN_NO_USER        LOGIN_RESULT = 2
	LOGIN_RESULT_LOGIN_UNKNOW_ERROR   LOGIN_RESULT = 3
)

var LOGIN_RESULT_name = map[int32]string{
	0: "LOGIN_SUCCESS",
	1: "LOGIN_WRONG_PASSWORD",
	2: "LOGIN_NO_USER",
	3: "LOGIN_UNKNOW_ERROR",
}
var LOGIN_RESULT_value = map[string]int32{
	"LOGIN_SUCCESS":        0,
	"LOGIN_WRONG_PASSWORD": 1,
	"LOGIN_NO_USER":        2,
	"LOGIN_UNKNOW_ERROR":   3,
}

func (x LOGIN_RESULT) String() string {
	return proto.EnumName(LOGIN_RESULT_name, int32(x))
}
func (LOGIN_RESULT) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type REGISTER_RESULT int32

const (
	REGISTER_RESULT_REGISTER_SUCCESS          REGISTER_RESULT = 0
	REGISTER_RESULT_REGISTER_ALREADY_EXIST    REGISTER_RESULT = 1
	REGISTER_RESULT_REGISTER_INVALID_NAME     REGISTER_RESULT = 2
	REGISTER_RESULT_REGISTER_INVALID_PASSWORD REGISTER_RESULT = 3
)

var REGISTER_RESULT_name = map[int32]string{
	0: "REGISTER_SUCCESS",
	1: "REGISTER_ALREADY_EXIST",
	2: "REGISTER_INVALID_NAME",
	3: "REGISTER_INVALID_PASSWORD",
}
var REGISTER_RESULT_value = map[string]int32{
	"REGISTER_SUCCESS":          0,
	"REGISTER_ALREADY_EXIST":    1,
	"REGISTER_INVALID_NAME":     2,
	"REGISTER_INVALID_PASSWORD": 3,
}

func (x REGISTER_RESULT) String() string {
	return proto.EnumName(REGISTER_RESULT_name, int32(x))
}
func (REGISTER_RESULT) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type CANCEL_MATCH_RESULT int32

const (
	CANCEL_MATCH_RESULT_CANCEL_MATCH_SUCCESS   CANCEL_MATCH_RESULT = 0
	CANCEL_MATCH_RESULT_CANCEL_MATCH_NO_ROOM   CANCEL_MATCH_RESULT = 1
	CANCEL_MATCH_RESULT_CANCEL_MATCH_NO_USER   CANCEL_MATCH_RESULT = 2
	CANCEL_MATCH_RESULT_CANCEL_MATCH_NOT_LOGIN CANCEL_MATCH_RESULT = 3
)

var CANCEL_MATCH_RESULT_name = map[int32]string{
	0: "CANCEL_MATCH_SUCCESS",
	1: "CANCEL_MATCH_NO_ROOM",
	2: "CANCEL_MATCH_NO_USER",
	3: "CANCEL_MATCH_NOT_LOGIN",
}
var CANCEL_MATCH_RESULT_value = map[string]int32{
	"CANCEL_MATCH_SUCCESS":   0,
	"CANCEL_MATCH_NO_ROOM":   1,
	"CANCEL_MATCH_NO_USER":   2,
	"CANCEL_MATCH_NOT_LOGIN": 3,
}

func (x CANCEL_MATCH_RESULT) String() string {
	return proto.EnumName(CANCEL_MATCH_RESULT_name, int32(x))
}
func (CANCEL_MATCH_RESULT) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ProtoMessage struct {
	MessageType int32  `protobuf:"varint,1,opt,name=MessageType" json:"MessageType,omitempty"`
	Data        []byte `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *ProtoMessage) Reset()                    { *m = ProtoMessage{} }
func (m *ProtoMessage) String() string            { return proto.CompactTextString(m) }
func (*ProtoMessage) ProtoMessage()               {}
func (*ProtoMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ProtoMessage) GetMessageType() int32 {
	if m != nil {
		return m.MessageType
	}
	return 0
}

func (m *ProtoMessage) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginReply struct {
	StateCode LOGIN_RESULT `protobuf:"varint,1,opt,name=StateCode,enum=LOGIN_RESULT" json:"StateCode,omitempty"`
	UserInfo  *UserInfo    `protobuf:"bytes,2,opt,name=UserInfo" json:"UserInfo,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginReply) GetStateCode() LOGIN_RESULT {
	if m != nil {
		return m.StateCode
	}
	return LOGIN_RESULT_LOGIN_SUCCESS
}

func (m *LoginReply) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type RegisterRequest struct {
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
}

func (m *RegisterRequest) Reset()                    { *m = RegisterRequest{} }
func (m *RegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterRequest) ProtoMessage()               {}
func (*RegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RegisterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RegisterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RegisterReply struct {
	StateCode REGISTER_RESULT `protobuf:"varint,1,opt,name=StateCode,enum=REGISTER_RESULT" json:"StateCode,omitempty"`
	UserInfo  *UserInfo       `protobuf:"bytes,2,opt,name=UserInfo" json:"UserInfo,omitempty"`
}

func (m *RegisterReply) Reset()                    { *m = RegisterReply{} }
func (m *RegisterReply) String() string            { return proto.CompactTextString(m) }
func (*RegisterReply) ProtoMessage()               {}
func (*RegisterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RegisterReply) GetStateCode() REGISTER_RESULT {
	if m != nil {
		return m.StateCode
	}
	return REGISTER_RESULT_REGISTER_SUCCESS
}

func (m *RegisterReply) GetUserInfo() *UserInfo {
	if m != nil {
		return m.UserInfo
	}
	return nil
}

type MatchRequest struct {
	Roomid int32 `protobuf:"varint,2,opt,name=Roomid" json:"Roomid,omitempty"`
}

func (m *MatchRequest) Reset()                    { *m = MatchRequest{} }
func (m *MatchRequest) String() string            { return proto.CompactTextString(m) }
func (*MatchRequest) ProtoMessage()               {}
func (*MatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MatchRequest) GetRoomid() int32 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

type CancelMatchRequest struct {
	Roomid int32 `protobuf:"varint,2,opt,name=Roomid" json:"Roomid,omitempty"`
}

func (m *CancelMatchRequest) Reset()                    { *m = CancelMatchRequest{} }
func (m *CancelMatchRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelMatchRequest) ProtoMessage()               {}
func (*CancelMatchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CancelMatchRequest) GetRoomid() int32 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

type MatchReply struct {
	Roomid      int32            `protobuf:"varint,1,opt,name=Roomid" json:"Roomid,omitempty"`
	MatchedUser []*UserBriefInfo `protobuf:"bytes,2,rep,name=MatchedUser" json:"MatchedUser,omitempty"`
}

func (m *MatchReply) Reset()                    { *m = MatchReply{} }
func (m *MatchReply) String() string            { return proto.CompactTextString(m) }
func (*MatchReply) ProtoMessage()               {}
func (*MatchReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MatchReply) GetRoomid() int32 {
	if m != nil {
		return m.Roomid
	}
	return 0
}

func (m *MatchReply) GetMatchedUser() []*UserBriefInfo {
	if m != nil {
		return m.MatchedUser
	}
	return nil
}

type CancelMatchReply struct {
	Result   CANCEL_MATCH_RESULT `protobuf:"varint,1,opt,name=Result,enum=CANCEL_MATCH_RESULT" json:"Result,omitempty"`
	RoomInfo *MatchReply         `protobuf:"bytes,2,opt,name=RoomInfo" json:"RoomInfo,omitempty"`
}

func (m *CancelMatchReply) Reset()                    { *m = CancelMatchReply{} }
func (m *CancelMatchReply) String() string            { return proto.CompactTextString(m) }
func (*CancelMatchReply) ProtoMessage()               {}
func (*CancelMatchReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CancelMatchReply) GetResult() CANCEL_MATCH_RESULT {
	if m != nil {
		return m.Result
	}
	return CANCEL_MATCH_RESULT_CANCEL_MATCH_SUCCESS
}

func (m *CancelMatchReply) GetRoomInfo() *MatchReply {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

type CardInfo struct {
	CardIndex int32  `protobuf:"varint,1,opt,name=CardIndex" json:"CardIndex,omitempty"`
	Weight    Weight `protobuf:"varint,2,opt,name=Weight,enum=Weight" json:"Weight,omitempty"`
	Color     Suits  `protobuf:"varint,3,opt,name=Color,enum=Suits" json:"Color,omitempty"`
}

func (m *CardInfo) Reset()                    { *m = CardInfo{} }
func (m *CardInfo) String() string            { return proto.CompactTextString(m) }
func (*CardInfo) ProtoMessage()               {}
func (*CardInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CardInfo) GetCardIndex() int32 {
	if m != nil {
		return m.CardIndex
	}
	return 0
}

func (m *CardInfo) GetWeight() Weight {
	if m != nil {
		return m.Weight
	}
	return Weight_Three
}

func (m *CardInfo) GetColor() Suits {
	if m != nil {
		return m.Color
	}
	return Suits_Club
}

type PlayCardInfo struct {
	User      int32       `protobuf:"varint,1,opt,name=User" json:"User,omitempty"`
	CardsInfo []*CardInfo `protobuf:"bytes,2,rep,name=CardsInfo" json:"CardsInfo,omitempty"`
}

func (m *PlayCardInfo) Reset()                    { *m = PlayCardInfo{} }
func (m *PlayCardInfo) String() string            { return proto.CompactTextString(m) }
func (*PlayCardInfo) ProtoMessage()               {}
func (*PlayCardInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *PlayCardInfo) GetUser() int32 {
	if m != nil {
		return m.User
	}
	return 0
}

func (m *PlayCardInfo) GetCardsInfo() []*CardInfo {
	if m != nil {
		return m.CardsInfo
	}
	return nil
}

type UserBriefInfo struct {
	Userid   int64  `protobuf:"varint,1,opt,name=Userid" json:"Userid,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=UserName" json:"UserName,omitempty"`
	UserIcon int32  `protobuf:"varint,3,opt,name=UserIcon" json:"UserIcon,omitempty"`
	Rank     int32  `protobuf:"varint,4,opt,name=Rank" json:"Rank,omitempty"`
}

func (m *UserBriefInfo) Reset()                    { *m = UserBriefInfo{} }
func (m *UserBriefInfo) String() string            { return proto.CompactTextString(m) }
func (*UserBriefInfo) ProtoMessage()               {}
func (*UserBriefInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UserBriefInfo) GetUserid() int64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *UserBriefInfo) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserBriefInfo) GetUserIcon() int32 {
	if m != nil {
		return m.UserIcon
	}
	return 0
}

func (m *UserBriefInfo) GetRank() int32 {
	if m != nil {
		return m.Rank
	}
	return 0
}

type UserInfo struct {
	BasicInfo *UserBriefInfo `protobuf:"bytes,1,opt,name=BasicInfo" json:"BasicInfo,omitempty"`
	Shit      string         `protobuf:"bytes,2,opt,name=Shit" json:"Shit,omitempty"`
}

func (m *UserInfo) Reset()                    { *m = UserInfo{} }
func (m *UserInfo) String() string            { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()               {}
func (*UserInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UserInfo) GetBasicInfo() *UserBriefInfo {
	if m != nil {
		return m.BasicInfo
	}
	return nil
}

func (m *UserInfo) GetShit() string {
	if m != nil {
		return m.Shit
	}
	return ""
}

func init() {
	proto.RegisterType((*ProtoMessage)(nil), "ProtoMessage")
	proto.RegisterType((*LoginRequest)(nil), "LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "LoginReply")
	proto.RegisterType((*RegisterRequest)(nil), "RegisterRequest")
	proto.RegisterType((*RegisterReply)(nil), "RegisterReply")
	proto.RegisterType((*MatchRequest)(nil), "MatchRequest")
	proto.RegisterType((*CancelMatchRequest)(nil), "CancelMatchRequest")
	proto.RegisterType((*MatchReply)(nil), "MatchReply")
	proto.RegisterType((*CancelMatchReply)(nil), "CancelMatchReply")
	proto.RegisterType((*CardInfo)(nil), "CardInfo")
	proto.RegisterType((*PlayCardInfo)(nil), "PlayCardInfo")
	proto.RegisterType((*UserBriefInfo)(nil), "UserBriefInfo")
	proto.RegisterType((*UserInfo)(nil), "UserInfo")
	proto.RegisterEnum("REQUEST_TYPE", REQUEST_TYPE_name, REQUEST_TYPE_value)
	proto.RegisterEnum("REPLY_TYPE", REPLY_TYPE_name, REPLY_TYPE_value)
	proto.RegisterEnum("Suits", Suits_name, Suits_value)
	proto.RegisterEnum("Weight", Weight_name, Weight_value)
	proto.RegisterEnum("Identity", Identity_name, Identity_value)
	proto.RegisterEnum("CardsType", CardsType_name, CardsType_value)
	proto.RegisterEnum("PlayerRole", PlayerRole_name, PlayerRole_value)
	proto.RegisterEnum("LOGIN_RESULT", LOGIN_RESULT_name, LOGIN_RESULT_value)
	proto.RegisterEnum("REGISTER_RESULT", REGISTER_RESULT_name, REGISTER_RESULT_value)
	proto.RegisterEnum("CANCEL_MATCH_RESULT", CANCEL_MATCH_RESULT_name, CANCEL_MATCH_RESULT_value)
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1060 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0x16, 0x75, 0xb2, 0x34, 0x3a, 0x64, 0xff, 0x4d, 0x7e, 0x43, 0x31, 0x52, 0xd4, 0x10, 0xda,
	0xc6, 0x50, 0x0d, 0xa1, 0x70, 0x9f, 0x40, 0x96, 0x68, 0x47, 0xb1, 0x4c, 0x2a, 0x4b, 0x2a, 0xae,
	0xaf, 0x58, 0xda, 0x5c, 0xcb, 0x84, 0x29, 0xae, 0x4b, 0x52, 0x71, 0x74, 0x51, 0xf4, 0x5d, 0xfa,
	0x0a, 0x7d, 0xc1, 0x62, 0x76, 0x79, 0x90, 0x9c, 0x5c, 0x14, 0xed, 0xdd, 0xcc, 0x37, 0xb3, 0xdf,
	0x7e, 0x73, 0x58, 0x4a, 0xd0, 0x59, 0xf1, 0x38, 0x76, 0x97, 0x7c, 0xf8, 0x18, 0x89, 0x44, 0xf4,
	0x27, 0xd0, 0x9e, 0xa3, 0x71, 0xa9, 0x50, 0x7a, 0x08, 0xad, 0xd4, 0xb4, 0x37, 0x8f, 0xbc, 0xa7,
	0x1d, 0x6a, 0x47, 0x35, 0xb6, 0x0d, 0x51, 0x0a, 0xd5, 0x89, 0x9b, 0xb8, 0xbd, 0xf2, 0xa1, 0x76,
	0xd4, 0x66, 0xd2, 0xee, 0x9f, 0x41, 0x7b, 0x26, 0x96, 0x7e, 0xc8, 0xf8, 0x6f, 0x6b, 0x1e, 0x27,
	0xf4, 0x00, 0x1a, 0x8b, 0x98, 0x47, 0xa1, 0xbb, 0x52, 0x14, 0x4d, 0x96, 0xfb, 0x18, 0x9b, 0xbb,
	0x71, 0xfc, 0x24, 0x22, 0x4f, 0x72, 0x34, 0x59, 0xee, 0xf7, 0x7f, 0x05, 0x48, 0x79, 0x1e, 0x83,
	0x0d, 0xfd, 0x11, 0x9a, 0x56, 0xe2, 0x26, 0x7c, 0x2c, 0x3c, 0x45, 0xd3, 0x3d, 0xe9, 0x0c, 0x67,
	0xe6, 0xf9, 0xd4, 0x70, 0x98, 0x6e, 0x2d, 0x66, 0x36, 0x2b, 0xe2, 0xf4, 0x7b, 0x75, 0xe5, 0x34,
	0xbc, 0x13, 0x92, 0xb6, 0x75, 0xd2, 0x1c, 0x66, 0x00, 0xcb, 0x43, 0xfd, 0x29, 0xbc, 0x60, 0x7c,
	0xe9, 0xc7, 0x09, 0x8f, 0xfe, 0xab, 0xd8, 0x3b, 0xe8, 0x14, 0x54, 0xa8, 0x77, 0xf8, 0xa5, 0x5e,
	0x32, 0x64, 0xfa, 0xf9, 0xd4, 0xb2, 0x75, 0xf6, 0xef, 0x25, 0xff, 0x00, 0xed, 0x4b, 0x37, 0xb9,
	0xbd, 0xcf, 0xf4, 0xee, 0x43, 0x9d, 0x09, 0xb1, 0xf2, 0x95, 0xa2, 0x1a, 0x4b, 0xbd, 0xfe, 0x31,
	0xd0, 0xb1, 0x1b, 0xde, 0xf2, 0xe0, 0x1f, 0x65, 0x7f, 0x04, 0x48, 0xf3, 0x50, 0x7a, 0x91, 0xa5,
	0x6d, 0x67, 0xd1, 0x9f, 0xa0, 0x25, 0xb3, 0xb8, 0x87, 0x72, 0x7a, 0xe5, 0xc3, 0xca, 0x51, 0xeb,
	0xa4, 0x2b, 0x55, 0x9e, 0x46, 0x3e, 0xbf, 0x93, 0x52, 0xb7, 0x53, 0xfa, 0x3e, 0x90, 0x1d, 0x15,
	0xc8, 0x7e, 0x0c, 0x75, 0xc6, 0xe3, 0x75, 0x90, 0xa4, 0x5d, 0x79, 0x35, 0x1c, 0x8f, 0x8c, 0xb1,
	0x3e, 0x73, 0x2e, 0x47, 0xf6, 0xf8, 0x5d, 0xd6, 0x99, 0x34, 0x87, 0xbe, 0x85, 0x06, 0xde, 0xbe,
	0xd5, 0x96, 0xd6, 0xb0, 0x20, 0x63, 0x79, 0xb0, 0xbf, 0x84, 0xc6, 0xd8, 0x8d, 0x3c, 0xb4, 0xe9,
	0x1b, 0x68, 0x2a, 0xdb, 0xe3, 0x9f, 0xd3, 0x1a, 0x0a, 0x80, 0x7e, 0x0b, 0xf5, 0x2b, 0xee, 0x2f,
	0xef, 0x13, 0x49, 0xd8, 0x3d, 0xd9, 0x1b, 0x2a, 0x97, 0xa5, 0x30, 0x7d, 0x03, 0xb5, 0xb1, 0x08,
	0x44, 0xd4, 0xab, 0xc8, 0x78, 0x7d, 0x68, 0xad, 0xfd, 0x24, 0x66, 0x0a, 0xec, 0x5f, 0x40, 0x7b,
	0x1e, 0xb8, 0x9b, 0xfc, 0x32, 0x0a, 0x55, 0xd9, 0x0e, 0x75, 0x8f, 0xb4, 0xe9, 0x5b, 0x25, 0x20,
	0x4e, 0x65, 0x57, 0xe4, 0x34, 0xb3, 0x13, 0xac, 0x88, 0xf5, 0x63, 0xe8, 0xec, 0xb4, 0x0f, 0x7b,
	0x8f, 0x40, 0xda, 0xfb, 0x0a, 0x4b, 0xbd, 0x6c, 0x2f, 0x0d, 0xdc, 0xcb, 0x72, 0xb1, 0x97, 0x46,
	0xba, 0x97, 0x72, 0x3f, 0x6e, 0x45, 0x28, 0x25, 0xd7, 0x58, 0xee, 0xa3, 0x3a, 0xe6, 0x86, 0x0f,
	0xbd, 0xaa, 0x52, 0x87, 0x76, 0x7f, 0x56, 0xac, 0x1a, 0x3d, 0x86, 0xe6, 0xa9, 0x1b, 0xfb, 0xb7,
	0x52, 0xa9, 0x26, 0x1b, 0xfc, 0x7c, 0xa2, 0x45, 0x02, 0xb2, 0x59, 0xf7, 0x7e, 0x92, 0x2a, 0x90,
	0xf6, 0xc0, 0x82, 0x36, 0xd3, 0x3f, 0x2c, 0x74, 0xcb, 0x76, 0xec, 0xeb, 0xb9, 0x4e, 0xbb, 0x00,
	0x86, 0x70, 0xd2, 0x8d, 0x23, 0x25, 0xda, 0x84, 0x9a, 0x7c, 0xc6, 0x44, 0xa3, 0x6d, 0x68, 0x64,
	0x8f, 0x84, 0x94, 0x31, 0x20, 0x27, 0x49, 0x80, 0xbe, 0x80, 0xd6, 0xd6, 0x9e, 0x90, 0xd6, 0x20,
	0x00, 0x60, 0xfa, 0x7c, 0x76, 0xad, 0x28, 0x3b, 0xd0, 0x94, 0x94, 0xb8, 0x11, 0xa4, 0x44, 0x49,
	0xfa, 0x81, 0xc9, 0x10, 0x8d, 0xbe, 0x2c, 0x1e, 0x72, 0x06, 0x96, 0x31, 0x4d, 0xd2, 0x65, 0x08,
	0xd0, 0xfd, 0x9d, 0x47, 0x91, 0xe1, 0xad, 0xc1, 0x18, 0x6a, 0x72, 0xc4, 0xb4, 0x01, 0xd5, 0x71,
	0xb0, 0xbe, 0x21, 0x25, 0xda, 0x82, 0xbd, 0x89, 0xef, 0xae, 0x44, 0xe8, 0x11, 0x0d, 0x95, 0xbe,
	0xe3, 0x6e, 0x94, 0x28, 0xd1, 0xd6, 0xa3, 0xeb, 0x71, 0x52, 0x51, 0xaa, 0x42, 0x2e, 0x4f, 0x92,
	0xea, 0xe0, 0x4f, 0x2d, 0xdb, 0x2b, 0x4c, 0xb2, 0xef, 0x23, 0xce, 0x49, 0x09, 0x19, 0xcf, 0xc4,
	0x3a, 0x22, 0x9a, 0xb4, 0xfc, 0x4f, 0x9c, 0x94, 0xe9, 0x1e, 0x54, 0x2c, 0xff, 0x33, 0xa9, 0x48,
	0x32, 0xfe, 0x89, 0x87, 0xa4, 0x8a, 0xa6, 0x8e, 0x67, 0x49, 0x0d, 0x13, 0x0d, 0x3f, 0xe4, 0xa4,
	0x8e, 0x89, 0x36, 0x0f, 0xc9, 0x1e, 0x42, 0xef, 0xdd, 0xdb, 0x07, 0xd2, 0xc0, 0xbc, 0x0f, 0x6b,
	0xce, 0x43, 0xd2, 0x44, 0xf0, 0xc2, 0x0f, 0x97, 0x04, 0x30, 0xcf, 0x0c, 0x39, 0x69, 0xc9, 0x03,
	0x4f, 0x82, 0xb4, 0x29, 0x40, 0xdd, 0x7a, 0x2f, 0x1e, 0x78, 0x44, 0x3a, 0x68, 0xcf, 0x94, 0xdd,
	0x1d, 0x7c, 0x07, 0x8d, 0xa9, 0xc7, 0xc3, 0xc4, 0x4f, 0x36, 0x88, 0x9f, 0xb9, 0xd1, 0x8a, 0x47,
	0xa4, 0x84, 0x93, 0x99, 0xb9, 0xa1, 0x17, 0x88, 0xc8, 0x23, 0xda, 0xe0, 0x2f, 0x2d, 0xdd, 0x5f,
	0xf9, 0x8d, 0xef, 0x02, 0x2c, 0xc2, 0x87, 0x50, 0x3c, 0x21, 0x44, 0x4a, 0x58, 0xb7, 0xa4, 0x3b,
	0x15, 0x62, 0xa5, 0xea, 0x92, 0x56, 0x19, 0x03, 0x66, 0x18, 0x6c, 0x54, 0xe9, 0x15, 0x1c, 0xaa,
	0x34, 0x47, 0xa1, 0x87, 0xea, 0xaa, 0xdb, 0x00, 0xaa, 0xac, 0xe1, 0xad, 0x56, 0x12, 0xb9, 0xb2,
	0xee, 0x3a, 0xa5, 0xd0, 0x9d, 0x88, 0xf5, 0x4d, 0xc0, 0x73, 0x6c, 0x0f, 0x31, 0x3b, 0xf2, 0x1f,
	0xb7, 0xb0, 0x06, 0xea, 0x56, 0x79, 0xa4, 0x29, 0xeb, 0xf4, 0xc3, 0x65, 0xc0, 0x09, 0x0c, 0x9e,
	0x00, 0xf0, 0x61, 0xf2, 0x88, 0x89, 0x80, 0xe3, 0x00, 0x67, 0xfe, 0x4d, 0xe4, 0x46, 0x1b, 0x25,
	0x59, 0x85, 0x50, 0x88, 0x56, 0xb8, 0x28, 0xa3, 0x8c, 0xba, 0x52, 0x37, 0x55, 0xde, 0xcd, 0x98,
	0xe4, 0xe8, 0xaa, 0x5b, 0x3e, 0x0e, 0xb0, 0x56, 0x9c, 0xc7, 0x31, 0xd6, 0x07, 0x77, 0xd0, 0xde,
	0xfe, 0x21, 0xa2, 0xff, 0x83, 0x8e, 0xf2, 0xad, 0xc5, 0x78, 0xac, 0x5b, 0x16, 0x29, 0xd1, 0x1e,
	0xbc, 0x52, 0xd0, 0x15, 0x33, 0x8d, 0x73, 0x67, 0x3e, 0xb2, 0xac, 0x2b, 0x93, 0x4d, 0x88, 0x56,
	0x24, 0x1b, 0xa6, 0xb3, 0xb0, 0x74, 0x46, 0xca, 0xb8, 0xa6, 0x0a, 0x5a, 0x18, 0x17, 0x86, 0x79,
	0xe5, 0xe8, 0x8c, 0x99, 0x8c, 0x54, 0x06, 0xbf, 0xc3, 0x8b, 0x67, 0x3f, 0x20, 0xf4, 0x15, 0x90,
	0x1c, 0x2a, 0x6e, 0x3b, 0x80, 0xfd, 0x1c, 0x1d, 0xcd, 0x98, 0x3e, 0x9a, 0x5c, 0x3b, 0xfa, 0x2f,
	0x53, 0xcb, 0x26, 0x1a, 0x7d, 0x0d, 0xff, 0xcf, 0x63, 0x53, 0xe3, 0xe3, 0x68, 0x36, 0x9d, 0x38,
	0xc6, 0xe8, 0x52, 0x27, 0x65, 0xfa, 0x0d, 0xbc, 0xfe, 0x22, 0x94, 0x2b, 0xad, 0x0c, 0xfe, 0x80,
	0x97, 0x5f, 0xf9, 0x52, 0x63, 0x69, 0x3b, 0xf0, 0x4e, 0xd1, 0x3b, 0x11, 0xc3, 0x74, 0x98, 0x69,
	0x5e, 0x12, 0xed, 0x6b, 0x91, 0xb4, 0xf6, 0x03, 0xd8, 0x7f, 0x16, 0xb1, 0x1d, 0xd9, 0x0c, 0x52,
	0xb9, 0xa9, 0xcb, 0x7f, 0x29, 0x3f, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xa4, 0x92, 0x7f,
	0xb6, 0x08, 0x00, 0x00,
}
