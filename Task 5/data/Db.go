package data

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client{

	err := godotenv.Load()
	if err != nil{
		log.Fatal("err loading .env file")
	}

	db_string := os.Getenv("CONNECTION_STRING")

	clientOptions := options.Client().ApplyURI(db_string)
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

var Client *mongo.Client = DbInstance()

func OpenCollection(client *mongo.Client, colName string) *mongo.Collection{
	err := godotenv.Load()
	if err != nil{
		log.Fatal("err loading .env file")
	}
	db_Name := os.Getenv("DB_NAME")
	collection := client.Database(db_Name).Collection(colName)
	return collection
}