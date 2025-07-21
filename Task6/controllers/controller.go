package controllers

import (
	"Task6/data"
	"Task6/middleware"
	"Task6/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)
//User Handlers

func Register(ctx *gin.Context){
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Role == ""{
		user.Role = "user"
	}

	if err := data.Register(user); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User Registered successfully"})
}

func Login(ctx *gin.Context){
	var credentials models.User
	if err := ctx.ShouldBindJSON(&credentials); err != nil{
		ctx.JSON(http.StatusBadRequest,gin.H{"error": err.Error()})
		return
	}
	user, err := data.Login(credentials.Username,credentials.Password)
	if err != nil{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	expire := time.Now().Add(24 * time.Hour)
	claims := &middleware.Claims{
		Username: user.Username,
		Role : user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString(middleware.JwtKey())
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK,gin.H{"token":tokenString})
}

// Task Handlers
func GetTasks(ctx *gin.Context) {
	tasks, err := data.GetTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := data.GetTaskById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func CreateTask(ctx *gin.Context) {
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := data.CreateTask(newTask)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, created)
}

func UpdateTask(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := data.UpdateTask(id, updatedTask)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updated)
}

func DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	err := data.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
