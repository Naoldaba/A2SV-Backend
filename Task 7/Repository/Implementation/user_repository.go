package implemenation

import (
	"context"
	"errors"

	"task_manager_api/Domain"
	"task_manager_api/Repository/Interfaces"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *UserRepository) PromoteUser(id string) (*domain.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }
	
    filter := bson.M{"_id": objID}
    update := bson.M{"$set": bson.M{"role": "ADMIN"}}
    opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

    var updatedUser domain.User
    err = repo.collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&updatedUser)
    if err != nil {
        return nil, err
    }
    return &updatedUser, nil
}

func (repo *UserRepository) GetAllUsers() ([]*domain.User, error) {
    var users []*domain.User
    cursor, err := repo.collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background()) 

    err = cursor.All(context.TODO(), &users)
    if err != nil {
        return nil, err
    }

    return users, nil
}
