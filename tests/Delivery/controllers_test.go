package tests

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"

	"testing"

	"github.com/Tamiru-Alemnew/task-manager/Delivery/controllers"
	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	"github.com/Tamiru-Alemnew/task-manager/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type controllerSuite struct {
	suite.Suite
	taskUsecase    *mocks.TaskUsecase
	userUsecase    *mocks.UserUsecase
	taskController controllers.TaskController
	userController controllers.UserController
	testingServer  *httptest.Server
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (suite *controllerSuite) SetupSuite() {
	suite.taskUsecase = new(mocks.TaskUsecase)
	suite.userUsecase = new(mocks.UserUsecase)
	suite.taskController = controllers.TaskController{
		TaskUsecase: suite.taskUsecase, 
	}

	suite.userController = controllers.UserController{
		UserUsecase: suite.userUsecase,
	}

	router := gin.Default()

	// Task routes
	router.GET("/tasks", suite.taskController.GetAllTasks)
	router.GET("/tasks/:id", suite.taskController.GetTaskByID)
	router.POST("/tasks", suite.taskController.CreateTask)
	router.PUT("/tasks/:id", suite.taskController.UpdateTask)
	router.DELETE("/tasks/:id", suite.taskController.DeleteTask)

	// User routes
	router.POST("/signup", suite.userController.SignUp)
	router.POST("/login", suite.userController.Login)
	router.PATCH("/promote/:id", suite.userController.Promote)

	suite.testingServer = httptest.NewServer(router)
}

func (suite *controllerSuite) SetupTest() {
	suite.taskUsecase = new(mocks.TaskUsecase)
	suite.userUsecase = new(mocks.UserUsecase)
	suite.taskController.TaskUsecase = suite.taskUsecase
	suite.userController.UserUsecase = suite.userUsecase
}

func (suite *controllerSuite) TearDownSuite() {
	defer suite.testingServer.Close()
}

// Tests for TaskController

func (suite *controllerSuite) TestGetAllTasks_Positive() {
    task := domain.Task{
        ID:          1,
        Title:       "title",
        Description: "description",
        DueDate:     "2023-12-31T23:59:59Z",
        Status:      "pending",
    }

    // Set up the mock to expect the GetAllTasks method call and return the appropriate values
    suite.taskUsecase.On("GetAll", mock.Anything).Return([]domain.Task{task}, nil)

    response, err := http.Get(suite.testingServer.URL + "/tasks")
    if response != nil {
        defer response.Body.Close()
    }

    suite.NoError(err, "no errors in request")
    suite.Equal(http.StatusOK, response.StatusCode)
    suite.taskUsecase.AssertExpectations(suite.T())

    var tasks []domain.Task
    err = json.NewDecoder(response.Body).Decode(&tasks)
    suite.NoError(err, "no error during body decoding")
    suite.Equal(1, len(tasks), "sends data correctly")
}

func (suite *controllerSuite) TestGetAllTasks_Negative() {
	sampleErr := errors.New("sample error message")
	suite.taskUsecase.On("GetAll", mock.Anything).Return(nil, sampleErr)
	response, err := http.Get(suite.testingServer.URL + "/tasks")
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusInternalServerError, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestGetTaskByID_Positive() {
	task := domain.Task{
		ID:          1,
		Title:       "title",
		Description: "description",
		DueDate:     "2023-12-31T23:59:59Z",
		Status:      "pending",
	}

	// Update mock setup to match actual method signature
	suite.taskUsecase.On("GetByID", mock.Anything, task.ID).Return(&task, nil)


	response, err := http.Get(suite.testingServer.URL + "/tasks/" + strconv.Itoa(task.ID))
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusOK, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())

	var fetchedTask domain.Task
	err = json.NewDecoder(response.Body).Decode(&fetchedTask)
	suite.NoError(err, "no error during body decoding")
	suite.Equal(task.ID, fetchedTask.ID, "sends task with correct ID")
	suite.Equal(task.Title, fetchedTask.Title, "sends task with correct Title")
	suite.Equal(task.Description, fetchedTask.Description, "sends task with correct Description")
	suite.Equal(task.DueDate, fetchedTask.DueDate, "sends task with correct DueDate")
	suite.Equal(task.Status, fetchedTask.Status, "sends task with correct Status")
}


func (suite *controllerSuite) TestGetTaskByID_Negative() {
	wrongID := 2
	sampleErr := errors.New("task not found")
	suite.taskUsecase.On("GetByID", mock.Anything, wrongID).Return(&domain.Task{}, sampleErr)
	response, err := http.Get(suite.testingServer.URL + "/tasks/" + strconv.Itoa(wrongID))
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusNotFound, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestAddTask_Positive() {
	newTask := domain.Task{}
	client := http.Client{}

	suite.taskUsecase.On("Create",mock.Anything, &newTask).Return(&newTask, nil)

	requestBody, err := json.Marshal(&newTask)
	suite.NoError(err, "can not marshal struct to json")

	request, _ := http.NewRequest(http.MethodPost, suite.testingServer.URL+"/tasks", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestAddTask_Negative() {
	newTask := domain.Task{}
	client := http.Client{}
	sampleErr := errors.New("failed to add task")
	suite.taskUsecase.On("Create", mock.Anything, &newTask).Return(nil , sampleErr)

	requestBody, err := json.Marshal(&newTask)
	suite.NoError(err, "can not marshal struct to json")

	request, _ := http.NewRequest(http.MethodPost, suite.testingServer.URL+"/tasks", bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusInternalServerError, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestUpdateTask_Positive() {
	taskID := 1
	taskUpdates := domain.Task{
		Title:       "Updated title",
		Description: "Updated description",
		Status:      "completed",
	}

	client := http.Client{}
	suite.taskUsecase.On("Update", mock.Anything, taskID, &taskUpdates).Return(&taskUpdates, nil)

	requestBody, err := json.Marshal(&taskUpdates)
	suite.NoError(err, "can not marshal struct to json")

	request, _ := http.NewRequest(http.MethodPut, suite.testingServer.URL+"/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusOK, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())

	var fetchedTask domain.Task
	err = json.NewDecoder(response.Body).Decode(&fetchedTask)
	suite.NoError(err, "no error during body decoding")
	suite.Equal(taskUpdates.Title, fetchedTask.Title, "sends task with correct Title")
	suite.Equal(taskUpdates.Description, fetchedTask.Description, "sends task with correct Description")
	suite.Equal(taskUpdates.Status, fetchedTask.Status, "sends task with correct Status")
}

func (suite *controllerSuite) TestUpdateTask_Negative() {
	taskID := 1
	taskUpdates := domain.Task{}
	client := http.Client{}
	sampleErr := errors.New("failed to update task")
	suite.taskUsecase.On("Update", mock.Anything, taskID, mock.Anything).Return(&taskUpdates, sampleErr)

	requestBody, err := json.Marshal(&taskUpdates)
	suite.NoError(err, "can not marshal struct to json")

	request, _ := http.NewRequest(http.MethodPut, suite.testingServer.URL+"/tasks/"+strconv.Itoa(taskID), bytes.NewBuffer(requestBody))
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusInternalServerError, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestDeleteTask_Positive() {
	taskID := 1
	client := http.Client{}
	suite.taskUsecase.On("Delete", mock.Anything, taskID).Return(nil)

	request, _ := http.NewRequest(http.MethodDelete, suite.testingServer.URL+"/tasks/"+strconv.Itoa(taskID), nil)
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusNoContent, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestDeleteTask_Negative() {
	taskID := 1
	client := http.Client{}
	sampleErr := errors.New("failed to delete task")
	suite.taskUsecase.On("Delete", mock.Anything, taskID).Return(sampleErr)

	request, _ := http.NewRequest(http.MethodDelete, suite.testingServer.URL+"/tasks/"+strconv.Itoa(taskID), nil)
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusInternalServerError, response.StatusCode)
	suite.taskUsecase.AssertExpectations(suite.T())
}

// Tests for UserController

func (suite *controllerSuite) TestSignUp_Positive() {
	user := &domain.User{
		Username: "Tamiru",
		Password: "",
		Role: "user",
	}

	suite.userUsecase.On("SignUp", mock.Anything, user).Return(user ,nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(suite.testingServer.URL+"/signup", "application/json", bytes.NewBuffer(requestBody))
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestSignUp_Negative() {
	user := &domain.User{
		ID:       1,
		Username:    "Tamiru",
		Password: "",
	}

	sampleErr := errors.New("failed to sign up user")
	suite.userUsecase.On("SignUp", mock.Anything, user).Return(nil , sampleErr)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(suite.testingServer.URL+"/signup", "application/json", bytes.NewBuffer(requestBody))
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusInternalServerError, response.StatusCode)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestLogin_Positive() {
	user := domain.User{}

	token := "sample_token"
	suite.userUsecase.On("Login", mock.Anything, user.Username, user.Password).Return(&user , token, nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(suite.testingServer.URL+"/login", "application/json", bytes.NewBuffer(requestBody))
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusOK, response.StatusCode)
	suite.userUsecase.AssertExpectations(suite.T())

	var tokenResponse TokenResponse
	err = json.NewDecoder(response.Body).Decode(&tokenResponse)
	suite.NoError(err, "no error during body decoding")
	suite.Equal(token, tokenResponse.Token, "sends correct token")
}

func (suite *controllerSuite) TestLogin_Negative() {
	user := domain.User{}

	sampleErr := errors.New("invalid credentials")
	suite.userUsecase.On("Login", mock.Anything, user.Username, user.Password).Return(nil, "", sampleErr)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(suite.testingServer.URL+"/login", "application/json", bytes.NewBuffer(requestBody))
	if response != nil {
		defer response.Body.Close()
	}
	suite.NoError(err, "no errors in request")
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestPromote_Positive() {
	userID := 1
	client := http.Client{}
	suite.userUsecase.On("Promote", mock.Anything, userID).Return(nil)

	request, _ := http.NewRequest(http.MethodPatch, suite.testingServer.URL+"/promote/"+strconv.Itoa(userID), nil)
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusOK, response.StatusCode)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *controllerSuite) TestPromote_Negative() {
	userID := 1
	client := http.Client{}
	sampleErr := errors.New("failed to promote user")
	suite.userUsecase.On("Promote", mock.Anything, userID).Return(sampleErr)

	request, _ := http.NewRequest(http.MethodPatch, suite.testingServer.URL+"/promote/"+strconv.Itoa(userID), nil)
	response, err := client.Do(request)
	if response != nil {
		defer response.Body.Close()
	}

	suite.NoError(err, "no errors in request")
	suite.Equal(http.StatusInternalServerError, response.StatusCode)
	suite.userUsecase.AssertExpectations(suite.T())
}

func TestControllerSuite(t *testing.T) {
	suite.Run(t, new(controllerSuite))
}
