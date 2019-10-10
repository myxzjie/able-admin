package services

import "github.com/myxzjie/able-admin/models"

type AccountService struct {
	BaseService
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (m *AccountService) FindAccount(name string) (res *models.Account, err error) {
	err = getOrm().QueryTable("able_account").Filter("name", name).One(res)
	return
}
