package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Todo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Todo: NewTodoRepository(db),
	}
}
