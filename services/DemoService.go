package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/myxzjie/go-cms/models"
)

type DemoService struct {
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (m *DemoService) FindByIdentify(id int, cols ...string) (result *models.Demo, err error) {
	o := orm.NewOrm()
	result = &models.Demo{}
	err = o.QueryTable("demo").Filter("id", id).One(result)
	return
}
