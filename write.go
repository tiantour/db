package db

import "database/sql"

// Write write
type Write struct{}

// NewWrite new write
func NewWrite() *Write {
	return &Write{}
}

// Item write
func (w Write) Item(query string, args ...interface{}) (sql.Result, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}

// List write list
func (w Write) List(query string, args []interface{}) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	tx, err := po.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	argsLength := len(args)
	for i := 0; i < argsLength; i++ {
		_, err = stmt.Exec(args[i])
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// ItemNamed wirte item named
func (w Write) ItemNamed(query string, args interface{}) (sql.Result, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args)
}

// ListNamed write list named
func (w Write) ListNamed(query string, args []interface{}) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	tx, err := po.DB.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	argsLength := len(args)
	for i := 0; i < argsLength; i++ {
		_, err = stmt.Exec(args[i])
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
