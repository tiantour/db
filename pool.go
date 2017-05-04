package db

import (
	"runtime"

	"github.com/jmoiron/sqlx"
)

var (
	po   *pool
	conn chan int
)

// pool pool
type pool struct {
	sqlx.DB
}

// newPool new pool
func newPool() {
	cap := runtime.NumCPU()
	conn = make(chan int, cap)
	for i := 0; i < cap; i++ {
		conn <- 1
	}
	po = &pool{*db}
}
