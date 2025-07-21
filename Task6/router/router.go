package router

import (
	"Task6/controllers"
	"Task6/middleware"

	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.GET("/tasks", controllers.GetTasks)
		api.POST("/tasks", controllers.CreateTask)
		api.GET("/tasks/:id", controllers.GetTaskById)
		api.PUT("/tasks/:id", controllers.UpdateTask)
		api.DELETE("/tasks/:id", controllers.DeleteTask)
	}

	admin := api.Group("/admin")
	admin.Use(middleware.AdminOnly())
	{
		admin.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome Admin"})
		})
	}
	return r
}
