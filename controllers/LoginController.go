package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/myxzjie/able-admin/models"
	"github.com/myxzjie/able-admin/services"
	"github.com/myxzjie/able-admin/utils"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Get() {
	this.TplName = viewTemplate("login")
}

func (this *LoginController) Post() {

	// username := this.GetString("username")
	// password := this.GetString("password")
	m := models.Account{}
	if err := this.ParseForm(&m); err != nil {
		this.Data["massage"] = "parm error"
		// this.jsonResult(1, err.Error())
		// this.Redirect("/login?massage=error", 302)
		// this.StopRun()
		this.TplName = viewTemplate("login")
		return
	}
	logs.Debug("parm:", m)
	data, err := services.NewAccountService().AccountLogin(m.Name, m.Password)
	// logs.Debug(">>data", data)
	if err != nil {
		logs.Error(">> err", err, data)
		this.Data["massage"] = "error"
		// this.jsonResult(1, err.Error())
		// this.Redirect("/login?massage=error", 302)
		// this.StopRun()
		this.TplName = viewTemplate("login")
		return
	}

	this.SetSession(utils.SESSION_USER, *data)

	// this.jsonResult(0, "ok")
	this.Redirect("/", 302)
	this.StopRun()
}

func (this *LoginController) Logout() {
	// this.DelSession(utils.SESSION_USER)
	this.DestroySession()
	this.Redirect("/login", 302)
	this.StopRun()
}
