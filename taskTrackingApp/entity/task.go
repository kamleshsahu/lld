package entity

import "time"

type Status string
type Tag string

var (
	TODO       Status = "TODO"
	COMPLETED  Status = "COMPLETED"
	INPROGRESS Status = "INPROGRESS"
	DELETED    Status = "DELETED"
)

var (
	HIGHPRIORITY Tag = "HIGHPRIORITY"
	TRIP         Tag = "TRIP"
)

type EventLog struct {
	Status Status
	Time   time.Time
}

type Task struct {
	Name              string
	Id                int
	Description       string
	EventLog          []EventLog
	ExpectedStartDate time.Time
	ExpectedEndDate   time.Time
	Status            Status
	Tags              map[Tag]bool
}

func (t *Task) UpdateStatus(status Status) {
	t.Status = status
	t.AddEventLog(status)
}

func (t *Task) AddEventLog(status Status) {
	t.EventLog = append(t.EventLog, EventLog{Status: status, Time: time.Now()})
}

func (t *Task) Clone() *Task {
	return &Task{
		Name:              t.Name,
		Id:                t.Id,
		Description:       t.Description,
		Status:            t.Status,
		EventLog:          t.EventLog,
		ExpectedStartDate: t.ExpectedStartDate,
		ExpectedEndDate:   t.ExpectedEndDate,
	}
}
