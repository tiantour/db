package db

import "database/sql"

// Item
func (w *write) Item(query string, args ...interface{}) (sql.Result, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := pool.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}

// List
func (w *write) List(query string, args []interface{}) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	tx, err := pool.DB.Begin()
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

// Item
func (w *write) ItemNamed(query string, args interface{}) (sql.Result, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := pool.DB.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args)
}

// List
func (w *write) ListNamed(query string, args []interface{}) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	tx, err := pool.DB.Beginx()
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
