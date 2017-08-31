package logindb

import (
	"fmt"
	"testing"
)

func Test_SQL(T *testing.T){
	OpenDB()
	DeleteUserByUsername("wei")
	UpdateUser("hong","tyty")
	tables:=QueryLoginUserTable()
	fmt.Println(tables)
	hasUser,hong:=QueryLoginUser("hong")
	fmt.Println(hasUser,hong)
	hasUser,hong=QueryLoginUser("hong")
	fmt.Println(hasUser,hong)
	hasUser,liu:=QueryUserInfo(2)
	fmt.Println(hasUser,liu)
}
