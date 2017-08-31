package logindb

import (
	"fmt"
)
var DELETE_LOGINUSER_BY_USERNAME string
var DELETE_LOGINNUSER_BY_USERID string
func init(){
	DELETE_LOGINUSER_BY_USERNAME =fmt.Sprintf("DELETE FROM %s WHERE %s=?",TABLE_LOGIN_USER,COLUMN_USERNAME)
	DELETE_LOGINNUSER_BY_USERID =fmt.Sprintf("DELETE FROM %s WHERE %s=?",TABLE_LOGIN_USER,COLUMN_USERID)
}
func DeleteUserByUsername(UserName string ){
	db.Exec(DELETE_LOGINUSER_BY_USERNAME,UserName)
}

func DeleteUserByUserid(UserID string ){
	db.Exec(DELETE_LOGINNUSER_BY_USERID,UserID)
}
