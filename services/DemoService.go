package services

import (
	"github.com/myxzjie/able-admin/models"
)

type DemoService struct {
	BaseService
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (m *DemoService) FindByIdentify(id int, cols ...string) (result *models.Demo, err error) {
	// o := orm.NewOrm()
	result = &models.Demo{}
	err = getOrm().QueryTable("demo").Filter("id", id).One(result)
	return
}
