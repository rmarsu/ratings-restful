package repository

import (
	"github.com/jmoiron/sqlx"
)

type Stars interface {
}
type Repository struct {
	Stars
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Stars: NewRatingsPostgres(db),
	}

}
