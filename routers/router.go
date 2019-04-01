package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers"
)

func init() {
	// admin
	beego.Router("/admin/", &controllers.UserController{}, "get,post:Login")
	// 用户
	beego.Router("/admin/member/", &controllers.UserController{}, "post:Create;put:OpAll")
	beego.Router("/admin/member/:uid", &controllers.UserController{}, "put:Put;delete:Delete")
	// 目录
	beego.Router("/admin/maincate/", &controllers.MainCateController{})
	beego.Router("/admin/subcate/", &controllers.SubCateController{})
	// 登出
	beego.Router("/admin/logout/", &controllers.UserController{}, "*:Logout")
	beego.AutoRouter(&controllers.AdminController{})

	// 静态文件
	beego.SetStaticPath("/static","static")
}

