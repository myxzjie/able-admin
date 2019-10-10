package services

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/myxzjie/able-admin/models"
)

type AccountService struct {
	BaseService
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (m *AccountService) FindAccount(name string) (res *models.Account, err error) {
	res = &models.Account{}
	err = getOrm().QueryTable(new(models.Account)).Filter("name", name).One(res)
	return
}

func (m *AccountService) AccountLogin(name string, password string) (res *models.Account, e error) {

	data, err := m.FindAccount(name)
	if err != nil {
		logs.Error(">> errr", err)
	}
	if data.Password != password {
		e = errors.New("password error")
	}
	res = data
	return

}
