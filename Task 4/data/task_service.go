package data

import (
	"errors"
	"task_manager_api/models"
)

type TaskService struct {
    tasks []models.Task
    nextTaskID int
}


func CreateTaskSerive() *TaskService{
	return &TaskService{
		tasks: []models.Task{},
		nextTaskID: 1,
	}
}


func (s *TaskService) GetTasks() []models.Task{
	return s.tasks
}


func (s *TaskService) GetTaskById(id int) (models.Task, error) {
    for _, t := range s.tasks {
        if t.ID == id {
            return t, nil
        }
    }
    return models.Task{}, errors.New("task not found")
}


func (s *TaskService) AddTask(task models.Task) models.Task {
    task.ID = s.nextTaskID
    s.nextTaskID++
    s.tasks = append(s.tasks, task)
    return task
}


func (s *TaskService) UpdateTask(id int, updatedTask models.Task) (models.Task, error) {
    for i, t := range s.tasks {
        if t.ID == id {
            s.tasks[i].Title = updatedTask.Title
            s.tasks[i].Description = updatedTask.Description
            s.tasks[i].DueDate = updatedTask.DueDate
            s.tasks[i].Status = updatedTask.Status
            return s.tasks[i], nil
        }
    }
    return models.Task{}, errors.New("task not found")
}


func (s *TaskService) DeleteTask(id int) error {
    for i, t := range s.tasks {
        if t.ID == id {
            s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
            return nil
        }
    }
    return errors.New("task not found")
}