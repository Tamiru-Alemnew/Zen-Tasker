package tests

import (
	"context"
	"errors"

	// "net/http"
	"net/http/httptest"
	"testing"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	usecases "github.com/Tamiru-Alemnew/task-manager/Usecases"
	"github.com/Tamiru-Alemnew/task-manager/tests/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)
type taskUsecaseSuite struct{
	suite.Suite
	taskRepository *mocks.TaskRepository
	taskUsecase usecases.TaskUsecase
	testingServer *httptest.Server
}

func (suite *taskUsecaseSuite) SetupSuite(){
	suite.taskRepository = new(mocks.TaskRepository)
	suite.taskUsecase = usecases.TaskUsecase{
		TaskRepository: suite.taskRepository,
	}
}

func(suite *taskUsecaseSuite) SetupTest(){
	suite.taskRepository = new(mocks.TaskRepository)
	suite.taskUsecase.TaskRepository = suite.taskRepository
}

func(suite *taskUsecaseSuite) TearDownSuite(){
	defer suite.testingServer.Close()
}


func(suite *taskUsecaseSuite) GetAllTasks(){
	suite.taskRepository.On("GetAll", mock.Anything).Return([]domain.Task{}, nil)
    tasks , err := suite.taskUsecase.GetAll(context.TODO())

	suite.NoError(err, "no error")
	suite.Equal(0 , len(tasks))
	suite.taskRepository.AssertCalled(suite.T(), "GetAll", mock.Anything)
}

func (suite *taskUsecaseSuite)GetByID(){
	task := domain.Task{}

	suite.taskRepository.On("GetByID" , mock.Anything ,task.ID).Return(&task , nil )
	_ , err := suite.taskUsecase.GetByID(context.TODO() , task.ID)
	suite.NoError(err , "no error")
	suite.taskRepository.AssertCalled(suite.T() , "GetByID" , mock.Anything , task.ID)
}

func (suite *taskUsecaseSuite)CreateTask_Positive(){
	task :=&domain.Task{
      Title: "documentation",
	  Description: "write documentation",
	  DueDate: "2021-09-01",
	  Status: "completed",
	}

	suite.taskRepository.On("Create" , mock.Anything , task).Return(task , nil)
	suite.taskRepository.On("GetByID", mock.Anything, task.ID).Return(domain.Task{}, nil)
	NewTask , err := suite.taskUsecase.Create(context.TODO() , task)

	suite.NoError(err , "no error")
	suite.Equal(task , NewTask)
	suite.taskRepository.AssertCalled(suite.T() , "Create" , mock.Anything , task)
	suite.taskRepository.AssertCalled(suite.T() , "GetByID" , mock.Anything , task.ID)
}

func (suite *taskUsecaseSuite)CreateTask_Negative(){
	task :=&domain.Task{
	  Title: "documentation",
	  Description: "write documentation",
	  DueDate: "2021-09-01",
	  Status: "completed",
	}
	error := errors.New("error")
	suite.taskRepository.On("Create" , mock.Anything , task).Return(nil , nil)
	suite.taskRepository.On("GetById" , mock.Anything , task.ID).Return(nil, error)
	NewTask , err := suite.taskUsecase.Create(context.TODO() , task)

	suite.Error(err , "error")
	suite.Nil(NewTask)
	suite.taskRepository.AssertCalled(suite.T() , "Create" , mock.Anything , task)
	suite.taskRepository.AssertCalled(suite.T() , "GetById" , mock.Anything , task)

}

func (suite *taskUsecaseSuite)Update_Positive (){
	task := &domain.Task{
		ID : 1,
		Title: "update doc",
		Description: "update my documentation",
	}

	suite.taskRepository.On("Update" , mock.Anything , task.ID , task).Return(task , nil)
	_ , err := suite.taskUsecase.Update(context.TODO() , task.ID , task)

	suite.NoError(err , " no error")
	suite.taskRepository.AssertCalled(suite.T() , "Update" , task.ID , task) 
	
}

func (suite *taskUsecaseSuite) Update_Negative() {
    task := &domain.Task{
        ID:          1,
        Title:       "update doc",
        Description: "update my documentation",
    }

    suite.taskRepository.On("Update", mock.Anything, task.ID, task).Return(nil, errors.New("update failed"))
    _, err := suite.taskUsecase.Update(context.TODO(), task.ID, task)


    suite.Error(err, "expected an error")
    suite.EqualError(err, "update failed", "error message should match")
    suite.taskRepository.AssertCalled(suite.T(), "Update", task.ID, task)
}

func (suite *taskUsecaseSuite)Delete_Positive(){
	taskID := 1

	suite.taskRepository.On("Delete" , mock.Anything , taskID).Return(nil)
	err:= suite.taskUsecase.Delete(context.TODO() , taskID)

	suite.NoError(err , "no error")
	suite.taskRepository.AssertCalled(suite.T() , "Delete" , mock.Anything , taskID)
}

func (suite *taskUsecaseSuite)Delete_Negative(){
	taskID := 1

	suite.taskRepository.On("Delete" , mock.Anything , taskID).Return(errors.New("delete failed"))
	err:= suite.taskUsecase.Delete(context.TODO() , taskID)

	suite.Error(err , "error")
	suite.EqualError(err , "delete failed" , "error message should match")
	suite.taskRepository.AssertCalled(suite.T() , "Delete" , mock.Anything , taskID)
}

func TestUsecaseSuite(t *testing.T) {
	suite.Run(t, new(taskUsecaseSuite))
}
