package models

import "github.com/astaxie/beego/orm"

type Demo struct {
	Id   int
	Age  int
	Name string
}

func NewDemo() *Demo {
	return &Demo{}
}

func (m *Demo) FindByIdentify(id int, cols ...string) (result *Demo, err error) {
	o := orm.NewOrm()
	result = &Demo{}
	err = o.QueryTable("demo").Filter("id", id).One(result)
	return
}
