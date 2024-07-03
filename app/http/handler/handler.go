package handler

import "github.com/quvonchbe05/golang-clean-architecture-template/internal/service"

type Handler struct {
	*TodoHandler
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		TodoHandler: NewTodoHandler(service),
	}
}
