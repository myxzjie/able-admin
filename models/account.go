package models

import "time"

type Account struct {
	Id         int64
	Name       string `form:"username"`
	NickName   string
	Phone      string
	Email      string
	Password   string `form:"password"`
	Salt       string
	State      string
	CreateDate time.Time
	Remarks    string
	Stype      string
	Sex        string
	Card       string
	Birtn      time.Time
	CreateUser string
	Locked     string
	OrgId      string
	Avatar     string
}

func NewAccount() *Account {
	return &Account{}
}

// TableName 获取对应数据库表名.
// func (m *Account) TableName() string {
// 	return new(Account).(string)
// }
