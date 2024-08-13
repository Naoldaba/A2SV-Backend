package implemenation

import (
	"context"
	"errors"

	"task_manager_api/Domain"
	"task_manager_api/Repository/Interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct{
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) interfaces.IUserRepository{
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