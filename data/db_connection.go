package data

import (
	"context"
	"errors"
	"log"

	// "github.com/Tamiru-Alemnew/task-manager/models"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
    TaskCollection *mongo.Collection
	UserCollection *mongo.Collection
    client         *mongo.Client
)

func InitMongoDB(mongoURI string) {

    clientOptions :=options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(context.TODO(), clientOptions)

    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    TaskCollection = client.Database("taskmanager").Collection("tasks")
	UserCollection = client.Database("taskmanager").Collection(("users"))
}


func DisconnectMongoDB() error {
    if client == nil {
        return errors.New("no MongoDB client to disconnect")
    }
    err := client.Disconnect(context.TODO())
    if err != nil {
        return err
    }
    client = nil
    return nil
}
