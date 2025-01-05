// Package storage - contains a storage designer for application.
package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Storage - object of database.
type Storage struct {
	db *sql.DB
}

// NewStore - constractor of storage.
func NewStore(storePath string) (*Storage, error) {
	const op = "storage.NewStore"

	db, err := sql.Open("sqlite3", storePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}
