package models

import "time"

type Account struct {
	UserId     int64
	Name       string
	NickName   string
	Phone      string
	Email      string
	Password   string
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
