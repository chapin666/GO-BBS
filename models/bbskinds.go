package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

type BBSKinds struct {
	KindId    int64     `orm:"auto;column(id)" form:"-"`
	KindName  string    `orm:"column(kindName)" valid:"MinSize(2);MaxSize(18)" form:"kindName"`
	Kindintro string    `orm:"column(kindintro)" form:"kindIntro"`
	PicPath   string    `orm:"column(picPath)" form="-"`
	AddTime   time.Time `orm:"auto_now_add;type(datetime);column(addTime)" form="-"`
	AddAuthor string    `orm:"column(addAuthor)" form="-"`
}

func (this *BBSKinds) TableName() string {
	return "bbsKinds"
}

func init() {
	orm.RegisterModel(new(BBSKinds))
}

func (this BBSKinds) Valid(v *validation.Validation) {
	beego.Info("检验成功")
}

func AddKind(kinds *BBSKinds) (rows int64, err error) {
	o := Orm()
	rows, err = o.Insert(kinds)
	return
}

func GetKindsCount() (int64, error) {
	bbs := new(BBSKinds)
	o := Orm()
	qs := o.QueryTable(bbs)
	num, err := qs.Count()
	return num, err
}

func GetKindsList() ([]*BBSKinds, error) {
	var kinds []*BBSKinds
	kind := new(BBSKinds)
	o := Orm()
	_, err := o.QueryTable(kind).All(&kinds)

	if err != nil {
		return nil, err
	}

	return kinds, nil
}

func GetKindsListOffset(offset int, page_size int64) ([]*BBSKinds, error) {
	var kinds []*BBSKinds
	kind := new(BBSKinds)
	o := Orm()
	_, err := o.QueryTable(kind).Limit(page_size, offset).All(&kinds)

	if err != nil {
		return nil, err
	}

	return kinds, nil
}

func DelById(id int64) (int64, error) {
	o := Orm()
	num, err := o.Delete(&BBSKinds{KindId: id})

	return num, err
}
