package controllers

import (
	"net/http"
	"strconv"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context) {
	tasks := data.GetTasks()
	ctx.JSON(http.StatusOK, tasks)
}

func GetTaskByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	task, found := data.GetTaskByID(id)
	if !found {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
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
	createdTask := data.CreateTask(newTask)
	ctx.JSON(http.StatusCreated, createdTask)
}

func UpdateTask(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	task, found := data.UpdateTask(id,updatedTask)
	if !found{
		ctx.JSON(http.StatusNotFound,gin.H{"message": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK,task)
}

func DeleteTask(ctx *gin.Context){
	id, _ := strconv.Atoi(ctx.Param("id"))
	found := data.DeleteTask(id)
	if !found {
		ctx.JSON(http.StatusNotFound,gin.H{"message":"Task not found"})
		return
	}
	ctx.JSON(http.StatusOK,gin.H{"message":"Task deleted successfully"})
}