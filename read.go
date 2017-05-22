package db

import "github.com/jmoiron/sqlx"

// Read Read
type Read struct{}

// NewRead new read
func NewRead() *Read {
	return &Read{}
}

// List read list
func (r Read) List(query string, args ...interface{}) (*sqlx.Rows, error) {
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

// Item read item
func (r Read) Item(query string, args ...interface{}) (*sqlx.Row, error) {
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

// ListStruct read list struct
func (r Read) ListStruct(dest interface{}, query string, args ...interface{}) error {
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

// ItemStruct read item struct
func (r Read) ItemStruct(dest interface{}, query string, args ...interface{}) error {
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

// ListNamed read list named
func (r Read) ListNamed(query string, args map[string]interface{}) (*sqlx.Rows, error) {
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

// ItemNamed read item named
func (r Read) ItemNamed(query string, args map[string]interface{}) (*sqlx.Row, error) {
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

// ListStructNamed list struct named
func (r Read) ListStructNamed(dest interface{}, query string, args map[string]interface{}) error {
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

// ItemStructNamed item struct named
func (r Read) ItemStructNamed(dest interface{}, query string, args map[string]interface{}) error {
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
