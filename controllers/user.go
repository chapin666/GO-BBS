package controllers

import (
	m "git.oschina.net/scitc/bbs/models"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
)

type UserController struct {
	SimpleController
}

func (this *UserController) getP(str string) string {
	count := m.GetUserCount()

	pager := this.paginator(PAGE_SIZE, count)
	total := pager.PageNums()

	cur, _ := strconv.Atoi(str)

	//如果当获取的页数大于总页数
	if cur > total {
		str = strconv.Itoa(total)
	}

	return str
}

func (this *UserController) Get() {
	this.TplNames = "register.html"
}

// user login
func (this *UserController) Login() {
	flash := beego.ReadFromRequest(&this.Controller)
	username := this.Input().Get("userName")
	password := this.Input().Get("password")

	if user := m.ValidateUser(username, password); user.UserName != "" {
		this.SetSession("user", user)
		this.Redirect("/", 302)
	} else {
		this.Data[INFO] = FAILED
		flash.Store(&this.Controller)
		flash.Notice("1")
		flash.Store(&this.Controller)
		this.Redirect("/", 302)
	}
}

func (this *UserController) Register() {
	beego.Info("Register()")
	u := m.User{UserType: "普通用户"}

	if err := this.ParseForm(&u); err != nil {
		beego.Debug(err.Error())
		this.Data[INFO] = FAILED
	}

	//如果用户名存在
	username := strings.Trim(u.UserName, " ")
	exists := m.FindUserByUserName(username)
	if exists {
		this.Data[INFO] = "对不起，用户名已经存在"
		this.TplNames = "register.html"
		return
	}

	_, hd, err := this.GetFile("photo")
	if err != nil {
		beego.Error("no photo")
	}

	fileName := this.uploadFile(hd, "photo")
	if fileName != "" {
		u.PicUrl = fileName
	}

	id, err := m.AddUser(&u)
	if err == nil && id > 0 {
		beego.Info("register success!!")
		this.Data[INFO] = SUCCESS
	} else {
		beego.Info("register failed!!")
		this.Data[INFO] = FAILED
	}

	this.Data["User"] = m.FindUserById(id)
	this.Layout = INDEX
	this.TplNames = "admin/personalInfo.html"
}

// user list
func (this *UserController) UserList() {

	// paginator
	count := m.GetUserCount()
	pager := this.paginator(PAGE_SIZE, count)

	users, _ := m.GetUserList(pager.Offset(), PAGE_SIZE, "userId")
	this.Data["Users"] = &users

	p := this.GetString("p")
	if p == "" {
		p = `1`
	}
	this.Data["p"] = p

	this.Layout = INDEX
	this.TplNames = "admin/userList.html"
}

// add user menu
func (this *UserController) UserAdd() {
	this.Data[INFO] = DEFAULT
	this.Layout = INDEX
	this.TplNames = "admin/userAdd.html"
}

// add user submit
func (this *UserController) PostAdd() {

	u := m.User{UserType: "系统管理员"}
	if err := this.ParseForm(&u); err != nil {
		this.Data[INFO] = FAILED
	}

	id, err := m.AddUser(&u)

	if err == nil && id > 0 {
		beego.Info("register success!!")
		this.Data[INFO] = SUCCESS
	} else {
		beego.Info("register failed!!", err.Error())
		this.Data[INFO] = FAILED
	}

	this.Layout = INDEX
	this.TplNames = "admin/userAdd.html"
}

// delete user from database
func (this *UserController) UserDelete() {
	id, err := this.GetInt("id")
	p := this.GetString("p")

	if err != nil {
		this.Data[INFO] = FAILED
	} else {
		status, err := m.DeletUserById(id)

		if err == nil && status > 0 {
			this.Data[INFO] = SUCCESS
		} else {
			this.Data[INFO] = FAILED
		}
	}

	p = this.getP(p)
	this.Redirect("/admin/userList?p="+p, 302)
}

//update user
func (this *UserController) UserUpdate() {
	this.Layout = INDEX
	this.TplNames = "admin/personalEdit.html"
}
