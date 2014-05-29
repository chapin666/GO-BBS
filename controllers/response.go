package controllers

import (
	m "git.oschina.net/scitc/bbs/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type ResController struct {
	SimpleController
}

func (this *ResController) PostAdd() {
	response := m.Response{}

	id, err := this.GetInt("id")

	u, ok := this.Data["user"].(m.User)
	if !ok {
		flash := beego.NewFlash()
		flash.Store(&this.Controller)
		flash.Notice("1")
		flash.Store(&this.Controller)
		this.Redirect("/list?id="+strconv.FormatInt(id, 10), 302)
		return
	}
	response.AddAuthor = &m.User{UserId: u.UserId}

	//ID Error
	if err != nil {

	}

	// Form Error
	if err := this.ParseForm(&response); err != nil {

	}

	valid := validation.Validation{}
	b, err := valid.Valid(&response)

	// Validate Error
	if err != nil {
	}

	// Validate not Fomart
	if b {

		response.BBSId = &m.BBS{BBSId: id}

		count, err := m.AddRes(&response)

		//Add Error
		if err != nil {

		}

		if count > 0 && err == nil {
			this.Redirect("/list?id="+strconv.FormatInt(id, 10), 302)
		} else {

		}
	}
}
