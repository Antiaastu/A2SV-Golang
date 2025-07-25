package main

import (
	"Task7/config"
	"Task7/delivery/controllers"
	"Task7/delivery/routers"
	"Task7/repositories"
	"Task7/usecases"
	"fmt"
)

func main() {
	db := config.DB

	userRepo := repositories.NewMongoUserRepo(db.Collection("users"))
	taskRepo := repositories.NewMongoTaskRepo(db.Collection("tasks"))

	userUC := usecases.NewUserUsecase(userRepo)
	taskUC := usecases.NewTaskUsecase(taskRepo)

	ctrl := controllers.NewController(*userUC, *taskUC)
	r := routers.SetupRouter(ctrl)
	r.Run(":8080")
	fmt.Println("Server running on port 8080")
}
