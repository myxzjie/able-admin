package database

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// RegisterDataBase 注册数据库
func RegisterDataBase() {

	driver := beego.AppConfig.String("db_driver")

	alias := beego.AppConfig.String("db_alias")

	host := beego.AppConfig.String("db_host")
	port := beego.AppConfig.String("db_port")

	database := beego.AppConfig.String("db_database")
	username := beego.AppConfig.String("db_username")
	password := beego.AppConfig.String("db_password")

	charset := beego.AppConfig.String("db_charset")

	switch driver {
	case "mysql":
		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, database, charset)

		orm.RegisterDataBase(alias, driver, dataSource)
	case "sqlite3":
		orm.RegisterDataBase(alias, driver, database)
	}

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

}

func MysqlDataSource() {

}
