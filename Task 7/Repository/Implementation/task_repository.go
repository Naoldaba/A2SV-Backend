package implemenation

import (
	"context"
	"errors"
	"log"
	"os"

	"task_manager_api/Domain"
	"task_manager_api/Repository/Interfaces"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository struct{
	collection *mongo.Collection
}

func NewTaskRepository(client *mongo.Client) interfaces.ITaskRepository{
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	collection := client.Database(os.Getenv("DB_NAME")).Collection("Tasks")
	return &TaskRepository{
		collection: collection,
	}
}

func (repo *TaskRepository) AddTask(task *domain.Task, id string) (*domain.Task, error) {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil { 
		return nil, err
	}
	task.UserID = userID
	result, err := repo.collection.InsertOne(context.TODO(), task)
	if err != nil { 
		return nil, err
	}
	task.ID = result.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (repo *TaskRepository) GetTasks() ([]*domain.Task, error) {
	var tasks []*domain.Task
	cursor, err := repo.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}

	err = cursor.All(context.TODO(), &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repo *TaskRepository) GetTaskById(id string) (*domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	var task *domain.Task
	err = repo.collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}

	return task, err
} 

func (repo *TaskRepository) UpdateTask(id string, updatedTask *domain.Task) (*domain.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid task ID")
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
	result := repo.collection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	var task *domain.Task
	
	err = result.Decode(&task)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("task not found")
	}
	return task, err	
}

func (repo *TaskRepository) DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	_, err = repo.collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err == mongo.ErrNoDocuments {
		return errors.New("task not found")
	}
    
	return err
}



