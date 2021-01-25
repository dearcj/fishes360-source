package main

import (
	"database/sql"
	"github.com/dearcj/golangproj/analytics"
	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

type Database interface {
	SaveAccounts(accs map[string]*Account) error
	Register(login string, password string) error
	Login(login string, password string, address string) (error, bool, *Account)
	ConnectDB() (Database, error)
}

type TestDB struct {
}

func (d *TestDB) SaveAccounts(accs map[string]*Account) error {
	return nil
}

func (d *TestDB) Register(login string, password string) error { return nil }

func (d *TestDB) Login(login string, password string, address string) (error, bool, *Account) {
	account := CreateAccount(nil, "", "")
	return nil, true, account
}

func (d *TestDB) ConnectDB() (Database, error) {
	return d, nil
}

type db struct {
	db       *sql.DB
	mongo    analytics.AnalyticsInterface
	Accounts *mgo.Collection
}
