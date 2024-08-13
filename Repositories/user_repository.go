package repositories

import (
	"context"
	"errors"

	"github.com/Tamiru-Alemnew/task-manager/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	database   *mongo.Database
	collection string
}

func NewUserRepository(db *mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	var user domain.User
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"username": username}

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (ur *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
    collection := ur.database.Collection(ur.collection)
    var users []domain.User

    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var user domain.User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (ur *userRepository) Promote(ctx context.Context, id int) error {
	collection := ur.database.Collection(ur.collection)
	filter := bson.M{"id": id}

	// Update the user's role to admin
	update := bson.M{"$set": bson.M{"role": "admin"}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("user not found")
		}
		return err
	}

	return nil
}

