package data

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Tamiru-Alemnew/task-manager/data"
	"github.com/Tamiru-Alemnew/task-manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    taskCollection *mongo.Collection
    client         *mongo.Client
)

func InitMongoDB(mongoURI string) {
    clientOptions := options.Client().ApplyURI(mongoURI)
    var err error
    client, err = mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
    }

    taskCollection = client.Database("taskmanager").Collection("tasks")
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

func GetAllTasks() ([]models.Task, error) {
    var tasks []models.Task
    cursor, err := taskCollection.Find(context.TODO(), bson.D{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.TODO())

    for cursor.Next(context.TODO()) {
        var task models.Task
        err := cursor.Decode(&task)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return tasks, nil
}

func GetTaskByID(id int) (models.Task, error) {
    var task models.Task
    filter := bson.D{primitive.E{Key: "id", Value: id}}
    err := taskCollection.FindOne(context.TODO(), filter).Decode(&task)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return models.Task{}, errors.New("task not found")
        }
        return models.Task{}, err
    }
    return task, nil
}

func CreateTask(task models.Task) (models.Task, error) {
    task.ID = int(time.Now().UnixNano())
    _, err := taskCollection.InsertOne(context.TODO(), task)
    if err != nil {
        return models.Task{}, err
    }
    return task, nil
}

func UpdateTask(id int, newTask models.Task) (models.Task, error) {
    filter := bson.D{primitive.E{Key: "id", Value: id}}
    update := bson.D{primitive.E{Key: "$set", Value: newTask}}
    result, err := data.TaskCollection.UpdateOne(context.TODO(), filter, update)
    if err != nil {
        return models.Task{}, err
    }
    if result.MatchedCount == 0 {
        return models.Task{}, errors.New("task not found")
    }
    return newTask, nil
}

func DeleteTask(id int) error {
    filter := bson.D{primitive.E{Key: "id", Value: id}}
    result, err := data.TaskCollection.DeleteOne(context.TODO(), filter)
    if err != nil {
        return err
    }
    if result.DeletedCount == 0 {
        return errors.New("task not found")
    }
    return nil
}