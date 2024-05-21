package repository

import (

	"github.com/jmoiron/sqlx"
)

const (
	ActualContent = "content"
	Rating        = "rating"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
