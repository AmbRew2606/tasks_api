package storage

import "tz_todo_list_1/internal/models"

//Методы интерфейса Storage
type Storage interface {
	GetTasks() ([]models.Task, error)
	CreateTask(t models.Task) (int, error)
	UpdateTask(t models.Task) error
	DeleteTask(taskID int) error
	Close()
}
