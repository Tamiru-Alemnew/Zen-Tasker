package tests

// import (
// 	"context"
// 	"log"
// 	"os"

// 	"testing"
// 	"time"

// 	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
// 	repositories "github.com/Tamiru-Alemnew/task-manager/Repositories"
// 	"github.com/joho/godotenv"
// 	"github.com/stretchr/testify/suite"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type taskRespositorySuite struct {
// 	suite.Suite
// 	TaskRepository *repositories.TaskRepository
// 	collection     *mongo.Collection
// }

// func (suite *taskRespositorySuite) SetupSuite() {
// 	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
// 	client, connectionErr := mongo.Connect(context.TODO(), clientOptions)
// 	if connectionErr != nil {
// 		log.Fatalf("Error: %v", connectionErr.Error())
// 	}

// 	databse := client.Database(os.Getenv("MONGO_URI"))
// 	collection := Database("taskmanager")
// 	collection.DeleteMany(context.TODO(), bson.D{{}})
// 	suite.collection = collection
// 	SetupTaskCollection(suite.collection)
// 	suite.TaskRepository = &repositories.TaskRepository{Collection: collection}
// }

// func (suite *taskRespositorySuite) SetupTest() {
// 	suite.TaskRepository.Collection.DeleteMany(context.TODO(), bson.D{{}})
// }

// // Tests GetAllTasks without adding any
// func (suite *taskRespositorySuite) TestGetTasks_Empty() {
// 	tasks, err := suite.TaskRepository.GetAllTasks(context.TODO())
// 	suite.NoError(err, "no error when fetching")
// 	suite.Equal(0, len(tasks), "lenght of slice returned is 0 when no objects are added")
// }

// // Tests GetAllTasks after adding two tasks
// func (suite *taskRespositorySuite) TestGetTasks_WithAdds() {
// 	task := domain.Task{
// 		ID:         1,
// 		Title:       "title",
// 		Description: "description 1",
// 		DueDate:     "2021-07-01",
// 		Status:      "pending",
// 	}

// 	err := suite.TaskRepository.AddTask(context.TODO(), task)
// 	suite.NoError(err, "no error when creating")
// 	task.ID = 1
// 	err = suite.TaskRepository.AddTask(context.TODO(), task)
// 	suite.NoError(err, "no error when creating")

// 	tasks, err := suite.TaskRepository.GetAllTasks(context.TODO())
// 	suite.NoError(err, "no error when creating")
// 	suite.Equal(2, len(tasks), "lenght of slice returned is 0 when no objects are added")
// }

// // Test GetTaskById after adding a task
// func (suite *taskRespositorySuite) TestGetByID() {
// 	task := domain.Task{
// 		ID:          1,
// 		Title:       "title",
// 		Description: "description 1",
// 		DueDate:     "2021-07-01",
// 		Status:      "pending",
// 	}

// 	err := suite.TaskRepository.AddTask(context.TODO(), task)
// 	suite.NoError(err, "no error when creating")
// 	foundTask, err := suite.TaskRepository.GetTaskByID(context.TODO(), task.ID)
// 	suite.NoError(err, "no error when fetching")
// 	suite.Equal(task.ID, foundTask.ID, "id of the two tasks match")
// }

// // Tests AddTask
// func (suite *taskRespositorySuite) TestAddTask_Positive() {
// 	task := domain.Task{
// 		ID:          1,
// 		Title:       "title",
// 		Description: "description 1",
// 		DueDate:     "2021-07-01",
// 		Status:      "pending",
// 	}

// 	err := suite.TaskRepository.AddTask(context.TODO(), task)
// 	suite.NoError(err, "no error when creating")
// }

// // Tests whether the second ID has been made unique in the DB
// func (suite *taskRespositorySuite) TestAddTask_Negative() {
// 	task := domain.Task{
// 		ID:          1,
// 		Title:       "title",
// 		Description: "description 1",
// 		DueDate:     "2021-07-01",
// 		Status:      "pending",
// 	}

// 	err := suite.TaskRepository.AddTask(context.TODO(), task)
// 	suite.NoError(err, "no error when creating")
// 	err = suite.TaskRepository.AddTask(context.TODO(), task)
// 	suite.Error(err, "error when creating an object with the same id")
// }

// // Tests UpdateTask
// func (suite *taskRespositorySuite) TestUpdateTask() {
// 	task := domain.Task{
// 		ID:          1,
// 		Title:       "title",
// 		Description: "description 1",
// 		DueDate:     "2021-07-01",
// 		Status:      "pending",
// 	}

// 	taskUpdates := domain.Task{
// 		Title:       "changed title",
// 		Description: "changed description 2",
// 		Status:      "completed",
// 	}

// 	err := suite.TaskRepository.AddTask(context.TODO(), task)
// 	suite.NoError(err, "no error when creating")

// 	updatedTask, err := suite.TaskRepository.UpdateTask(context.TODO(), task.ID, taskUpdates)
// 	suite.NoError(err, "no error when updating")

// 	suite.Equal(taskUpdates.Title, updatedTask.Title, "title updated successfully")
// 	suite.Equal(taskUpdates.Description, updatedTask.Description, "description updated successfully")
// 	suite.Equal(taskUpdates.Status, updatedTask.Status, "status updated successfully")
// }

// // test DeleteTask
// func (suite *taskRespositorySuite) TestDeleteTask() {
// 	task := domain.Task{
// 		ID:         1,
// 		Title:       "title",
// 		Description: "description 1",
// 		DueDate:     "2021-07-01",
// 		Status:      "pending",
// 	}

// 	err := suite.TaskRepository.AddTask(context.TODO(), task)
// 	suite.NoError(err, "no error when creating")

// 	err = suite.TaskRepository.DeleteTask(context.TODO(), task.ID)
// 	suite.NoError(err, "no error when deleting")
// 	_, err = suite.TaskRepository.GetTaskByID(context.TODO(), task.ID)
// 	suite.Error(err, "deleted task not found")
// }

// func TestTaskRepositorySuite(t *testing.T) {
// 	err := godotenv.Load("../../.env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	suite.Run(t, new(taskRespositorySuite))
// }