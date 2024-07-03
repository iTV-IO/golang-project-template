package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/quvonchbe05/golang-clean-architecture-template/app/http/handler"
)

func NewRoute(handler *handler.Handler) *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			todo := v1.Group("/todo")
			{
				todo.GET("/list", handler.TodoHandler.GetAllTodos)
				todo.GET("/:id", handler.TodoHandler.GetTodoById)
				todo.POST("/create", handler.TodoHandler.CreateTodo)
			}
		}
	}

	return router
}
