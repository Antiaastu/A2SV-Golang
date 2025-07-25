package routers

import(
	"Task7/delivery/controllers"
	"Task7/infrastructure"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router and sets up the routes.
func SetupRouter(ctrl *controllers.Controller) *gin.Engine {
	r := gin.Default()
	r.POST("/register", ctrl.Register)
	r.POST("/login", ctrl.Login)

	api := r.Group("/api", infrastructure.AuthMiddleware())
	api.POST("/tasks", ctrl.CreateTask)
	api.GET("/tasks", ctrl.GetTasks)
	api.GET("/tasks/:id", ctrl.GetTaskByID)
	api.PUT("/tasks/:id", ctrl.UpdateTask)
	api.DELETE("/tasks/:id", infrastructure.AuthMiddleware(), ctrl.DeleteTask)
	return r

}