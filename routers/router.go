package routers

import (
	"github.com/astaxie/beego"
	"specialTady/controllers/admin"
	"specialTady/controllers/api"
)

func Init() {

	//APP段接口
	beego.AutoRouter(&api.UserController{})
	beego.AutoRouter(&api.ArticleController{})

	//后台
	beego.Router("/admin", &admin.IndexController{}, "*:Index")
	beego.Router("/admin/login", &admin.AccountController{}, "*:Login")
	beego.Router("/admin/logout", &admin.AccountController{}, "*:Logout")
	beego.Router("/admin/account/profile", &admin.AccountController{}, "*:Profile")

	//用户管理
	beego.Router("/admin/user/list", &admin.UserController{}, "*:List")
	beego.Router("/admin/user/add", &admin.UserController{}, "*:Add")
	beego.Router("/admin/user/edit", &admin.UserController{}, "*:Edit")
	beego.Router("/admin/user/delete", &admin.UserController{}, "*:Delete")

	//内容管理
	beego.Router("/admin/article/add", &admin.ArticleController{}, "*:Add")
	beego.Router("/admin/article/upload", &admin.ArticleController{}, "*:Upload")
	beego.Router("/admin/article/save", &admin.ArticleController{}, "post:Save")
	beego.Router("/admin/article/delete", &admin.ArticleController{}, "*:Delete")
	beego.Router("/admin/article/batch", &admin.ArticleController{}, "*:Batch")
	beego.Router("/admin/article/list", &admin.ArticleController{}, "*:List")
	beego.Router("/admin/article/trash", &admin.ArticleController{}, "*:Trash")
	beego.Router("/admin/article/edit", &admin.ArticleController{}, "*:Edit")
	beego.Router("/admin/article/destroy", &admin.ArticleController{}, "*:Destroy")
	beego.Router("/admin/tag/list", &admin.TagController{}, "*:List")
	beego.Router("/admin/tag/delete", &admin.TagController{}, "*:Delete")
	beego.Router("/admin/tag/add", &admin.TagController{}, "*:Add")
}
