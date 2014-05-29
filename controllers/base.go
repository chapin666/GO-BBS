package controllers

import (
	"fmt"
	m "git.oschina.net/scitc/bbs/models"
	"github.com/astaxie/beego"
	"path"
	"reflect"
	"strconv"
	"time"
)

const (
	DEFAULT = iota
	SUCCESS
	FAILED
)

const (
	INDEX        = "admin/index.html"
	INFO         = "info"
	CURRENT_PAGE = 1
	PAGE_SIZE    = 12
	PAGE_COUNT   = 0
	UPLOAD_FILE  = "/file"
)

type SimpleController struct {
	beego.Controller
}

func (this *SimpleController) paginator(per int, nums int64) *m.Paginator {
	p := m.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

type Info struct {
	name string
}

func (this *SimpleController) Prepare() {
	u := this.GetSession("user")

	if u != nil {
		user, ok := u.(m.User)
		if ok {
			this.Data["user"] = user
			this.Data["info"] = SUCCESS
			if user.UserType == "系统管理员" {
				this.Data["type"] = 1
			}
		} else {
			this.Data["info"] = FAILED
		}
	} else {
		this.Data["info"] = FAILED
	}
}

func (this *SimpleController) uploadFile(hd interface{}, formName string) string {
	if hd != nil {
		timestap := time.Now().UnixNano()
		filepath := path.Join("static/files", strconv.FormatInt(timestap, 10)+".jpg")
		err := this.SaveToFile(formName, filepath)
		if err != nil {
			beego.Error("upload file error:", err.Error())
		}
		return filepath
	}
	return ""
}

//////util/////////////////
func ToInt64(value interface{}) (d int64, err error) {
	val := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		d = val.Int()
	case uint, uint8, uint16, uint32, uint64:
		d = int64(val.Uint())
	default:
		err = fmt.Errorf("ToInt64 need numeric not `%T`", value)
	}
	return
}
