package database

import "database/sql"

type Repository struct {
	Db *sql.DB
}

func NewFunc(db *sql.DB) *Repository {
	return &Repository{Db: db}
}
