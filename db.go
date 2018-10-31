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
	"github.com/tiantour/conf"
)

var (
	conn chan int
	po   *pool
)

// pool pool
type pool struct {
	sqlx.DB
}

func init() {
	c := conf.NewDB().Data
	if c.Uname == "" || c.Database == "" || c.Passwd == "" {
		log.Fatal("db conf is null")
	}
	if c.IP == "" {
		c.IP = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = ":6379"
	}

	db, err := sqlx.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8",
			c.Uname,
			c.Passwd,
			c.IP,
			c.Port,
			c.Database,
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
