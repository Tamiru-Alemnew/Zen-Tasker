package data

import (
	"context"
	"errors"
	"time"

	"github.com/Tamiru-Alemnew/task-manager/models"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)
var JwtKey []byte

func SetJWTKey(key []byte) {
    JwtKey = key
}

func UserRegistration(user models.User) (models.User,error) {

    // Check if the user already exists
    filter := bson.D{primitive.E{Key: "username", Value: user.Username}}
    var existingUser models.User
    err := UserCollection.FindOne(context.TODO(), filter).Decode(&existingUser)
    if err == nil {
        return models.User{} , errors.New("user already exists")
    }

    // Proceed with registration
    user.ID = int(time.Now().UnixNano())
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return models.User{} , errors.New("internal server error")
    }

    user.Password = string(hashedPassword)

    count , err := UserCollection.CountDocuments(context.TODO(), bson.D{})

    if err != nil {
        return models.User{} ,  errors.New("internal server error")
    }

    if count == 0 {
        user.Role = "admin"
    }else {
        user.Role = "user"
    }

    _, err = UserCollection.InsertOne(context.TODO(), user)
    if err != nil {
        return models.User{} ,  err
    }

    return user ,  nil
}


func UserCredentialValidation(user models.User) (models.User, string, error) {
    var logedinUser models.User
    filter := bson.D{primitive.E{Key: "username", Value: user.Username}}

    err := UserCollection.FindOne(context.TODO(), filter).Decode(&logedinUser)
    if err != nil {
        return models.User{} ,"", err
    }

    err = bcrypt.CompareHashAndPassword([]byte(logedinUser.Password), []byte(user.Password))
    if err != nil {
        return models.User{} , "", errors.New("invalid email or password")
    }



    if len(JwtKey) == 0 {
    return models.User{} , "", errors.New("JWT secret key is not set")
}

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id":  logedinUser.ID,
        "role": logedinUser.Role,
        "username": logedinUser.Username,
        "exp":      time.Now().Add(time.Hour * 365).Unix(),
    })

    jwtToken, err := token.SignedString([]byte(JwtKey))
    if err != nil {
        return models.User{} , "", err
    }

    return logedinUser, jwtToken, nil
}

func PromoteUser(id int) (error){

    filter := bson.D{primitive.E{Key: "_id", Value: id}}

    _, err := UserCollection.UpdateOne(
        context.TODO(),
        filter,
        bson.D{
            {Key: "$set", Value: bson.D{
                {Key: "role", Value: "admin"},
            }},
        },
    )

    if err != nil {
        return err
    }

    return nil
}


