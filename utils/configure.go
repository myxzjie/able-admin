package utils

import "github.com/astaxie/beego"

var (
	Suffix       = ".html"
	SESSION_USER = "ABLE_SESSION_USER"
)

var (
	SessionName    = beego.AppConfig.DefaultString("sessionname", "sid")
	DatabasePrefix = beego.AppConfig.DefaultString("database::db_prefix", "")
)

const (
	secureRandomGeneratorSize = 16
	AlgorithmName             = "md5"
	HashIterations            = 2
)
