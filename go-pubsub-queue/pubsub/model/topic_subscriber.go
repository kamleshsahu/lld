package model

import (
	"sync"
	"sync/atomic"
)

// TopicSubscriber relates a subscriber to a topic with an offset
type TopicSubscriber struct {
	offset     atomic.Int32
	subscriber Subscriber
	mutex      sync.Mutex
	cond       *sync.Cond
}

// NewTopicSubscriber creates a new topic subscriber
func NewTopicSubscriber(sub Subscriber) *TopicSubscriber {
	ts := &TopicSubscriber{
		subscriber: sub,
	}
	ts.offset.Store(0)
	ts.cond = sync.NewCond(&ts.mutex)
	return ts
}

// GetOffset returns the current offset
func (ts *TopicSubscriber) GetOffset() *atomic.Int32 {
	return &ts.offset
}

// GetSubscriber returns the subscriber
func (ts *TopicSubscriber) GetSubscriber() Subscriber {
	return ts.subscriber
}

// Lock locks the subscriber for synchronized operations
func (ts *TopicSubscriber) Lock() {
	ts.mutex.Lock()
}

// Unlock unlocks the subscriber
func (ts *TopicSubscriber) Unlock() {
	ts.mutex.Unlock()
}

// Wait causes the goroutine to wait on the condition
func (ts *TopicSubscriber) Wait() {
	ts.cond.Wait()
}

// Signal wakes up a goroutine waiting on the condition
func (ts *TopicSubscriber) Signal() {
	ts.cond.Signal()
}

// Subscriber interface for consumers of messages
type Subscriber interface {
	GetID() string
	Consume(message *Message) error
}
