package db

import (
	"database/sql"
)

// Write write
type Write struct{}

// NewWrite new *Write
func NewWrite() *Write {
	return &Write{}
}

// List write list
func (w *Write) List(query string, args ...[]interface{}) (sql.Result, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := db.Preparex(db.Rebind(query))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var result sql.Result
	for _, v := range args {
		result, err = stmt.Exec(v...)
		if err != nil {
			return nil, err
		}
	}
	return result, tx.Commit()
}

// ListNamed write list named
func (w *Write) ListNamed(query string, args ...interface{}) (sql.Result, error) {
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := db.PrepareNamed(db.Rebind(query))
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var result sql.Result
	for _, v := range args {
		result, err = stmt.Exec(v)
		if err != nil {
			return nil, err
		}
	}
	return result, tx.Commit()
}

// Item write
func (w *Write) Item(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.Preparex(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

// ItemNamed wirte item named
func (w *Write) ItemNamed(query string, args interface{}) (sql.Result, error) {
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Exec(args)
}
