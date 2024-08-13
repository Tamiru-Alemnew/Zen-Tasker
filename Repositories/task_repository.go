package repositories

import (
	"context"
	"fmt"

	"github.com/Tamiru-Alemnew/task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/internal/errors"
)

type taskRepository struct{
	database mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database , collection string) domain.TaskRepository{
	return &taskRepository{
		database: db,
		collection: collection,
	}
}

func(tr *taskRepository) Create(c context.Context, task *domain.Task) error{
	collection := tr.database.Collection(tr.collection)

	_ , err := collection.InsertOne(c , task)

	return err
}

func(tr *taskRepository) Update(c context.Context ,id int , task *domain.Task) error{
	collection := tr.database.Collection(tr.collection)
	filter := bson.M{"id": id}
	update := bson.D{{Key: "$set", Value: task}}

	result , err := collection.UpdateOne(c , filter , update)

	if result.MatchedCount == 0 {
		return errors.New("no task associated with this id")
	}
	return err
	
}

func(tr *taskRepository) Delete (c context.Context , id int ) error{
	collection := tr.database.Collection(tr.collection)
	filter := bson.M{"id" : id}
	result , err := collection.DeleteOne( c , filter)

	if result.DeletedCount == 0 {
        return errors.New("task not found")
    }
	return err
}

func (tr *taskRepository)GetAll(c context.Context)([]domain.Task , error){
	var tasks []domain.Task
	collection := tr.database.Collection(tr.collection)

	cursor , err := collection.Find(c , bson.D{})

	if err != nil{
		return nil , err
	}
	defer cursor.Close(c)
	
	for cursor.Next(c){
		var task domain.Task

		err := cursor.Decode(&task)
		
		if err != nil {
			return nil , err
		}
		tasks = append(tasks, task)
	}


	return tasks , nil

}

func (tr *taskRepository) GetByID (c context.Context, id int )(*domain.Task , error){
	var task domain.Task

	collection := tr.database.Collection(tr.collection)
	filter := bson.M{"id":id}

	err := collection.FindOne(c , filter).Decode(&task)

	if err != nil {
    if err == mongo.ErrNoDocuments {
        return nil, fmt.Errorf("task with id %d not found", id)
    }
    return nil, err
	}

	return &task , nil

}