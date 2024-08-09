package service

import (
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/dto"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/repository"
)

type TodoServiceImpl struct {
	repo repository.Todo
}

func NewTodoServiceImpl(repo repository.Todo) *TodoServiceImpl {
	return &TodoServiceImpl{repo: repo}
}

func (s *TodoServiceImpl) GetAllTodos() ([]dto.TodoOutputDTO, error) {
	return s.repo.GetAllTodos()
}

func (s *TodoServiceImpl) GetTodoById(id int) (dto.TodoOutputDTO, error) {
	return s.repo.GetTodoById(id)
}

func (s *TodoServiceImpl) CreateTodo(todo dto.TodoCreateDTO) (int, error) {
	return s.repo.CreateTodo(todo)
}
