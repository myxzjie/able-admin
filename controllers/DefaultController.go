package controllers

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["IsDashboard"] = true
	// this.Data["Email"] = "astaxie@gmail.com"
	// logs.Info(">>" + viewTemplate("index"))
	this.TplName = viewTemplate("index")
}
