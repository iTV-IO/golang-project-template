package controller

import "github.com/quvonchbe05/golang-clean-architecture-template/internal/service"

type Controller struct {
	*TodoController
}

func NewController(service *service.Service) *Controller {
	return &Controller{
		TodoController: NewTodoController(service),
	}
}
