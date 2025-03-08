package main

import (
	"fmt"
	"lld/go-pubsub-queue/pubsub/model"
	"time"
)

type SleepingSubscriber struct {
	id              string
	sleepTimeInMsec int
}

// NewSleepingSubscriber creates a new sleeping subscriber
func NewSleepingSubscriber(id string, sleepTimeInMsec int) model.Subscriber {
	return &SleepingSubscriber{
		id:              id,
		sleepTimeInMsec: sleepTimeInMsec,
	}
}

// GetID returns the subscriber's ID
func (s *SleepingSubscriber) GetID() string {
	return s.id
}

// Consume processes a message with a deliberate sleep
func (s *SleepingSubscriber) Consume(message *model.Message) error {
	fmt.Printf("Subscriber: %s started consuming: %s\n", s.id, message.GetMsg())
	time.Sleep(time.Duration(s.sleepTimeInMsec) * time.Millisecond)
	fmt.Printf("Subscriber: %s done consuming: %s\n", s.id, message.GetMsg())
	return nil
}
