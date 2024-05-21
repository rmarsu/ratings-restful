package repository

import "github.com/jmoiron/sqlx"

type StarsPostgres struct {
	db *sqlx.DB
}

func NewRatingsPostgres(db *sqlx.DB) *StarsPostgres {
	return &StarsPostgres{db: db}
}
