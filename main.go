package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hisshihi/todo-list-go/api"
)

func main() {
	r := gin.Default()
	r.GET("/tasks", api.GetTasks)
	r.POST("/tasks", api.CreateTask)
	r.GET("/tasks/:id", api.GetTaskByID)
	r.PUT("/tasks/:id", api.UpdateTask)
	r.DELETE("/tasks/:id", api.DeleteTask)
	r.Run(":8080")
}