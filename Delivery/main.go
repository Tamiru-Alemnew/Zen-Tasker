package main

import (
	"context"
	"log"
	"os"

	"github.com/Tamiru-Alemnew/task-manager/Delivery/router"
	infrastructure "github.com/Tamiru-Alemnew/task-manager/Infrastructures"
	"github.com/Tamiru-Alemnew/task-manager/Repositories"
	usecases "github.com/Tamiru-Alemnew/task-manager/Usecases"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

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

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		log.Fatalf("MONGODB_URI not set in .env file")
	}

	jwtKey := os.Getenv("JWT_SECRET")
	if jwtKey == "" {
		log.Fatalf("JWT secret key is not set")
	}

	// Initialize MongoDB client
	client, err := InitMongoDB(mongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
	}()

	// Initialize MongoDB collections
	taskCollection := client.Database("taskmanager")
	userCollection := client.Database("taskmanager")

	// Initialize repositories
	taskRepo := repositories.NewTaskRepository(taskCollection , "tasks")
	userRepo := repositories.NewUserRepository(userCollection , "users")

	// Initialize usecases
	taskUsecase := usecases.NewTaskUsecase(taskRepo)
	userUsecase := usecases.NewUserUsecase(userRepo , infrastructure.NewPasswordService(bcrypt.DefaultCost) , infrastructure.NewJWTService(os.Getenv("JWT_SECRET")))
	
	// Setup router with dependencies injected
	r := router.SetupRouter(taskUsecase, userUsecase)

	r.Run(":8080")
}
