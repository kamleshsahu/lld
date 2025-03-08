package handler

import (
	"fmt"

	"lld/go-pubsub-queue/pubsub/model"
)

// SubscriberWorker handles the message consumption for a subscriber
type SubscriberWorker struct {
	topic           *model.Topic
	topicSubscriber *model.TopicSubscriber
}

// NewSubscriberWorker creates a new subscriber worker
func NewSubscriberWorker(topic *model.Topic, topicSubscriber *model.TopicSubscriber) *SubscriberWorker {
	return &SubscriberWorker{
		topic:           topic,
		topicSubscriber: topicSubscriber,
	}
}

// Start begins the worker process
func (sw *SubscriberWorker) Start() {
	go sw.run()
}

// WakeUpIfNeeded signals the subscriber to check for new messages
func (sw *SubscriberWorker) WakeUpIfNeeded() {
	sw.topicSubscriber.Lock()
	defer sw.topicSubscriber.Unlock()
	sw.topicSubscriber.Signal()
}

// run processes messages for the subscriber
func (sw *SubscriberWorker) run() {
	for {
		sw.topicSubscriber.Lock()
		curOffset := int(sw.topicSubscriber.GetOffset().Load())
		messages := sw.topic.GetMessages()

		if curOffset >= len(messages) {
			sw.topicSubscriber.Wait()
			sw.topicSubscriber.Unlock()
			continue
		}

		message := messages[curOffset]
		sw.topicSubscriber.Unlock()

		// Process the message
		err := sw.topicSubscriber.GetSubscriber().Consume(message)
		if err != nil {
			fmt.Printf("Error consuming message: %v\n", err)
			// Consider adding backoff or recovery logic here
		}

		sw.topicSubscriber.Lock()
		// Compare and set offset only if it hasn't changed
		if int(sw.topicSubscriber.GetOffset().Load()) == curOffset {
			sw.topicSubscriber.GetOffset().Add(1)
		}
		sw.topicSubscriber.Unlock()
	}
}
