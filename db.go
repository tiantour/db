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
	source := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8",
		c.Uname,
		c.Passwd,
		c.IP,
		c.Port,
		c.Database,
	)

	var err error
	db, err = sqlx.Open("mysql", source)
	if err != nil {
		log.Fatalf("open db err: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("ping db err: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
}
