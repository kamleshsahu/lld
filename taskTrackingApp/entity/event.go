package entity

import "time"

type Event struct {
	Time     time.Time
	TaskId   int
	TaskName string
	Action   string
	TaskMeta Task
}
