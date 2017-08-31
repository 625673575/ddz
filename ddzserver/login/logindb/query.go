package logindb

import (
	"log"
	"database/sql"
	"fmt"
	"ddzserver/message"
)
const TABLE_LOGIN_USER="loginuser"
const TABLE_USER_INFO="userinfo"
const COLUMN_USERNAME="username"
const COLUMN_PASSWORD="password"
const COLUMN_USERID="userid"
var SELECT_LOGINUSER_BY_USERNAME string
var SELECT_LOGINUSER_ALL string
var SELECT_USERINFO_BY_ID string
func init(){
	SELECT_LOGINUSER_BY_USERNAME =fmt.Sprintf("SELECT * FROM %s WHERE %s=?",TABLE_LOGIN_USER,COLUMN_USERNAME)
	SELECT_LOGINUSER_ALL =fmt.Sprintf("SELECT * FROM %s",TABLE_LOGIN_USER)
	SELECT_USERINFO_BY_ID=fmt.Sprintf("SELECT * FROM %s WHERE %s=?",TABLE_USER_INFO,COLUMN_USERID)
}
type LoginUser struct {
	Username string
	Password string
	Userid int64
}
func QueryLoginUser(UserName string)(bool,LoginUser){
	row:=db.QueryRow(SELECT_LOGINUSER_BY_USERNAME,UserName)
	var username string
	var password string
	var userid sql.NullInt64
	err:=row.Scan(&username,&password,&userid)
	hasUser:=err==nil
	if(!hasUser){
		log.Println("没有找到相关的用户",err.Error())
	}
	user:=LoginUser{Username:username,Password:password,Userid:userid.Int64}
	return hasUser,user
}

func QueryLoginUserTable()[]LoginUser{
	rows, err := db.Query(SELECT_LOGINUSER_ALL)
	if err != nil {
		log.Fatal(err)
	}
	userList:=make([]LoginUser,0)
	for rows.Next() {
		var username string
		var password string
		var userid sql.NullInt64
		if err := rows.Scan(&username,&password,&userid); err != nil {
			log.Fatal(err)
		}
		user:=LoginUser{Username:username,Password:password,Userid:userid.Int64}
		userList=append(userList,user)
	}
	return userList
}


func QueryUserInfo(userid int64) (bool,message.UserInfo){
	row:=db.QueryRow(SELECT_USERINFO_BY_ID,userid)
	var username string
var rank sql.NullInt64
	err:=row.Scan(&userid,&username,&rank)
	hasUser:=err==nil
	if !hasUser{
		log.Println("没有找到相关的用户",err.Error())
	}
	userbrief:=&message.UserBriefInfo{Userid:userid,UserName:username,UserIcon:1,Rank:int32(rank.Int64)}
	user:=message.UserInfo{BasicInfo:userbrief}
	return hasUser,user
}