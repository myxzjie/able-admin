package services

import "github.com/astaxie/beego/orm"

type BaseService struct{}

func getOrm() orm.Ormer {
	o := orm.NewOrm()
	return o
}
