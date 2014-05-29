package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	//register database driver
	orm.RegisterDriver("mysql", orm.DR_MySQL)

	//register database
	orm.RegisterDataBase("default", "mysql", "root:root@/bbs?charset=utf8&loc=Asia%2FShanghai", 30)
}

func Orm() orm.Ormer {
	o := orm.NewOrm()
	o.Using("default")
	return o
}
