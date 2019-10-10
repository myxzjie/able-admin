package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/myxzjie/go-cms/register"
	_ "github.com/myxzjie/go-cms/routers"
)

func main() {
	logs.Info(">> go-cms server starting ....")
	register.InitRegister()
	// beego.Informational("Started sending signals.")
	beego.Run()

}
