package db

import "github.com/jmoiron/sqlx"

type read struct{}

// Read Read
var Read = &read{}

// List Query
func (r *read) List(query string, args ...interface{}) (*sqlx.Rows, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.Preparex(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Queryx(args...)
}

// Item QueryRow
func (r *read) Item(query string, args ...interface{}) (*sqlx.Row, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.Preparex(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.QueryRowx(args...), nil
}

// Listx Select
func (r *read) ListStruct(dest interface{}, query string, args ...interface{}) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.Select(dest, args...)
}

// Itemx Get
func (r *read) ItemStruct(dest interface{}, query string, args ...interface{}) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.Get(dest, args...)
}

// Itemx Get
func (r *read) ListNamed(query string, args map[string]interface{}) (*sqlx.Rows, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Queryx(args)
}

// Itemx Get
func (r *read) ItemNamed(query string, args map[string]interface{}) (*sqlx.Row, error) {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.QueryRowx(args), nil
}

// Listx Select
func (r *read) ListStructNamed(dest interface{}, query string, args map[string]interface{}) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.Select(dest, args)
}

// Itemx Get
func (r *read) ItemStructNamed(dest interface{}, query string, args map[string]interface{}) error {
	<-conn
	defer func() {
		conn <- 1
	}()
	stmt, err := po.DB.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.Get(dest, args)
}
