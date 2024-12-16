package service

import (
	"container/list"
	"errors"
	"fmt"
	"lld/taskTrackingApp/entity"
	"sync"
)

type TaskManager struct {
	Observable
	taskList   map[int]*entity.Task
	nextTaskId int
}

func (t *TaskManager) GetAllTasks(status *entity.Status) ([]entity.Task, error) {
	taskList := make([]entity.Task, 0)
	for _, task := range t.taskList {
		if status == nil || task.Status == *status {
			taskList = append(taskList, *task.Clone())
		}
	}
	return taskList, nil
}

func (t *TaskManager) AddTask(task entity.Task) (*entity.Task, error) {
	t.nextTaskId++
	task.Id = t.nextTaskId
	task.Status = entity.TODO
	t.taskList[t.nextTaskId] = &task
	taskCopy := task.Clone()
	event := entity.Event{TaskId: task.Id, TaskName: taskCopy.Name, Action: string(task.Status), TaskMeta: *taskCopy}
	t.Fire(event)
	return taskCopy, nil
}

func (t *TaskManager) DeleteTask(id int) error {
	if _, ok := t.taskList[id]; !ok {
		return errors.New("task Not Found")
	}
	taskCopy := t.taskList[id].Clone()
	delete(t.taskList, id)

	event := entity.Event{TaskId: taskCopy.Id, TaskName: taskCopy.Name, Action: string(entity.DELETED), TaskMeta: *taskCopy}
	t.Fire(event)
	return nil
}

func (t *TaskManager) UpdateTaskStatus(id int, status entity.Status) (*entity.Task, error) {
	if _, ok := t.taskList[id]; !ok {
		return nil, errors.New(fmt.Sprintf("taskId : %d not found", id))
	}
	t.taskList[id].UpdateStatus(status)
	taskCopy := t.taskList[id].Clone()

	event := entity.Event{TaskId: taskCopy.Id, TaskName: taskCopy.Name, Action: string(taskCopy.Status), TaskMeta: *taskCopy}
	t.Fire(event)
	return taskCopy, nil
}

var singleton sync.Once
var taskManagerInstance *TaskManager

func NewTaskManager() *TaskManager {
	singleton.Do(func() {
		taskManagerInstance = &TaskManager{taskList: make(map[int]*entity.Task), Observable: Observable{subscribers: list.New()}}
	})
	return taskManagerInstance
}
