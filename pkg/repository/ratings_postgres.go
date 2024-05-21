package repository

import "github.com/jmoiron/sqlx"

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewRatingsPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}
