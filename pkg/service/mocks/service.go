package service

import (
	rate "rate"
	"rate/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user rate.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type rateList interface {
	Create(userId int, list rate.rateList) (int, error)
	GetAll(userId int) ([]rate.rateList, error)
	GetById(userId, listId int) (rate.rateList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input rate.UpdateListInput) error
}

type rateItem interface {
	Create(userId, listId int, item rate.rateItem) (int, error)
	GetAll(userId, listId int) ([]rate.rateItem, error)
	GetById(userId, itemId int) (rate.rateItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input rate.UpdateItemInput) error
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Ratings: NewrateListService(repos.Ratings),
	}
}
