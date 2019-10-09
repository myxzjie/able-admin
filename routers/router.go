package routers

import (
	"github.com/astaxie/beego"
	"github.com/myxzjie/go-cms/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})
}
