package repository

import "github.com/quvonchbe05/golang-clean-architecture-template/internal/dto"

type TodoRepository interface {
	GetAllTodos() ([]dto.TodoOutputDTO, error)
	GetTodoById(id int) (dto.TodoOutputDTO, error)
	CreateTodo(todo dto.TodoCreateDTO) (int, error)
}
