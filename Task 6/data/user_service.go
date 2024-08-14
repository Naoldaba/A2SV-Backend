package data

import (
	"context"
	"task_manager_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct{}

func (s *UserService) Register(c *mongo.Collection, user models.User) error {
	_, err := c.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUser(c *mongo.Collection, email string) (*models.User, error) {
    var user models.User
    err := c.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}


func (s *UserService) GetAllUsers(c *mongo.Collection) ([]*models.User, error){
	var users []*models.User
	cursor, err := c.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	} 
	err = cursor.All(context.TODO(), &users)
	if err != nil{
		return nil, err
	}
	return users, nil
}

func (s *UserService) PromoteUser(c *mongo.Collection, id string)  (*models.User, error){
	var existingUser *models.User
	user_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = c.FindOne(context.TODO(), bson.M{"_id": user_id}).Decode(&existingUser)
	if err != nil{
		return nil, err
	}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	var newUser *models.User
	err = c.FindOneAndUpdate(context.TODO(), bson.M{"_id":user_id}, bson.M{"$set": bson.M{"role": "ADMIN"}},opts).Decode(&newUser)
	if err != nil{
		return nil, err
	}
	return newUser, err
}