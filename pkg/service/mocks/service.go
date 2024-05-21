package service

import (
	rate "rate"
	"rate/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go
type Stars interface {
	Create(int, list rate.Stars) (int, error)
}

type Ratings interface {
}

type Service struct {
	Ratings
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Ratings: NewRatings(repos.Stars),
	}
}
