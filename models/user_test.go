package models

import (
	"fmt"
	"testing"
)

func Test_AddUser(t *testing.T) {
	/*u := User{}
	u.UserName = "chegbin"
	u.Password = "1234566"
	u.Email = "997155658@qq.com"
	u.PicUrl = "122"
	u.Sex = "男"
	u.UserType = "系统管理员"

	id, err := AddUser(&u)

	if err == nil && id > 0 {
		fmt.Println(id)
	}*/
}

func Test_Query(t *testing.T) {
	params, _ := GetUserList(1, 12, "id")
	for _, m := range params {
		fmt.Println(m["UserId"], m["UserName"])
	}

}

func TestQuery(t *testing.T) {
	user := FindUserById(23)
	fmt.Println(user)
}
