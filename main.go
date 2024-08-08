package main

import (
	"log"
	"os"

	"github.com/Tamiru-Alemnew/task-manager/data"
	"github.com/Tamiru-Alemnew/task-manager/router"
	"github.com/joho/godotenv"
)

func main() {

   // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatalf("MONGODB_URI not set in .env file")
    }

    data.InitMongoDB(mongoURI)

    // close the connection when the main function ends
    defer data.DisconnectMongoDB()

    r := router.SetupRouter()

    r.Run(":8080") 
}
