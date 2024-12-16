package service

import "lld/taskTrackingApp/entity"

type ITaskManager interface {
	AddTask(task entity.Task) (*entity.Task, error)
	DeleteTask(id int) error
	UpdateTaskStatus(id int, status entity.Status) (*entity.Task, error)
	GetAllTasks(status *entity.Status) ([]entity.Task, error)
}
