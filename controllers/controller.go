package controllers

import (
	"go/token"
	"net/http"
	"strconv"

	"github.com/Tamiru-Alemnew/task-manager/data"
	"github.com/Tamiru-Alemnew/task-manager/models"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
    tasks, err := data.GetAllTasks()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tasks)
}


func GetTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
        return
    }
    task, err := data.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}

func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    createdTask , err := data.CreateTask(task)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    
    c.JSON(http.StatusCreated, createdTask)
}

func UpdateTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
        return
    }
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    updatedTask, err := data.UpdateTask(id, task)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid task ID"})
        return
    }
    err = data.DeleteTask(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }
    c.Status(http.StatusNoContent)
}


func SignUp(c *gin.Context){
    var user models.User

    if err := c.ShouldBindBodyWithJSON(&user) ; err!= nil {
        c.JSON(http.StatusBadRequest , err)
        return 
    }

    err := data.UserRegistration(user)

    if err != nil {
        c.JSON(http.StatusBadRequest , err)
        return 
    }

    c.JSON(201, gin.H{"message": "Signup successful"})
}

func Login(c *gin.Context){
    var user models.User

    err := c.ShouldBindJSON(&user)

     if err != nil {
        c.JSON(http.StatusBadRequest , err)
        return 
    }
    
    jwtToken , err := data.UserCredentialValidation(user)
     if err != nil {
        c.JSON(http.StatusBadRequest , err)
        return 
    }

    c.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})
}