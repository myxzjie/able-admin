package register

import (
	"github.com/myxzjie/go-cms/database"
)

var (
	ConfigurationFile = "./conf/app.conf"
	WorkingDirectory  = "./"
	LogFile           = "./logs"
)

func InitRegister() {
	database.RegisterDataBase()
}
