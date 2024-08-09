package data

import (
	"context"
	"errors"
	"task_manager_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskService struct{}

func (s *TaskService) GetTasks(c *mongo.Collection) ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := c.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &tasks)
    if err != nil {
        return nil, err
    }
	return tasks, nil
}

func (s *TaskService) GetTaskById(c *mongo.Collection, id string) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid task ID")
	}

	var task models.Task
	err = c.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return models.Task{}, errors.New("task not found")
	}
	return task, err
}

func (s *TaskService) AddTask(c *mongo.Collection, task models.Task) (models.Task, error) {
	result, err := c.InsertOne(context.TODO(), task)
	if err != nil { 
		return models.Task{}, err
	}
	task.ID = result.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (s *TaskService) UpdateTask(c *mongo.Collection, id string, updatedTask models.Task) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid task ID")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

    filter := bson.M{
        "_id":objID,
    }
	result := c.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var task models.Task
	err = result.Decode(&task)
	if err == mongo.ErrNoDocuments {
		return models.Task{}, errors.New("task not found")
	}
	return task, err
}

func (s *TaskService) DeleteTask(c *mongo.Collection, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	_, err = c.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err == mongo.ErrNoDocuments {
		return errors.New("task not found")
	}
    
	return err
}
