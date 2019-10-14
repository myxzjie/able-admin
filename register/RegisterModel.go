package register

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/myxzjie/able-admin/models"
	"github.com/myxzjie/able-admin/utils"
)

// RegisterModel 注册Model
func RegisterModel() {
	orm.RegisterModelWithPrefix(utils.DatabasePrefix,
		new(models.Account),
		new(models.Demo))
}

func GetDatabasePrefix() string {
	return beego.AppConfig.DefaultString("database::db_prefix", "")
}
