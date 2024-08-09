package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	Todo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Todo: NewTodoRepositoryImpl(db),
	}
}
