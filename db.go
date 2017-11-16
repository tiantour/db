package db

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

var (
	conn chan int
	po   *pool
)

// pool pool
type pool struct {
	sqlx.DB
}

// New new db
func New(ip, port, database, uname, passwd string) {
	if uname == "" || database == "" || passwd == "" {
		log.Fatal("db conf is null")
	}
	if ip == "" {
		ip = "127.0.0.1"
	}
	if port == "" {
		port = ":6379"
	}
	db, err := sqlx.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8",
			uname,
			passwd,
			ip,
			port,
			database,
		))
	if err != nil {
		log.Fatal(err)
	}
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	cap := runtime.NumCPU()
	conn = make(chan int, cap)
	for i := 0; i < cap; i++ {
		conn <- 1
	}
	po = &pool{*db}
}
