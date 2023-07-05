package handlers

import (
	"database/sql"
	"gym/ankur/database"
)

type ApiHandler struct {
	Db      *sql.DB
	Storage *database.Repository
}

func NewApiHandler(db *sql.DB) *ApiHandler {
	return &ApiHandler{Db: db, Storage: database.NewFunc(db)}
}
