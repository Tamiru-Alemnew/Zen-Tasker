package tests

import (
	"context"
	"log"
	"testing"

	"os"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"

	// "errors"

	"github.com/Tamiru-Alemnew/task-manager/Domain"
	repositories "github.com/Tamiru-Alemnew/task-manager/Repositories"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type UserRepositorySuite struct {
	suite.Suite
	DB            *mongo.Database
	UserRepository domain.UserRepository
	Ctx           context.Context
	Collection    string
}

func (suite *UserRepositorySuite) SetupSuite() {
	err := godotenv.Load("../../.env")
	suite.NoError(err)

	mongoURI := os.Getenv("MONGODB_URI")
	suite.Require().NotEmpty(mongoURI, "MONGODB_URI must be set in the environment")

	clientOpts := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	suite.NoError(err)

	err = client.Ping(context.TODO(), readpref.Primary())
	suite.NoError(err)

	suite.Collection = "users_test"
	suite.DB = client.Database("task_manager_test")
	suite.UserRepository = repositories.NewUserRepository(suite.DB, suite.Collection)
}

func (suite *UserRepositorySuite) TearDownSuite() {
	err := suite.DB.Collection(suite.Collection).Drop(suite.Ctx)
	suite.NoError(err)
}

func (suite *UserRepositorySuite) TestFindByUsername() {
	user := domain.User{
		ID:       1,
		Username: "testuser",
		Password: "password123",
		Role:     "user",
	}

	_, err := suite.DB.Collection(suite.Collection).InsertOne(suite.Ctx, user)
	suite.NoError(err)

	result, err := suite.UserRepository.FindByUsername(suite.Ctx, "testuser")
	suite.NoError(err)
	suite.NotNil(result)
	suite.Equal(user.Username, result.Username)

	// Test case: User not found
	result, err = suite.UserRepository.FindByUsername(suite.Ctx, "nonexistentuser")
	suite.NoError(err)
	suite.Nil(result)
}

func (suite *UserRepositorySuite) TestCreate() {
	user := domain.User{
		ID:       2,
		Username: "newuser",
		Password: "password456",
		Role:     "user",
	}

	err := suite.UserRepository.Create(suite.Ctx, &user)
	suite.NoError(err)

	// Verify the user was created
	var result domain.User
	err = suite.DB.Collection(suite.Collection).FindOne(suite.Ctx, bson.M{"username": user.Username}).Decode(&result)
	suite.NoError(err)
	suite.Equal(user.Username, result.Username)
}

func (suite *UserRepositorySuite) TestGetAll() {
	users := []domain.User{
		{ID: 3, Username: "user1", Password: "pass1", Role: "user"},
		{ID: 4, Username: "user2", Password: "pass2", Role: "user"},
	}

	_, err := suite.DB.Collection(suite.Collection).InsertMany(suite.Ctx, []interface{}{users[0], users[1]})
	suite.NoError(err)

	result, err := suite.UserRepository.GetAll(suite.Ctx)
	suite.NoError(err)
	suite.Equal(4, len(result))
	suite.Equal(users[0].Username, result[2].Username)
	suite.Equal(users[1].Username, result[3].Username)
}

// test promote

func (suite *UserRepositorySuite) TestPromote() {
    user := domain.User{
        ID:       5,
        Username: "user5",
    	Password: "pass5",
        Role:     "user",
    }

    // Create the user
    err := suite.UserRepository.Create(suite.Ctx, &user)
    suite.NoError(err, "should create user without error")

    // Log the created user
    var createdUser domain.User
    err = suite.DB.Collection(suite.Collection).FindOne(suite.Ctx, bson.M{"_id": user.ID}).Decode(&createdUser)
    log.Println("Created User:", createdUser)
    suite.NoError(err, "should find created user without error")
    suite.Equal("user", createdUser.Role, "initial role should be 'user'")

    err = suite.UserRepository.Promote(suite.Ctx, user.ID)
    suite.NoError(err, "should promote user without error")

    var result domain.User
    err = suite.DB.Collection(suite.Collection).FindOne(suite.Ctx, bson.M{"_id": user.ID}).Decode(&result)
   
    suite.NoError(err, "should find promoted user without error")
    suite.Equal("admin", result.Role, "role should be 'admin' after promotion")
}



func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
