package service

import "github.com/quvonchbe05/golang-clean-architecture-template/internal/dto"

type TodoService interface {
	GetAllTodos() ([]dto.TodoOutputDTO, error)
	GetTodoById(id int) (dto.TodoOutputDTO, error)
	CreateTodo(todo dto.TodoCreateDTO) (int, error)
}