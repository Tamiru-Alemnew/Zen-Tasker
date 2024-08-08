package main

import (
    "github.com/Tamiru-Alemnew/task-manager/router"
    "context"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
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

    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
      panic(err)
    }

    defer func() {
      if err = client.Disconnect(context.TODO()); err != nil {
        panic(err)
      }
    }()


    // Send a ping to confirm a successful connection
    if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
      panic(err)
    }

    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

    r := router.SetupRouter()

    r.Run(":8080") // Run on port 8080
}
