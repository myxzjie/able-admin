package models

import "github.com/astaxie/beego/orm"

var (
	o orm.Ormer
)

func init() {
	// o = orm.NewOrm()
}

type M struct {
}

func (this *M) Object() orm.Ormer {
	return o
}
