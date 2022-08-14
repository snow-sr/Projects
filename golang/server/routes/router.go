package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/snow-sr/learningGo/controllers"
)

func Config(router *gin.Engine) *gin.Engine {
	main := router.Group("/")
	{
		todo := main.Group("/todo")
		{
			todo.GET("/", controllers.GetAll)
			todo.GET("/getNonCompleted", controllers.GetNonCompleted)
			todo.GET("/:id", controllers.GetById)
			todo.GET("/delete/:id", controllers.DeleteById)
			todo.POST("/", controllers.CreateTodo)
			todo.GET("/solve/:id", controllers.SolveTodo)
		}
	}
	return router
}
