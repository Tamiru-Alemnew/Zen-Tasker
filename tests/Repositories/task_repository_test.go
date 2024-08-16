package tests

import (
	"context"
	"log"
	"os"
	"testing"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	repositories "github.com/Tamiru-Alemnew/task-manager/Repositories"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositorySuite struct {
	suite.Suite
	TaskRepository *repositories.TaskRepository
	// collection     *mongo.Collection
}

func (suite *TaskRepositorySuite) SetupSuite() {
	mongoURI := os.Getenv("MONGODB_URI")
	client, err := InitMongoDB(mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	database := client.Database("taskmanager")
	suite.TaskRepository = &repositories.TaskRepository{
		Database:   database,
		Collection: "tasks",
	}
}

func (suite *TaskRepositorySuite) SetupTest() {
	_, err := suite.TaskRepository.Database.Collection(suite.TaskRepository.Collection).DeleteMany(context.TODO(), bson.D{{}})
	suite.NoError(err, "should delete all tasks before each test")
}

// Test GetAll when no tasks are added
func (suite *TaskRepositorySuite) TestGetTasks_Empty() {
	tasks, err := suite.TaskRepository.GetAll(context.TODO())
	suite.NoError(err, "should fetch tasks without error")
	suite.Equal(0, len(tasks), "length of slice returned should be 0 when no tasks are added")
}

// Test GetAll after adding two tasks
func (suite *TaskRepositorySuite) TestGetTasks_WithAdds() {
	task1 := domain.Task{
		ID:          1,
		Title:       "title1",
		Description: "description 1",
		DueDate:     "2021-07-01",
		Status:      "pending",
	}

	task2 := domain.Task{
		ID:          2,
		Title:       "title2",
		Description: "description 2",
		DueDate:     "2021-07-02",
		Status:      "pending",
	}

	err := suite.TaskRepository.Create(context.TODO(), &task1)
	suite.NoError(err, "should create first task without error")
	err = suite.TaskRepository.Create(context.TODO(), &task2)
	suite.NoError(err, "should create second task without error")

	tasks, err := suite.TaskRepository.GetAll(context.TODO())
	suite.NoError(err, "should fetch tasks without error")
	suite.Equal(2, len(tasks), "length of slice returned should be 2 after adding tasks")
}



// Test Create (AddTask)
func (suite *TaskRepositorySuite) TestAddTask_Positive() {
	task := domain.Task{
		ID:          1,
		Title:       "title",
		Description: "description 1",
		DueDate:     "2021-07-01",
		Status:      "pending",
	}

	err := suite.TaskRepository.Create(context.TODO(), &task)
	suite.NoError(err, "should create task without error")
}

// Test Create (AddTask) with duplicate ID

func (suite *TaskRepositorySuite) TestAddTask_DuplicateID() {

	task := domain.Task{
		ID:          1,
		Title:       "title",
		Description: "description 1",
		DueDate:     "2021-07-01",
		Status:      "pending",
	}

	err := suite.TaskRepository.Create(context.TODO(), &task)
	suite.NoError(err, "should create task without error")

	err = suite.TaskRepository.Create(context.TODO(), &task)
	suite.Error(err, "should not create task with duplicate ID")
}

// get by id

func (suite *TaskRepositorySuite) TestGetByID_Negative() {

	_, err := suite.TaskRepository.GetByID(context.TODO(), 200)
	suite.Error(err, "sempty db should return error")
	
}



// Helper function to initialize MongoDB connection
func InitMongoDB(mongoURI string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}


func TestTaskRepositorySuite(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	suite.Run(t, new(TaskRepositorySuite))
}
