package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	todo_request "github.com/quvonchbe05/golang-clean-architecture-template/app/http/request/todo"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/dto"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/service"
)

type TodoController struct {
	service service.TodoService
}

func NewTodoController(service service.TodoService) *TodoController {
	return &TodoController{service: service}
}

func (h *TodoController) GetAllTodos(c *gin.Context) {
	todos, err := h.service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": todos,
	})
}

func (h *TodoController) GetTodoById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	todos, err := h.service.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": todos,
	})
}

func (h *TodoController) CreateTodo(c *gin.Context) {
	var todo todo_request.TodoCreateRequest

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto := dto.TodoCreateDTO{
		Title: todo.Title,
	}

	id, err := h.service.CreateTodo(dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
