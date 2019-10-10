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
	username := this.GetString("username")
	// password := this.GetString("password")
	data, err := services.NewAccountService().FindAccount(username)
	if err != nil {

	}
	logs.Info(data)
	// data.name
	this.jsonResult(0, "ok")
}
