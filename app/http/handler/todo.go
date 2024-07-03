package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/model/request"
	"github.com/quvonchbe05/golang-clean-architecture-template/internal/service"
)

type TodoHandler struct {
	service *service.Service
}

func NewTodoHandler(service *service.Service) *TodoHandler {
	return &TodoHandler{service: service}
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.service.Todo.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": todos,
	})
}

func (h *TodoHandler) GetTodoById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid todo ID"})
		return
	}

	todos, err := h.service.Todo.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": todos,
	})
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo request.TodoCreate

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := h.service.Todo.CreateTodo(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
