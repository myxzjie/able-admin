package routers

import (
	"github.com/astaxie/beego"
	"github.com/myxzjie/able-admin/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})
	//
	beego.Router("/login", &controllers.LoginController{}, "get:Get")
	beego.Router("/login", &controllers.LoginController{}, "post:Post")
	beego.Router("/logout", &controllers.LoginController{}, "*:Logout")
	//
	beego.Router("/demo/index", &controllers.DemoController{}, "*:Get")
	//
	beego.Router("/account/index", &controllers.AccountController{}, "*:Get")
	beego.Router("/account/datapage", &controllers.AccountController{}, "*:DataPage")
}
