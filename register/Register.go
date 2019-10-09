package register

import (
	"github.com/myxzjie/go-cms/register"
)

var (
	ConfigurationFile = "./conf/app.conf"
	WorkingDirectory  = "./"
	LogFile           = "./logs"
)

func InitRegister() {
	register.RegisterDataBase()
}
