package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/myxzjie/able-admin/models"
	"github.com/myxzjie/able-admin/utils"
)

func init() {
	var FilterUser = func(ctx *context.Context) {
		data, ok := ctx.Input.Session(utils.SESSION_USER).(models.Account)
		logs.Debug(">> filter:", data, ok)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	}

	// var UrlManager = func(ctx *context.Context) {
	// 	// 数据库读取全部的 url mapping 数据
	// 	urlMapping := model.GetUrlMapping()
	// 	for baseurl,rule:=range urlMapping {
	// 		if baseurl == ctx.Request.RequestURI {
	// 			ctx.Input.RunController = rule.controller
	// 			ctx.Input.RunMethod = rule.method
	// 			break
	// 		}
	// 	}
	// }

	beego.InsertFilter("/", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/demo/*", beego.BeforeRouter, FilterUser)
	beego.InsertFilter("/account/*", beego.BeforeRouter, FilterUser)
}
