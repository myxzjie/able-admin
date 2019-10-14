package controllers

import (
	"github.com/myxzjie/able-admin/models"
	"github.com/myxzjie/able-admin/services"
)

type AccountController struct {
	BaseController
}

func (this *AccountController) Get() {
	this.Data["IsAccount"] = true
	this.TplName = viewTemplate("account/index")
}

func (this *AccountController) DataPage() {
	// showCount := 10
	// totalPage := 0
	page := models.DefaultPerson()
	page.CurrentPage = 1
	datas, total, _ := services.NewAccountService().GetDataListPage(page)

	// if val := (total % showCount); val == 0 {
	// 	totalPage = total / showCount
	// } else {
	// 	totalPage = (total / showCount) + 1
	// }

	page.TotalResult = total
	this.jsonResultPage(0, "ok", datas, page)
}
