package logindb

import (
	"fmt"
)
var UPDATE_LOGINUSER_PASSWORD_BY_USERNAME string

func init(){
	UPDATE_LOGINUSER_PASSWORD_BY_USERNAME =fmt.Sprintf("update %s set %s=? where %s=?",TABLE_LOGIN_USER,COLUMN_PASSWORD, COLUMN_USERNAME)
}
func UpdateUser(UserName string ,Password string ){
	db.Exec(UPDATE_LOGINUSER_PASSWORD_BY_USERNAME,Password,UserName)
}
