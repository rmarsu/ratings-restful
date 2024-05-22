package service

import (
	rate "rate"
	"rate/pkg/repository"
)

type Stars interface {
	Create(list rate.Stars) (int, error)
}


type Service struct {
	Stars
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Stars:   NewRatingsService(repos.Stars),
	}
}
