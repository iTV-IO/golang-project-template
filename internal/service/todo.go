package service

import (
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/repository"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/model/request"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/model/response"
)

type Todo interface {
	GetAllTodos() ([]response.Todo, error)
	GetTodoById(id int) (response.Todo, error)
	CreateTodo(todo *request.TodoCreate) (int, error)
}

type TodoService struct {
	repo repository.Todo
}

func NewTodoService(repo repository.Todo) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) GetAllTodos() ([]response.Todo, error) {
	return s.repo.GetAllTodos()
}

func (s *TodoService) GetTodoById(id int) (response.Todo, error) {
	return s.repo.GetTodoById(id)
}

func (s *TodoService) CreateTodo(todo *request.TodoCreate) (int, error) {
	return s.repo.CreateTodo(todo)
}
