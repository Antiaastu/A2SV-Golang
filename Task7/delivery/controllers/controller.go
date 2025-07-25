package controllers

import (
	"Task7/dto"
	"Task7/usecases"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	UserUC usecases.UserUsecase
	TaskUC usecases.TaskUsecase
}

func NewController(u usecases.UserUsecase, t usecases.TaskUsecase) *Controller {
	return &Controller{u, t}
}

func (c *Controller) Register(ctx *gin.Context) {
	var req dto.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := c.UserUC.Register(context.Background(), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (c *Controller) Login(ctx *gin.Context) {
	var req dto.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	token, err := c.UserUC.Login(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *Controller) CreateTask(ctx *gin.Context) {
	var req dto.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	owner := ctx.GetString("username")
	t, err := c.TaskUC.Create(context.Background(), req, owner)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, dto.TaskResponse(t))
}

func (c *Controller) GetTasks(ctx *gin.Context) {
	owner := ctx.GetString("username")
	tasks, err := c.TaskUC.GetAll(context.Background(), owner)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var taskResponses []dto.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponse(task))
	}
	ctx.JSON(http.StatusOK, taskResponses)
}

func (c *Controller) GetTaskByID(ctx *gin.Context) {
	taskID := ctx.Param("id")
	owner := ctx.GetString("username")
	task, err := c.TaskUC.GetByID(context.Background(), taskID, owner)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dto.TaskResponse(task))
}

func (c *Controller) UpdateTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	var req dto.CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	owner := ctx.GetString("username")
	task, err := c.TaskUC.Update(context.Background(), req, taskID, owner)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dto.TaskResponse(task))
}

func (c *Controller) DeleteTask(ctx *gin.Context) {
	taskID := ctx.Param("id")
	owner := ctx.GetString("username")
	if err := c.TaskUC.Delete(context.Background(), taskID, owner); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Task deleted successfully"})
}
