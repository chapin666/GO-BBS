package routers

import (
	"git.oschina.net/scitc/bbs/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	// UserController
	beego.Router("/register", &controllers.UserController{}, "get:Get;post:Register")
	beego.Router("/login", &controllers.UserController{}, "post:Login")
	beego.Router("/admin/userList", &controllers.UserController{}, "get:UserList")
	beego.Router("/admin/userAdd", &controllers.UserController{}, "get:UserAdd;post:PostAdd")
	beego.Router("/admin/userUpdate", &controllers.UserController{}, "get:UserUpdate")
	beego.Router("/admin/userDelete", &controllers.UserController{}, "get:UserDelete")

	// KindsController{}
	beego.Router("/page", &controllers.KindsController{})
	beego.Router("/admin/kindsList", &controllers.KindsController{}, "get:KindsList")
	beego.Router("/admin/kindsAdd", &controllers.KindsController{}, "get:KindsAdd;post:PostAdd")
	beego.Router("/admin/kindsDel", &controllers.KindsController{}, "get:KindsDel")

	// ContentsList
	beego.Router("/list", &controllers.ContentsList{})
	beego.Router("/admin/bbsAdd", &controllers.ContentsList{}, "get:BBSAdd;post:PostAdd")
	beego.Router("/admin/bbsList", &controllers.ContentsList{}, "get:BBSList")
	beego.Router("/admin/publishList", &controllers.ContentsList{}, "get:PublishList")
	beego.Router("/admin/responseList", &controllers.ContentsList{}, "get:ResponseList")

	//ResponseController
	beego.Router("/response", &controllers.ResController{}, "post:PostAdd")

	// AdminController
	beego.Router("/admin/userInfo", &controllers.AdminController{})
	beego.Router("/admin/logout", &controllers.AdminController{}, "get:Logout")
}
