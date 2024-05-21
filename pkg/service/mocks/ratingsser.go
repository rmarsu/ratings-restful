package service

import (
    "rate/pkg/repository"
)

type TodoListService struct {
	repo repository.Stars
}

func NewTodoListService(repo repository.Stars) *TodoListService {
    return &TodoListService{repo: repo}
}