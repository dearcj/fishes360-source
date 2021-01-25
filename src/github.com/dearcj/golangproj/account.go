package main

import (
	"github.com/dearcj/golangproj/msutil"
	"github.com/dearcj/golangproj/network"
)

type UserInfo struct {
	Id      string
	Balance float64
}

type UserCredentials struct {
	*UserInfo
	Token string
	Puuid string
}

type Account struct {
	UserCredentials
}

func (a *Account) Reset() {
}

func (a *Account) GetAccountGeneral() *data.AccountGeneral {
	return &data.AccountGeneral{
		Money:    float32(a.Balance),
		Username: a.Id,
	}
}

func (a *Account) Insert(s *msutil.XServerDataMsg) {
	s.WriteToMsg().AccountGeneral = a.GetAccountGeneral()
}

func defaultAccount() *Account {
	acc := &Account{}

	return acc
}

func CreateAccount(info *UserInfo, token string, puuid string) *Account {
	var acc *Account = &Account{
		UserCredentials{
			info,
			token,
			puuid,
		},
	}

	//acc.CurrentLoot.Inventory = CreateInventory(BIG_INV_SIZE)
	return acc
}
