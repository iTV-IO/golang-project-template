package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/model/request"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/model/response"
)

type Todo interface {
	GetAllTodos() ([]response.Todo, error)
	GetTodoById(id int) (response.Todo, error)
	CreateTodo(todo *request.TodoCreate) (int, error)
}

type TodoRepository struct {
	db *sqlx.DB
}

func NewTodoRepository(db *sqlx.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (t *TodoRepository) GetAllTodos() ([]response.Todo, error) {
	var todos []response.Todo

	query := `
		SELECT * FROM todos ORDER BY id DESC
	`

	err := t.db.Select(&todos, query)
	if err != nil {
		return todos, fmt.Errorf("failed to get todos list: %w", err)
	}

	return todos, nil
}

func (t *TodoRepository) GetTodoById(id int) (response.Todo, error) {
	var todo response.Todo

	query := `
		SELECT * FROM todos WHERE id=$1
	`

	err := t.db.Get(&todo, query, id)
	if err != nil {
		return todo, fmt.Errorf("failed to get todo by id: %w", err)
	}

	return todo, nil
}

func (t *TodoRepository) CreateTodo(todo *request.TodoCreate) (int, error) {
	var id int

	query := `
		INSERT INTO todos (title) VALUES ($1) RETURNING id
	`

	err := t.db.QueryRow(query, todo.Title).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create todo: %w", err)
	}

	return id, nil
}
