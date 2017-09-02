package main

import (
	"ddzserver/login/logindb"
	"log"
	"net"
	"github.com/alecthomas/log4go"
	"ddzserver/session"
)
var TEST_ENV=false

func initlog() {
	if TEST_ENV{
		return
	}
	log.Println("初始化 log4go")
	log4go.AddFilter("file", log4go.DEBUG, log4go.NewFileLogWriter("./temp/Logs/test.log", false)) //输出到文件,级别为DEBUG,文件名为test.log,每次追加该原文件
	log4go.AddFilter("file", log4go.ERROR, log4go.NewFileLogWriter("./temp/Logs/test.log", false)) //输出到文件,级别为ERROR,文件名为test.log,每次追加该原文件

	//l4g.LoadConfiguration("log.xml")//使用加载配置文件,类似与java的log4j.propertites
	log4go.Debug("start server")
	defer log4go.Close()
}



func main() {
	initlog()
	logindb.OpenDB()
	StartSocketListen()
}
func StartSocketListen() {

	service := ":5000"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)

	listener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listener.Accept()
		if conn == nil {
			log4go.Error(err.Error())
			continue
		}
		addr := conn.RemoteAddr().String()
		if addr == "" {
			continue
		}
		session.NewSession(&conn)
		log.Println("Receive connect request from ", conn.RemoteAddr().String())
		if err != nil {
			log4go.Error(err.Error())
			continue
		}
	}
}
