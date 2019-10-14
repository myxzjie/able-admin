package controllers

type AccountController struct {
	BaseController
}

func (this *AccountController) Get() {

	this.Data["IsAccount"] = true
	this.TplName = viewTemplate("account/index")
}
