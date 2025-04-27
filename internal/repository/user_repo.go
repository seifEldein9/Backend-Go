package repository

import (
	"context"
	"fmt"
	"go-backend-app/internal/config"
	"go-backend-app/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func init() {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetMongoURI()))
    if err != nil {
        log.Fatal(err)
    }

    userCollection = client.Database("goappdb").Collection("users")
}

func GetAllUsers() ([]models.User, error) {
    var users []models.User
    cursor, err := userCollection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var user models.User
        cursor.Decode(&user)
        users = append(users, user)
    }
    return users, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}


func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func CreateUser(user models.User) error {
 	existingUser, _ := GetUserByEmail(user.Email)
	if existingUser.Email != "" {
			return fmt.Errorf("user with email %s already exists", user.Email)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
			return err
	}
	user.Password = string(hash)

	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
			if mongo.IsDuplicateKeyError(err) {
					return fmt.Errorf("user with email %s already exists", user.Email)
			}
			return err
	}
	return nil
}