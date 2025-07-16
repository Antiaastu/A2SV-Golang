package data

import (
	"task_manager/models"
)

var tasks []models.Task
var nextID = 1

func GetTasks() []models.Task {
	return tasks
}

func GetTaskByID(id int) (models.Task, bool) {
	for _, task := range tasks{
		if task.ID == id{
			return task, true
		}
	}
	return models.Task{}, false
}

func CreateTask(task models.Task) models.Task{
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id int, updated models.Task) (models.Task, bool){
	for i, task := range tasks{
		if task.ID == id{
			updated.ID = id
			tasks[i] = updated
			return updated, true
		}
	}
	return models.Task{}, false
}
func DeleteTask(id int) bool{
	for i, task := range tasks{
		if task.ID == id {
			tasks = append(tasks[:i],tasks[i+1:]...)
			return true
		}
	}
	return false
}