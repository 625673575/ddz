package logindb

import (
	"fmt"
)
var INSERT_LOGINUSER string
var INSERT_USERINFO string
func init(){
	INSERT_LOGINUSER =fmt.Sprintf("insert into %s(%s,%s,%s) values(?,?,?)",TABLE_LOGIN_USER,COLUMN_USERNAME,COLUMN_PASSWORD,COLUMN_USERID )
	INSERT_USERINFO=fmt.Sprintf("insert into %s(%s,%s,%s) values(?,?,?)",TABLE_USER_INFO,COLUMN_USERID,COLUMN_USERNAME,COLUMN_ICONID )
}
func InsertLoginUser(UserName string ,Password string ){
	db.Exec(INSERT_LOGINUSER,UserName,Password,0)
}

func InsertUserInfo(Userid int64,UserName string ,IconId int64 ){
	db.Exec(INSERT_USERINFO,Userid,UserName,IconId)
}