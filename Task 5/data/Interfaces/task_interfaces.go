package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
	"task_manager_api/models"
)


type ITaskService interface{
    GetTasks(c *mongo.Collection) ([]models.Task, error)
    GetTaskById(c *mongo.Collection, id string) 
    AddTask(task models.Task) (models.Task, error)
    UpdateTask(id string, updatedTask models.Task) (models.Task, error)
    DeleteTask(id string) error
}