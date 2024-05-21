package repository

import (
	"github.com/jmoiron/sqlx"
	rate "rate"
)

type Ratings interface {
}

type Repository struct {
	Ratings
}

type Stars interface {
	Create(userId int, list rate.Stars) (int, error)
}


func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Ratings: NewRatingsPostgres(db),
	}

}
