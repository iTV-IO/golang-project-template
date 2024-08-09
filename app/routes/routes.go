package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/quvonchbe05/golang-clean-architecture-template/app/http/controller"
)

func NewRoute(controller *controller.Controller) *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			todo := v1.Group("/todo")
			{
				todo.GET("/list", controller.TodoController.GetAllTodos)
				todo.GET("/:id", controller.TodoController.GetTodoById)
				todo.POST("/create", controller.TodoController.CreateTodo)
			}
		}
	}

	return router
}
