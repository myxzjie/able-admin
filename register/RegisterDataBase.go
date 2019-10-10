package register

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// RegisterDataBase 注册数据库
func RegisterDataBase() {
	driver := beego.AppConfig.String("database::db_driver")
	logs.Debug(">> driver:"+driver)
	switch driver {
	case "mysql":
		mysql()
	case "sqlite3":
		sqlite3()
	}

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
	
}

func mysql(){
	alias := beego.AppConfig.String("mysql::db_alias")
	host := beego.AppConfig.String("mysql::db_host")
	port := beego.AppConfig.String("mysql::db_port")
	database := beego.AppConfig.String("mysql::db_database")
	username := beego.AppConfig.String("mysql::db_username")
	password := beego.AppConfig.String("mysql::db_password")
	charset := beego.AppConfig.String("mysql::db_charset")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local", username, password, host, port, database, charset)
	
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(alias, "mysql", dataSource)
}

func sqlite3(){
	alias := beego.AppConfig.String("sqlite3::db_alias")
	database := beego.AppConfig.String("sqlite3::db_database")

	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase(alias, "sqlite3", database)
}

