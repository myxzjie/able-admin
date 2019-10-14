package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/myxzjie/able-admin/services"
)

type DemoController struct {
	BaseController
}

func (this *DemoController) Get() {
	data, err := services.NewDemoService().FindByIdentify(1)
	if err != nil {
		if err == orm.ErrNoRows {
			logs.Error(" not data")
		}
		logs.Error(" get data err")
	}
	logs.Info(data)
	this.Data["IsDemo"] = true
	this.TplName = viewTemplate("demo")
}
