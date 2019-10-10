package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/myxzjie/able-admin/services"
)

type MainController struct {
	BaseController
	// beego.Controller
}

func (this *MainController) Get() {
	// result := models.Demo{}
	// err := M.Object().QueryTable(new(models.Demo)).One(&result)
	// if err != nil {

	// }
	data, err := services.NewDemoService().FindByIdentify(1)
	if err != nil {
		if err == orm.ErrNoRows {
			logs.Error(" not data")
		}
		logs.Error(" get data err")
	}
	logs.Info(data)
	// models.NewDemo().FindByIdentify(1)
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	logs.Info(">>" + viewTemplate("index"))
	this.TplName = viewTemplate("index")
}
