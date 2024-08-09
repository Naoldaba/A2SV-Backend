package data

import (
	"context"
	"task_manager_api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct{}

func (s *UserService) Register(c *mongo.Collection, user models.User) error {
	_, err := c.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUser(c *mongo.Collection, email string) (models.User, error) {
	var user models.User
	err := c.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
