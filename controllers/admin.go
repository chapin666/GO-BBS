package controllers

import (
	m "git.oschina.net/scitc/bbs/models"
)

type AdminController struct {
	SimpleController
}

func (this *AdminController) Get() {
	u := this.GetSession("user")
	if u != nil {
		user, ok := u.(m.User)
		if ok && user.UserName != "" {
			this.Data["user"] = user
		} else {
			this.Data["user"] = nil
		}
	} else {
		this.Data["user"] = nil
	}

	this.Layout = INDEX
	this.TplNames = "admin/personalInfo.html"
}

func (this *AdminController) Logout() {
	this.DelSession("user")
	this.DestroySession()
	this.Redirect("/", 302)
}
