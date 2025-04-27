package repository

import (
	"context"
	"go-backend-app/internal/config"
	"go-backend-app/internal/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func CreateUser(user models.User) error {
    _, err := userCollection.InsertOne(context.Background(), user)
    return err
}
