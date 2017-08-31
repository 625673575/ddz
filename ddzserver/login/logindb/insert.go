package logindb

import (
	"fmt"
)
var INSERT_LOGINUSER string

func init(){
	INSERT_LOGINUSER =fmt.Sprintf("insert into %s(%s,%s,%s) values(?,?,?)",TABLE_LOGIN_USER,COLUMN_USERNAME,COLUMN_PASSWORD,COLUMN_USERID )
}
func InsertUser(UserName string ,Password string ){
	db.Exec(INSERT_LOGINUSER,UserName,Password,0)
}

