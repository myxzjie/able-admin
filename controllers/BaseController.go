package controllers

import (
	"github.com/astaxie/beego"
	"github.com/myxzjie/go-cms/models"
)

type BaseController struct {
	beego.Controller
}

var (
	M = new(models.M)
)
