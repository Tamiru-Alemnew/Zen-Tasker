package controllers

import (
	"net/http"
	"strconv"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
    taskUsecase domain.TaskUsecase
}
type UserController struct {
    userUsecase domain.UserUsecase
}


func NewTaskController(tc domain.TaskUsecase) *TaskController {
    return &TaskController{
        taskUsecase: tc,
    }
}

func NewUserController(uc domain.UserUsecase) *UserController {
    return &UserController{
        userUsecase: uc,
    }
}


func (ctrl *TaskController) GetAllTasks(c *gin.Context) {
    tasks, err := ctrl.taskUsecase.GetAll(c.Request.Context())
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

func (ctrl *TaskController) CreateTask(c *gin.Context) {
    var task domain.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    createdTask, err := ctrl.taskUsecase.Create(c.Request.Context(), &task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdTask)
}

func (ctrl *TaskController) GetTaskByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    task, err := ctrl.taskUsecase.GetByID(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, task)
}

func (ctrl *TaskController) UpdateTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    var task domain.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    updatedTask, err := ctrl.taskUsecase.Update(c.Request.Context(), id, &task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedTask)
}


func (ctrl *TaskController) DeleteTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
        return
    }

    err = ctrl.taskUsecase.Delete(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.Status(http.StatusNoContent)
}

func (ctrl *UserController) SignUp(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    createdUser, err := ctrl.userUsecase.SignUp(c.Request.Context(), &user)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, createdUser)
}

func (ctrl *UserController) Login(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    loggedInUser, token, err := ctrl.userUsecase.Login(c.Request.Context(), user.Username, user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": loggedInUser, "token": token})
}

func (ctrl *UserController) Promote(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    err = ctrl.userUsecase.Promote(c.Request.Context(), id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User promoted to admin successfully"})
}

