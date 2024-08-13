package router

import (
	"github.com/Tamiru-Alemnew/task-manager/Delivery/controllers"
	middleware "github.com/Tamiru-Alemnew/task-manager/Infrastructures"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()
    r.POST("/register" , controllers.SignUp)
    r.POST("/login" , controllers.Login)

    r.GET("/tasks",middleware.AuthMiddleware(), controllers)
    r.GET("/tasks/:id",middleware.AuthMiddleware(), controllers.GetTask)
    
    r.PATCH("promote/:id", middleware.AuthMiddleware() , middleware.RoleAuthorizationMiddleware("admin") , controllers.Promote)
    r.POST("/tasks", middleware.AuthMiddleware(),middleware.RoleAuthorizationMiddleware("admin"),controllers.CreateTask)
    r.PUT("/tasks/:id", middleware.AuthMiddleware(),middleware.RoleAuthorizationMiddleware("admin") , controllers.UpdateTask)
    r.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.RoleAuthorizationMiddleware("admin"),controllers.DeleteTask)
    return r
}