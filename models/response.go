package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

type Response struct {
	ResId     int64     `orm:"auto;column(id)" form:"-"`
	BBSId     *BBS      `orm:"rel(fk);column(bbsId)" form:"-"`
	Content   string    `orm:"column(content)" form:"content" valid:"Required;MinSize(1)"`
	AddTime   time.Time `orm:"auto_now_add;type(datetime);column(addTime)" form:"-"`
	AddAuthor *User     `orm:"rel(fk);column(addAuthor)" form:"-"`
}

func (this *Response) TableName() string {
	return "response"
}

func init() {
	orm.RegisterModel(new(Response))
}

func (this Response) Valid(v *validation.Validation) {
	beego.Info("回复表单校验成功")
}

func AddRes(res *Response) (rows int64, err error) {
	o := Orm()
	rows, err = o.Insert(res)
	return
}

func GetResCount(id int64) (count int64) {
	res := new(Response)
	o := Orm()
	count, _ = o.QueryTable(res).RelatedSel().Filter("BBSId__BBSId", id).Count()
	return
}

func GetResByBBSId(id int64) ([]*Response, int64, error) {
	var responses []*Response
	res := new(Response)
	o := Orm()
	seter := o.QueryTable(res).RelatedSel().Filter("BBSId__BBSId", id)

	count, err := seter.Count()
	if err != nil {
		return nil, 0, err
	}

	_, err = seter.All(&responses)
	if err != nil {
		return nil, 0, err
	}

	return responses, count, nil
}

func GetUserResCount(id int64) (count int64) {
	res := new(Response)
	o := Orm()
	count, _ = o.QueryTable(res).RelatedSel().Filter("AddAuthor__UserId", id).Count()
	return
}

func GetResByUserId(id int64, offset int, page_size int64) ([]*Response, error) {
	var responses []*Response
	res := new(Response)
	o := Orm()
	_, err := o.QueryTable(res).RelatedSel().Filter("AddAuthor__UserId", id).Limit(page_size, offset).All(&responses)

	if err != nil {
		return nil, err
	}

	return responses, nil
}
