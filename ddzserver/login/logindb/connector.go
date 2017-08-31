package logindb

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import (
	"github.com/alecthomas/log4go"
)


var db *sql.DB
var err error
func OpenDB() {
	db, err = sql.Open("mysql", "root:root@/ddz")
	if err != nil{
		log4go.Error("数据库加载失败")
		return
	}
	db.Begin()
}