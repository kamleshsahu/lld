package handler

import (
	"sync"

	"lld/go-pubsub-queue/pubsub/model"
)

// TopicHandler manages message distribution for a topic
type TopicHandler struct {
	topic             *model.Topic
	subscriberWorkers map[string]*SubscriberWorker
	mutex             sync.RWMutex
}

// NewTopicHandler creates a new topic handler
func NewTopicHandler(topic *model.Topic) *TopicHandler {
	return &TopicHandler{
		topic:             topic,
		subscriberWorkers: make(map[string]*SubscriberWorker),
	}
}

// Publish distributes messages to all subscribers
func (th *TopicHandler) Publish() {
	subscribers := th.topic.GetSubscribers()
	for _, subscriber := range subscribers {
		th.StartSubscriberWorker(subscriber)
	}
}

// StartSubscriberWorker creates and starts a subscriber worker if needed
func (th *TopicHandler) StartSubscriberWorker(topicSubscriber *model.TopicSubscriber) {
	subscriberID := topicSubscriber.GetSubscriber().GetID()

	th.mutex.Lock()

	if _, exists := th.subscriberWorkers[subscriberID]; !exists {
		subscriberWorker := NewSubscriberWorker(th.topic, topicSubscriber)
		th.subscriberWorkers[subscriberID] = subscriberWorker
		subscriberWorker.Start()
	}

	subscriberWorker := th.subscriberWorkers[subscriberID]
	th.mutex.Unlock()

	subscriberWorker.WakeUpIfNeeded()
}
