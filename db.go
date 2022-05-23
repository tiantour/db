package db

import (
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func New(Driver, Source string) {
	var err error
	db, err = sqlx.Connect(Driver, Source)
	if err != nil {
		log.Fatalf("connect db err: %v", err)
	}

	db.SetMaxIdleConns(64)
	db.SetMaxOpenConns(32)
	db.SetConnMaxLifetime(128 * time.Minute)

	db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
}
