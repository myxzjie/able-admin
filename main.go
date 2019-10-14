package main

import (
	"encoding/gob"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/astaxie/beego/session/mysql"
	_ "github.com/astaxie/beego/session/redis"
	"github.com/myxzjie/able-admin/models"
	"github.com/myxzjie/able-admin/register"
	_ "github.com/myxzjie/able-admin/routers"
)

func init() {
	gob.Register(models.Account{})
}

func main() {
	logs.Info(">> able-admin server starting ....")
	register.InitRegister()
	// beego.Informational("Started sending signals.")
	beego.Run()

}
