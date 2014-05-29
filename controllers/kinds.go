package controllers

import (
	m "git.oschina.net/scitc/bbs/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

type KindsController struct {
	SimpleController
}

func (this *KindsController) Get() {

	id, _ := this.GetInt("id")
	//根据类型获取BBS

	count, _ := m.GetBBSCount()
	pager := this.paginator(PAGE_SIZE, count)

	bbsList, err := m.GetBBSByKindOffset(id, pager.Offset(), PAGE_SIZE)
	if err != nil {
	}

	//获取BBS分类
	datas, err := m.GetKindsList()

	if err != nil {
	}

	p := this.GetString("p")
	if p == "" {
		p = `1`
	}
	this.Data["p"] = p

	this.Data["bbs"] = bbsList
	this.Data["id"] = id
	this.Data["kinds"] = datas
	this.TplNames = "pageList.html"
}

func (this *KindsController) KindsList() {

	count, _ := m.GetKindsCount()
	pager := this.paginator(PAGE_SIZE, count)

	datas, err := m.GetKindsListOffset(pager.Offset(), PAGE_SIZE)

	if err == nil {
		this.Data["count"] = 0
		this.Data["kinds"] = datas
	} else {
		this.Data["count"] = 0
		this.Data["kinds"] = nil
	}

	p := this.GetString("p")
	if p == "" {
		p = `1`
	}
	this.Data["p"] = p

	this.Layout = INDEX
	this.TplNames = "admin/bbsKindsList.html"
}

func (this *KindsController) KindsAdd() {
	this.Layout = INDEX
	this.Data[INFO] = -1
	this.TplNames = "admin/bbsKindsAdd.html"
}

func (this *KindsController) PostAdd() {
	kinds := m.BBSKinds{}
	this.Layout = INDEX

	if err := this.ParseForm(&kinds); err != nil {
		//this[INFO] = FAILED
		this.TplNames = "admin/bbsKindsAdd.html"
	}

	valid := validation.Validation{}
	b, err := valid.Valid(kinds)

	if err != nil {
		this.TplNames = "admin/bbsKindsAdd.html"
	}

	if b {
		_, hd, err := this.GetFile("kindLogo")
		if err != nil {
			beego.Error("no photo")
		}

		fileName := this.uploadFile(hd, "kindLogo")
		if fileName != "" {
			kinds.PicPath = fileName
		}
		kinds.AddAuthor = "Admin"

		beego.Info(kinds)

		id, err := m.AddKind(&kinds)
		if err == nil && id > 0 {
			beego.Info("add kinds success!!")
			this.TplNames = "admin/bbsKindsList.html"
		} else {
			beego.Info("add kinds failed!!")
			this.Data[INFO] = FAILED
			this.TplNames = "admin/bbsKindsAdd.html"
		}

	} else {
		for _, err := range valid.Errors {
			msg := err.Message
			if msg == "Minimum size is 2" {
				this.Data[INFO] = "帖子类型名称长度不能少用2位"
			} else if msg == "Maximum size is 18" {
				this.Data[INFO] = "帖子类型名称长度不能超过18位"
			}

			break
		}

		this.TplNames = "admin/bbsKindsAdd.html"
	}
}

func (this *KindsController) KindsDel() {
	id, err := this.GetInt("id")
	if err != nil {
	}
	_, err = m.DelById(id)

	if err != nil {

	} else {
		this.Redirect("/admin/kindsList", 302)
	}

}
