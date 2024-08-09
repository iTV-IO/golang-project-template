package repository

import (
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/dto"
	"gorm.io/gorm"
)

type Todo interface {
	GetAllTodos() ([]dto.TodoOutputDTO, error)
	GetTodoById(id int) (dto.TodoOutputDTO, error)
	CreateTodo(todo dto.TodoCreateDTO) (int, error)
}

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepositoryImpl(db *gorm.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{db: db}
}

func (t *TodoRepositoryImpl) GetAllTodos() ([]dto.TodoOutputDTO, error) {
	var todos []dto.TodoOutputDTO

	// Logic

	return todos, nil
}

func (t *TodoRepositoryImpl) GetTodoById(id int) (dto.TodoOutputDTO, error) {
	var todo dto.TodoOutputDTO

	// Logic

	return todo, nil
}

func (t *TodoRepositoryImpl) CreateTodo(todo dto.TodoCreateDTO) (int, error) {
	var id int

	// Logic

	return id, nil
}
