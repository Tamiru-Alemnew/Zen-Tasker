package data

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Tamiru-Alemnew/task-manager/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UserRegistration(user models.User)(models.User , error){
	user.ID = int(time.Now().UnixNano())
	_ , err := UserCollection.InsertOne(context.TODO() , user)

	if err != nil {
		return models.User{} , err
	}

	return user , nil
}

func UserLogin(user models.User)(models.User , error){
	filter = bson.D{primitive.E{Key: "username", Value: user.Username}}
	

}
