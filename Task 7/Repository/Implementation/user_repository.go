package implemenation

import (
	"context"
	"log"
	"os"
	"errors"

	"task_manager_api/Domain"
	"task_manager_api/Repository/Interfaces"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct{
	collection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) interfaces.IUserRepository{
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	collection := client.Database(os.Getenv("DB_NAME")).Collection("Users")
	return &UserRepository{
		collection: collection,
	}
}

func (repo *UserRepository) GetUser(email string) (*domain.User, error) {
	var user *domain.User
	err := repo.collection.FindOne(context.TODO(), bson.M{"email":email}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (repo *UserRepository) Register(user *domain.User) error{
	_, err := repo.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}