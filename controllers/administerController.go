package controllers

import "github.com/astaxie/beego/logs"

type AdministerController struct {
	BaseController
}

// Prepare 预处理.
func (this *AdministerController) Prepare() {
	this.BaseController.Prepare()
	logs.Debug(">> URI:", this.Ctx.Request.RequestURI)
	if data := this.GetAccountSession(); data.Id <= 0 {
		this.Redirect("/login", 302)
		this.StopRun()
	}
}
