package db

import (
	"fmt"
	"log"
	"strings"
	"time"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/tiantour/conf"
)

// db pool
var db *sqlx.DB

func init() {
	fmt.Println(1)
	c := conf.NewDB().Data
	fmt.Println(2, c)
	address := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8",
		c.Uname,
		c.Passwd,
		c.IP,
		c.Port,
		c.Database,
	)
	fmt.Println(3, address)

	var err error
	db, err = sqlx.Connect("mysql", address)
	fmt.Println(4, db, err)
	if err != nil {
		defer db.Close()
		log.Fatalf("open db err: %v", err)
	}

	err = db.Ping()
	fmt.Println(5, err)
	if err != nil {
		defer db.Close()
		log.Fatalf("ping db err: %v", err)
	}

	db.SetMaxIdleConns(120)
	db.SetMaxOpenConns(120)
	db.SetConnMaxLifetime(240 * time.Minute)
	fmt.Println(6)

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	fmt.Println(7)
}
