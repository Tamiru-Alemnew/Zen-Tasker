package data

import (
	"context"
	"errors"
	// "log"
	// "net/http"
	"time"

	"github.com/Tamiru-Alemnew/task-manager/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("jwt_secret_key")


func UserRegistration(user models.User) error {

    // Check if the user already exists
    filter := bson.D{primitive.E{Key: "username", Value: user.Username}}
    var existingUser models.User
    err := UserCollection.FindOne(context.TODO(), filter).Decode(&existingUser)
    if err == nil {
        return errors.New("user already exists")
    }


    // Proceed with registration
    user.ID = int(time.Now().UnixNano())
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return  errors.New("internal server error")
    }

    user.Password = string(hashedPassword)

    _, err = UserCollection.InsertOne(context.TODO(), user)
    if err != nil {
        return  err
    }

    return  nil
}


func UserCredentialValidation(user models.User) (string, error) {
    var logedinUser models.User
    filter := bson.D{primitive.E{Key: "username", Value: user.Username}}

    err := UserCollection.FindOne(context.TODO(), filter).Decode(&logedinUser)
    if err != nil {
        return "", err
    }

    err = bcrypt.CompareHashAndPassword([]byte(logedinUser.Password), []byte(user.Password))
    if err != nil {
        return "", errors.New("invalid email or password")
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  logedinUser.ID,
        "username": logedinUser.Username,
        "exp":      time.Now().Add(time.Hour * 365).Unix(),
    })

    jwtToken, err := token.SignedString(jwtKey)
    if err != nil {
        return "", errors.New("internal server error")
    }

    return jwtToken, nil
}
