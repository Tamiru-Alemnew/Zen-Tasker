package repositories

import (
	"context"
	"fmt"

	domain "github.com/Tamiru-Alemnew/task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository struct{
	Database *mongo.Database
	Collection string
}

func NewTaskRepository(db *mongo.Database , collection string) domain.TaskRepository{
	return &TaskRepository{
		Database: db,
		Collection: collection,
	}
}

func(tr *TaskRepository) Create(c context.Context, task *domain.Task) error{
	collection := tr.Database.Collection(tr.Collection)
	_ , err := collection.InsertOne(c , task)

	return err
}

func(tr *TaskRepository) Update(c context.Context ,id int , task *domain.Task) error{
	collection := tr.Database.Collection(tr.Collection)
	filter := bson.M{"_id": id}
	update := bson.D{{Key: "$set", Value: task}}

	_ , err := collection.UpdateOne(c , filter , update)

	return err
	
}

func(tr *TaskRepository) Delete (c context.Context , id int ) error{
	collection := tr.Database.Collection(tr.Collection)
	filter := bson.M{"_id" : id}
	_ , err := collection.DeleteOne( c , filter)

	return err
}

func (tr *TaskRepository)GetAll(c context.Context)([]domain.Task , error){
	var tasks []domain.Task
	collection := tr.Database.Collection(tr.Collection)

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

func (tr *TaskRepository) GetByID (c context.Context, id int )(*domain.Task , error){
	var task domain.Task

	collection := tr.Database.Collection(tr.Collection)
	filter := bson.M{"_id":id}

	err := collection.FindOne(c , filter).Decode(&task)

	if err != nil {
    if err == mongo.ErrNoDocuments {
        return nil, fmt.Errorf("task with id %d not found", id)
    }
    return nil, err
	}

	return &task , nil

}