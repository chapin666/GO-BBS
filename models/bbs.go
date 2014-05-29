package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
)

type BBS struct {
	BBSId      int64     `orm:"auto;column(id)" form:"-"`
	BBSTitle   string    `orm:"column(bbsTitle)" form:"bbsName"`
	BBSContent string    `orm:"column(bbsContent)" form:"content"`
	BBSKind    *BBSKinds `orm:"rel(fk);column(bbsKind)" form:"-"`
	AddTime    time.Time `orm:"auto_now_add;type(datetime);column(addTime)" form:"-"`
	AddAuthor  *User     `orm:"rel(fk);column(addAuthor)" form:"-"`
}

func (this *BBS) TableName() string {
	return "bbs"
}

func init() {
	orm.RegisterModel(new(BBS))
}

func (this *BBS) Valid(v *validation.Validation) {
	beego.Info("帖子检验成功")
}

func GetBBSCount() (int64, error) {
	bbs := new(BBS)
	o := Orm()
	qs := o.QueryTable(bbs)
	num, err := qs.Count()
	return num, err
}

func AddBBS(bbs *BBS) (rows int64, err error) {
	o := Orm()
	rows, err = o.Insert(bbs)
	return
}

func BBSList() ([]*BBS, error) {
	var bbsList []*BBS
	bbs := new(BBS)
	o := Orm()

	//Limit(page_size, offset).OrderBy(sort)
	_, err := o.QueryTable(bbs).RelatedSel().All(&bbsList)

	if err != nil {
		return nil, err
	}

	return bbsList, nil
}

func BBSListOffset(offset int, page_size int64, sort string) ([]*BBS, error) {
	var bbsList []*BBS
	bbs := new(BBS)
	o := Orm()

	_, err := o.QueryTable(bbs).RelatedSel().Limit(page_size, offset).OrderBy(sort).All(&bbsList)

	if err != nil {
		return nil, err
	}

	return bbsList, nil
}

func GetBBSByKind(kind int64) ([]*BBS, error) {
	var bbsList []*BBS
	bbs := new(BBS)
	o := Orm()

	_, err := o.QueryTable(bbs).Filter("BBSKind", kind).RelatedSel().All(&bbsList)

	if err != nil {
		return nil, err
	}

	return bbsList, nil
}

func GetBBSByKindOffset(kind int64, offset int, page_size int64) ([]*BBS, error) {
	var bbsList []*BBS
	bbs := new(BBS)
	o := Orm()

	_, err := o.QueryTable(bbs).Filter("BBSKind", kind).RelatedSel().Limit(page_size, offset).All(&bbsList)

	if err != nil {
		return nil, err
	}

	return bbsList, nil
}

func GetBBSById(id int64) (bbs BBS, err error) {
	b := new(BBS)
	o := Orm()
	err = o.QueryTable(b).Filter("BBSId", id).RelatedSel().One(&bbs)
	return
}

func GetBBSCountById(id int64) int64 {
	bbs := new(BBS)
	o := Orm()

	count, _ := o.QueryTable(bbs).Filter("BBSKind", id).Count()
	return count
}

func GetBBSByUserId(id int64) ([]*BBS, error) {
	var bbs []*BBS
	b := new(BBS)
	o := Orm()
	_, err := o.QueryTable(b).RelatedSel().Filter("AddAuthor__UserId", id).All(&bbs)

	if err != nil {
		return nil, err
	}

	return bbs, nil
}

func GetBBSByUserIdOffset(id int64, offset int, page_size int64) ([]*BBS, int64, error) {
	var bbs []*BBS
	b := new(BBS)
	o := Orm()
	seter := o.QueryTable(b).RelatedSel().Filter("AddAuthor__UserId", id)

	count, err := seter.Count()
	if err != nil {
		return nil, 0, err
	}

	_, err = seter.Limit(page_size, offset).All(&bbs)
	if err != nil {
		return nil, 0, err
	}

	return bbs, count, nil
}

func GetBBSCountByUser(id int64) (int64, error) {
	b := new(BBS)
	o := Orm()
	seter := o.QueryTable(b).RelatedSel().Filter("AddAuthor__UserId", id)

	count, err := seter.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}
