package repository

import (
	"fmt"
	rate "rate"

	"github.com/jmoiron/sqlx"
)

type StarsPostgres struct {
	db *sqlx.DB
}

func NewRatingsPostgres(db *sqlx.DB) *StarsPostgres {
	return &StarsPostgres{db: db}
}
func (r *StarsPostgres) Create(rates rate.Stars) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	StarsQuery := fmt.Sprintf("INSERT INTO %s (obj, stars) VALUES ($1, $2) RETURNING id", Rating)
	row := tx.QueryRow(StarsQuery, rates.Obj, rates.Stars)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
