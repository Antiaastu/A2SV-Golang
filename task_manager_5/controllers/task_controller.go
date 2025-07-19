package controllers
import (
	"net/http"
	"task_manager_5/data"
	"task_manager_5/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(ctx *gin.Context){
	tasks, err := data.GetTasks()
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,tasks)
}

func GetTaskById(ctx *gin.Context){
	id := ctx.Param("id")
	task, err := data.GetTaskById(id)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,task)
}

func CreateTask(ctx *gin.Context){
	var newTask models.Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	created, err := data.CreateTask(newTask)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated,created)
}

func UpdateTask(ctx *gin.Context){
	id := ctx.Param("id")
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updated, err := data.UpdateTask(id,updatedTask)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,updated)
}

func DeleteTask(ctx *gin.Context){
	id := ctx.Param("id")
	err := data.DeleteTask(id)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}