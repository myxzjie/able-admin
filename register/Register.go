package register

import (
	"log"
	"os"

	"github.com/astaxie/beego"
)

var (
	ConfigurationFile = "./conf/app.conf"
	WorkingDirectory  = "./"
	LogFile           = "./logs"
)

func InitRegister() {
	err := beego.LoadAppConfig("ini", ConfigurationFile)
	if err != nil {
		log.Println("An error occurred:", err)
		os.Exit(1)
	}
	RegisterDataBase()
	RegisterModel()
	RegisterLogger(LogFile)
}
