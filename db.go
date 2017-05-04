package db

import (
	"fmt"
	"log"
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
		"%s:%s@/%s",
		conf.Data.DB.Uname,
		conf.Data.DB.Passwd,
		conf.Data.DB.Database,
	)
	db, err = sqlx.Open("mysql", account)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
