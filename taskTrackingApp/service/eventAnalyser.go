package service

import (
	"lld/taskTrackingApp/entity"
	"time"
)

type EventAnalyser struct {
	activityLog []entity.Event
}

func (e *EventAnalyser) GetAllActivity() []entity.Event {
	events := make([]entity.Event, 0)

	for _, event := range e.activityLog {
		events = append(events, event)
	}

	return events
}

func (e *EventAnalyser) GetCompletedEvents(start *time.Time, end *time.Time) []entity.Event {
	events := make([]entity.Event, 0)

	for _, event := range e.activityLog {
		if (start != nil && event.Time.Before(*start)) || (end != nil && event.Time.After(*end)) {
			continue
		}
		events = append(events, event)
	}

	return events
}

func (e *EventAnalyser) Notify(data interface{}) error {
	event := data.(entity.Event)
	event.Time = time.Now()
	e.activityLog = append(e.activityLog, event)
	//fmt.Printf("new event captured : %v\n", event)
	return nil
}

func NewEventLogger() IEventAnalyser {
	return &EventAnalyser{}
}
