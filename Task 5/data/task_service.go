package data

import (
	"context"
	"errors"
	"log"
	"task_manager_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskService struct {
	collection *mongo.Collection
	client     *mongo.Client
}

func CreateTaskService(dbName, colName, connString string) *TaskService {
	clientOptions := options.Client().ApplyURI(connString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(dbName).Collection(colName)

	log.Println("DB connected successfully")
	return &TaskService{
		collection: collection,
		client:     client,
	}
}

func (s *TaskService) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := s.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &tasks)
    if err != nil {
        return nil, err
    }
	return tasks, nil
}

func (s *TaskService) GetTaskById(id string) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid task ID")
	}

	var task models.Task
	err = s.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return models.Task{}, errors.New("task not found")
	}
	return task, err
}

func (s *TaskService) AddTask(task models.Task) (models.Task, error) {
	result, err := s.collection.InsertOne(context.TODO(), task)
	if err != nil {
		return models.Task{}, err
	}
	task.ID = result.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (s *TaskService) UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
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
	result := s.collection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var task models.Task
	err = result.Decode(&task)
	if err == mongo.ErrNoDocuments {
		return models.Task{}, errors.New("task not found")
	}
	return task, err
}

func (s *TaskService) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	_, err = s.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err == mongo.ErrNoDocuments {
		return errors.New("task not found")
	}
    
	return err
}
