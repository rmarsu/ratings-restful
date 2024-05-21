package service

import (
	"rate/pkg/repository"
)

type RatingsService struct {
	repo repository.Stars
}

func NewRatings(repo repository.Stars) *RatingsService {
	return &RatingsService{repo: repo}
}
