package service

import (
	rate "rate"
	"rate/pkg/repository"
)

type RatingsService struct {
	repo repository.Stars
}

func NewRatingsService(repo repository.Stars) *RatingsService {
	return &RatingsService{repo: repo}
}

func (s *RatingsService) Create(list rate.Stars) (int, error) {
    return s.repo.Create(list)
}