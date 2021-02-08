package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Config struct {
	DatabaseFile string
}

func NewSqliteDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", cfg.DatabaseFile)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
