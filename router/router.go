package router

import (
	"github.com/Tamiru-Alemnew/task-manager/controllers"
	"github.com/Tamiru-Alemnew/task-manager/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/register" , controllers.SignUp)
    r.POST("/login" , controllers.Login)
    
    r.GET("/tasks",middleware.AuthMiddleware(), controllers.GetTasks)
    r.GET("/tasks/:id",middleware.AuthMiddleware(), controllers.GetTask)
    r.POST("/tasks", middleware.AuthMiddleware(),controllers.CreateTask)
    r.PUT("/tasks/:id", middleware.AuthMiddleware() , controllers.UpdateTask)
    r.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.RoleAuthorizationMiddleware("admin"),controllers.DeleteTask)
    return r
}