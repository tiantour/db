package db

import (
	"fmt"
	"os"
	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tiantour/conf"
)

var (
	db  *sqlx.DB
	err error
)

func init() {
	newServe()
	newPool()
}

func newServe() {
	account := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		conf.NewConf().DB.Uname,
		conf.NewConf().DB.Passwd,
		conf.NewConf().DB.Host,
		conf.NewConf().DB.Port,
		conf.NewConf().DB.Database,
	)
	db, err = sqlx.Open("mysql", account)
	if err != nil {
		os.Exit(1)
	}
}
