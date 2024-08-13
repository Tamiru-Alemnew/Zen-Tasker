package router

import (
	"github.com/Tamiru-Alemnew/task-manager/Delivery/controllers"
	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	middleware "github.com/Tamiru-Alemnew/task-manager/Infrastructures"
	"github.com/gin-gonic/gin"
)

func SetupRouter(taskUsecase domain.TaskUsecase, userUsecase domain.UserUsecase) *gin.Engine {
    router := gin.Default()

    // Inject controllers
    taskController := controllers.NewTaskController(taskUsecase)
    userController := controllers.NewUserController(userUsecase)

    router.POST("/register" , userController.SignUp)
    router.POST("/login" , userController.Login)

    router.GET("/tasks",middleware.AuthMiddleware(), taskController.GetAllTasks)
    router.GET("/tasks/:id",middleware.AuthMiddleware(), taskController.GetTaskByID)
    
    router.PATCH("promote/:id", middleware.AuthMiddleware() , middleware.RoleAuthorizationMiddleware("admin") , userController.Promote)
    router.POST("/tasks", middleware.AuthMiddleware(),middleware.RoleAuthorizationMiddleware("admin"), taskController.CreateTask)
    router.PUT("/tasks/:id", middleware.AuthMiddleware(),middleware.RoleAuthorizationMiddleware("admin") , taskController.UpdateTask)
    router.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.RoleAuthorizationMiddleware("admin"),taskController.DeleteTask)
    return router
}