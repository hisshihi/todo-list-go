package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"slices"
)

type Task struct {
	ID int `json:"id" binding:"required"`
	Title string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=3,max=1000"`
	Completed bool `json:"completed" binding:"required"`
}

type TaskRequest struct {
	Title string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=3,max=1000"`
}

type TaskResponse struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}

var tasks []Task

func CreateTask(ctx *gin.Context) {
	var req TaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	newTask := Task{
		ID: len(tasks) + 1,
		Title: req.Title,
		Description: req.Description,
		Completed: false,
	}

	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, TaskResponse{
		Title: newTask.Title,
		Description: newTask.Description,
		Completed: newTask.Completed,
	})
}

func GetTasks(ctx *gin.Context) {
	if len(tasks) == 0 {
		ctx.JSON(http.StatusNoContent, gin.H{"message": "No tasks found"})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

type TaskIDRequest struct {
	ID int `uri:"id" binding:"required,min=1"`
}

func GetTaskByID(ctx *gin.Context) {
	var req TaskIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, task := range tasks {
		if task.ID == req.ID {
			ctx.JSON(http.StatusOK, TaskResponse{
				Title:       task.Title,
				Description: task.Description,
				Completed:   task.Completed,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

type TaskUpdateRequest struct {
	Title string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=3,max=1000"`
	Completed bool `json:"completed" binding:"required"`
}

func UpdateTask(ctx *gin.Context) {
	var findId TaskIDRequest
	if err := ctx.ShouldBindUri(&findId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req TaskUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, task := range tasks {
		if task.ID == findId.ID {
			tasks[i] = Task{
				ID: task.ID,
				Title: req.Title,
				Description: req.Description,
				Completed: req.Completed,
			}

			ctx.JSON(http.StatusOK, TaskResponse{
				Title:       tasks[i].Title,
				Description: tasks[i].Description,
				Completed:   tasks[i].Completed,
			})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

func DeleteTask(ctx *gin.Context) {
	var req TaskIDRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	for i, task := range tasks {
		if task.ID == req.ID {
			tasks = slices.Delete(tasks, i, i+1)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
			return
		}
	}

	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}