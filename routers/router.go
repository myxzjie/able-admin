package routers

import (
	"github.com/astaxie/beego"
	"github.com/myxzjie/able-admin/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{}, "*:Get")
	beego.Router("/login/submit", &controllers.LoginController{}, "post:LoginSubmit")
}
