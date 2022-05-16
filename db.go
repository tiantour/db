package db

import (
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/tiantour/conf"
)

var (
	err     error
	db      *sqlx.DB
	address = fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8",
		conf.NewDB().Data.Uname,
		conf.NewDB().Data.Passwd,
		conf.NewDB().Data.IP,
		conf.NewDB().Data.Port,
		conf.NewDB().Data.Database,
	)
)

func init() {
	db, err = sqlx.Connect("mysql", address)
	if err != nil {
		defer db.Close()
		log.Fatalf("connect db err: %v", err)
	}

	db.SetMaxIdleConns(64)
	db.SetMaxOpenConns(32)
	db.SetConnMaxLifetime(128 * time.Minute)

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
}
