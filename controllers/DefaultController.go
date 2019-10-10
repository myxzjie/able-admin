package controllers

import (
	"github.com/myxzjie/go-cms/models"
)

type MainController struct {
	BaseController
	// beego.Controller
}

func (c *MainController) Get() {
	// result := models.Demo{}
	// err := M.Object().QueryTable(new(models.Demo)).One(&result)
	// if err != nil {

	// }
	models.NewDemo().FindByIdentify(1)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
