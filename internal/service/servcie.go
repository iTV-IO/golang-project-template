package service

import "github.com/quvonchbe05/golang-clean-architecture-template/internal/repository"

type Service struct {
	TodoService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		TodoService: NewTodoServiceImpl(repo),
	}
}
