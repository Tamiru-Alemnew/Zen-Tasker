package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(){
	router := gin.Default()

	router.GET("/tasks", GetTasks)
	router.GET("/tasks/:id", GetTask)
	router.POST("/tasks", CreateTask)
	router.PUT("/tasks/:id", UpdateTask)
	router.DELETE("/tasks/:id", DeleteTask)

}