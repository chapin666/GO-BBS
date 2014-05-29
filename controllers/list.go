package controllers

import (
	m "git.oschina.net/scitc/bbs/models"
	"github.com/astaxie/beego"
	"strconv"
)

type ContentsList struct {
	SimpleController
}

//添加帖子
func (this *ContentsList) BBSAdd() {
	kinds, err := m.GetKindsList()
	if err != nil {
	}
	this.Data["kinds"] = kinds
	this.Layout = INDEX
	this.TplNames = "admin/bbsAdd.html"
}

func (this *ContentsList) PostAdd() {
	bbs := m.BBS{}
	this.Layout = INDEX

	kid, _ := this.GetInt("bbsKind")

	if err := this.ParseForm(&bbs); err != nil {
		beego.Error("发布帖子失败，表单错误")
	}

	bbs.AddAuthor = &m.User{UserId: 4}
	bbs.BBSKind = &m.BBSKinds{KindId: kid}

	beego.Info(bbs)
	count, err := m.AddBBS(&bbs)

	if err == nil && count > 0 {
		this.Redirect("/admin/publishList", 302)
	} else {
		this.TplNames = "admin/bbsAdd.html"
	}
}

//帖子内容
func (this *ContentsList) Get() {
	id, _ := this.GetInt("id")
	bbs, err := m.GetBBSById(id)
	if err != nil {
	}
	responses, count, err := m.GetResByBBSId(bbs.BBSId)
	if err != nil {
	}

	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		this.Data["show"], _ = strconv.Atoi(n)
	}

	this.Data["responses"] = responses
	this.Data["bbs"] = bbs
	this.Data["id"] = id
	this.Data["count"] = count
	//this.Data[INFO] = SUCCESS
	this.TplNames = "pageContents.html"
}

//帖子列表
func (this *ContentsList) BBSList() {
	this.Layout = INDEX
	this.TplNames = "admin/bbsList.html"
}

// 发布的帖子
func (this *ContentsList) PublishList() {
	u, ok := this.Data["user"].(m.User)
	if ok {

		count, _ := m.GetBBSCountByUser(u.UserId)
		pager := this.paginator(PAGE_SIZE, count)
		bbs, _, _ := m.GetBBSByUserIdOffset(u.UserId, pager.Offset(), PAGE_SIZE)
		this.Data["bbs"] = bbs
	} else {
		this.Data["bbs"] = nil
	}

	p := this.GetString("p")
	if p == "" {
		p = `1`
	}
	this.Data["p"] = p

	this.Layout = INDEX
	this.TplNames = "admin/bbsPublish.html"
}

// 回复的帖子
func (this *ContentsList) ResponseList() {
	u, ok := this.Data["user"].(m.User)
	if ok {
		count := m.GetUserResCount(u.UserId)
		pager := this.paginator(PAGE_SIZE, count)

		res, _ := m.GetResByUserId(u.UserId, pager.Offset(), PAGE_SIZE)
		this.Data["res"] = res
	} else {
		this.Data["res"] = nil
	}

	p := this.GetString("p")
	if p == "" {
		p = `1`
	}
	this.Data["p"] = p

	this.Layout = INDEX
	this.TplNames = "admin/bbsResponse.html"
}
