package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/myxzjie/able-admin/services"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {

	this.TplName = viewTemplate("login")
}

func (this *LoginController) LoginSubmit() {
	// username := this.GetString("username")
	password := this.GetString("password")
	data, err := services.NewAccountService().AccountLogin("demo", password)
	if err != nil {
		logs.Error(">> err", err)
		this.jsonResult(1, err.Error())
	}
	logs.Info(data.Name)
	this.jsonResult(0, "ok")
}
