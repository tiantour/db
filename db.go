package db

import (
	"fmt"
	"log"
	"strings"
	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var (
	// IP ip
	IP = "127.0.0.1"
	// Port port
	Port = "3306"
	// Uname uname
	Uname = ""
	// Passwd passwd
	Passwd = ""
	// DB db
	DB  = ""
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
		Uname,
		Passwd,
		IP,
		Port,
		DB,
	)
	db, err = sqlx.Open("mysql", account)
	if err != nil {
		log.Fatal(err)
	}
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
}
