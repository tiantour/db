package db

import (
	"fmt"
	"runtime"
	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tiantour/conf"
)

var (
	conn    chan int
	cap     = runtime.NumCPU()
	pool    = NewPool()
	db, err = sqlx.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@/%s",
			conf.Options.DB.Uname,
			conf.Options.DB.Passwd,
			conf.Options.DB.Database,
		),
	)
)

// Pool pool
type Pool struct {
	sqlx.DB
}

// NewPool newpool
func NewPool() *Pool {
	conn = make(chan int, cap)
	for i := 0; i < cap; i++ {
		conn <- 1
	}
	return &Pool{*db}
}

var (
	// Read Read
	Read = &read{}
	// Write Write
	Write = &write{}
)

type (
	write struct{}
	read  struct{}
)
