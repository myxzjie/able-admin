package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/myxzjie/able-admin/register"
	_ "github.com/myxzjie/able-admin/routers"
)

func main() {
	logs.Info(">> able-admin server starting ....")
	register.InitRegister()
	// beego.Informational("Started sending signals.")
	beego.Run()

}
