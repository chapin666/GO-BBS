package main

import (
	m "git.oschina.net/scitc/bbs/models"
	_ "git.oschina.net/scitc/bbs/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func Index(in int) (out int) {
	out = in + 1
	return
}

func GetCount(id int64) (count int64) {
	count = m.GetBBSCountById(id)
	return
}

func GetResCount(id int64) (count int64) {
	count = m.GetResCount(id)
	return
}

var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("user").(m.User)
	if !ok && ctx.Request.RequestURI != "/" {
		ctx.Redirect(302, "/")
	}
}

func main() {
	beego.SessionOn = true
	beego.AddFuncMap("Index", Index)
	beego.AddFuncMap("Count", GetCount)
	beego.AddFuncMap("ResCount", GetResCount)
	beego.AddFilter("/admin/:all", "AfterStatic", FilterUser)
	beego.Run()
}
