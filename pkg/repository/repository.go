package repository

import (
	"github.com/jmoiron/sqlx"
	rate "rate"
)

type Stars interface {
	Create(list rate.Stars) (int, error)
}

type Repository struct {
	Stars
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Stars: NewRatingsPostgres(db),
	}

}
