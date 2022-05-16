package db

import (
	"github.com/jmoiron/sqlx"
)

// Read Read
type Read struct{}

// NewRead new read
func NewRead() *Read {
	return &Read{}
}

// List read list
func (r *Read) List(query string, args ...interface{}) (*sqlx.Rows, error) {
	stmt, err := db.Preparex(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Queryx(args...)
}

// ListNamed read list named
func (r *Read) ListNamed(query string, args interface{}) (*sqlx.Rows, error) {
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.Queryx(args)
}

// ListStruct read list struct
func (r *Read) ListStruct(dest interface{}, query string, args ...interface{}) error {
	stmt, err := db.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.Select(dest, args...)
}

// ListStructNamed list struct named
func (r *Read) ListStructNamed(dest interface{}, query string, args interface{}) error {
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.Select(dest, args)
}

// Item read item
func (r *Read) Item(query string, args ...interface{}) (*sqlx.Row, error) {
	stmt, err := db.Preparex(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowx(args...), nil
}

// ItemNamed read item named
func (r *Read) ItemNamed(query string, args interface{}) (*sqlx.Row, error) {
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return stmt.QueryRowx(args), nil
}

// ItemStruct read item struct
func (r *Read) ItemStruct(dest interface{}, query string, args ...interface{}) error {
	stmt, err := db.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.Get(dest, args...)
}

// ItemStructNamed item struct named
func (r *Read) ItemStructNamed(dest interface{}, query string, args interface{}) error {
	stmt, err := db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.Get(dest, args)
}
