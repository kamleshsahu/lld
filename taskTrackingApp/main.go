package main

import (
	"fmt"
	"lld/taskTrackingApp/entity"
	"lld/taskTrackingApp/service"
	"time"
)

func main() {

	taskManager := service.NewTaskManager()
	eventObserver := service.NewEventLogger()
	taskManager.Subscribe(eventObserver)

	taskList := []entity.Task{
		{Name: "T1", Description: "temp task 1"},
		{Name: "T2", Description: "temp task 2"},
		{Name: "T3", Description: "temp task 3"},
		{Name: "T4", Description: "temp task 4"},
		{Name: "T5", Description: "temp task 5"},
	}

	for i, task := range taskList {
		t, err := taskManager.AddTask(task)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(t)
		}
		taskList[i] = *t
		time.Sleep(500 * time.Millisecond)
	}

	err := taskManager.DeleteTask(taskList[2].Id)
	if err != nil {
		return
	}

	t2Updated, err := taskManager.UpdateTaskStatus(taskList[2].Id, entity.COMPLETED)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t2Updated)
	}

	t3Updated, err := taskManager.UpdateTaskStatus(taskList[3].Id, entity.COMPLETED)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t3Updated)
	}

	activityList := eventObserver.GetAllActivity()
	for _, activity := range activityList {
		fmt.Printf("%s %s %d %s\n", activity.Time.String(), activity.TaskName, activity.TaskId, activity.Action)
	}

	fmt.Println("events which are within given time")
	end := time.Now()
	start := time.Now().Add(-1 * time.Second)
	activityList = eventObserver.GetCompletedEvents(&start, &end)
	for _, activity := range activityList {
		fmt.Printf("%s %s %d %s\n", activity.Time.String(), activity.TaskName, activity.TaskId, activity.Action)
	}
}
