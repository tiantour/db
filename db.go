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
	c := conf.NewDB().Data
	address := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8",
		c.Uname,
		c.Passwd,
		c.IP,
		c.Port,
		c.Database,
	)

	var err error
	db, err = sqlx.Connect("mysql", address)
	if err != nil {
		log.Fatalf("open db err: %v", err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(15 * time.Minute)
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
}
