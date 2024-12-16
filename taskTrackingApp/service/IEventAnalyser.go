package service

import (
	"lld/taskTrackingApp/entity"
	"time"
)

type IEventAnalyser interface {
	IObserver
	GetAllActivity() []entity.Event
	GetCompletedEvents(start *time.Time, end *time.Time) []entity.Event
}
