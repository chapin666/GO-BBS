package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

// user struct
type User struct {
	UserId     int64     `orm:"auto;column(id)"`
	UserName   string    `orm:"column(userName)" form:"userName" valid:"Required;MaxSize(20);MinSize(6)`
	Password   string    `orm:"column(password)" form:"password" valid:"Required;MaxSize(20);MinSize(6)"`
	RePassword string    `orm:"-" form:"passRepeat" valid:"Required"`
	Sex        string    `orm:"column(sex)" form:"sex" valid:"Required"`
	Email      string    `orm:"column(email)" form:"email" valid:"Email"`
	RegTime    time.Time `orm:"auto_now_add;type(datetime);column(regTime)"`
	UserType   string    `orm:"column(userType)"`
	PicUrl     string    `orm:"column(picUrl)"`
}

// Implements TableName.
// return table bame.
func (u *User) TableName() string {
	return "users"
}

// override init method.
func init() {
	orm.RegisterModel(new(User))
}

// checkUser function check form.
func checkUser(user *User) (err error) {
	vaild := validation.Validation{}

	// Valid parse the struct's tag,
	// then check the struct value.
	b, _ := vaild.Valid(&user)

	if !b {
		for _, err := range vaild.Errors {
			return errors.New(err.Message)
		}
	}

	return nil
}

/////////////crud/////////////////////
// resiter system user
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}

	o := Orm()
	rows, err := o.Insert(u)

	return rows, err
}

func GetUserCount() int64 {
	u := new(User)
	o := Orm()
	qs := o.QueryTable(u)
	num, _ := qs.Count()
	return num
}

// query from database of user list
func GetUserList(offset int, page_size int64, sort string) (users []orm.Params, count int64) {
	u := new(User)
	o := Orm()
	qs := o.QueryTable(u)

	count, err := qs.Limit(page_size, offset).OrderBy(sort).Values(&users)

	if err != nil {
		fmt.Print(err.Error())
		return nil, count
	}

	return
}

//delete from users
func DeletUserById(id int64) (status int64, err error) {
	o := Orm()
	status, err = o.Delete(&User{UserId: id})
	return
}

//find user by name
func FindUserByUserName(userName string) bool {
	u := new(User)
	o := Orm()
	return o.QueryTable(u).Filter("userName", userName).Exist()
}

func FindUserById(userId int64) (user User) {
	u := new(User)
	o := Orm()
	o.QueryTable(u).Filter("UserId", userId).One(&user)
	return
}

func ValidateUser(userName string, pwd string) (user User) {
	u := new(User)
	o := Orm()
	o.QueryTable(u).Filter("userName", userName).Filter("password", pwd).One(&user)
	return
}
