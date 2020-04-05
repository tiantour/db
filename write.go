package db

import "database/sql"

// Write write
type Write struct{}

// NewWrite new *Write
func NewWrite() *Write {
	return &Write{}
}

// List write list
func (w *Write) List(query string, args []interface{}) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for i := 0; i < len(args); i++ {
		_, err = stmt.Exec(args[i])
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

// Item write
func (w *Write) Item(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return stmt.Exec(args...)
}

// ListNamed write list named
func (w *Write) ListNamed(query string, args []interface{}) error {
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for i := 0; i < len(args); i++ {
		_, err = stmt.Exec(args[i])
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
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
