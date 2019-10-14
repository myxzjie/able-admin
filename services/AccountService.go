package services

import (
	"errors"

	"github.com/astaxie/beego/logs"
	"github.com/myxzjie/able-admin/models"
	"github.com/myxzjie/able-admin/utils"
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
	// logs.Debug(">>", data)
	if err != nil {
		logs.Error(">>", err)
		// e = errors.New("data error")
		return nil, err
	}
	if (verify(*data, password)) == false {
		return nil, errors.New("password error")
	}
	// res = data
	// return
	return data, nil
}

func password(m models.Account) models.Account {
	salt := utils.SecureRandomGenerator()
	salt = m.Name + salt
	npass := utils.SimpleHash(utils.AlgorithmName, m.Password, salt, utils.HashIterations)

	m.Salt = salt
	m.Password = npass
	return m
}

func verify(m models.Account, pass string) bool {
	salt := m.Salt
	salt = m.Name + salt
	logs.Debug(">>salt:", salt, pass)
	npass := utils.SimpleHash(utils.AlgorithmName, pass, salt, utils.HashIterations)
	logs.Debug(">> verify password:", npass, m.Password)
	if npass == m.Password {
		return true
	}
	return false
}
