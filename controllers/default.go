package controllers

import (
	m "git.oschina.net/scitc/bbs/models"
	"github.com/astaxie/beego"
	"strconv"
)

type MainController struct {
	SimpleController
}

func (this *MainController) Get() {
	u := this.GetSession("user")
	if u != nil {
		usr, ok := u.(m.User)
		if ok && usr.UserName != "" {
			this.Data[INFO] = SUCCESS
			this.Data["user"] = usr
		} else {
			this.Data[INFO] = DEFAULT
		}
	} else {
		this.Data[INFO] = DEFAULT
	}

	datas, err := m.GetKindsList()

	if err != nil {
	}
	this.Data["kinds"] = datas
	this.Data["status"] = 0

	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		this.Data["status"], _ = strconv.Atoi(n)
	}

	this.TplNames = "index.html"
}
