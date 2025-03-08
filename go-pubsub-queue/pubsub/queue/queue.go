package queue

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"lld/go-pubsub-queue/pubsub/handler"
	"lld/go-pubsub-queue/pubsub/model"
)

// Queue is the main pubsub queue implementation
type Queue struct {
	topicHandlers map[string]*handler.TopicHandler
	mutex         sync.RWMutex
}

// NewQueue creates a new pubsub queue
func NewQueue() *Queue {
	return &Queue{
		topicHandlers: make(map[string]*handler.TopicHandler),
	}
}

// CreateTopic creates a new topic in the queue
func (q *Queue) CreateTopic(topicName string) *model.Topic {
	topicID := uuid.New().String()
	topic := model.NewTopic(topicName, topicID)

	topicHandler := handler.NewTopicHandler(topic)

	q.mutex.Lock()
	q.topicHandlers[topic.GetTopicID()] = topicHandler
	q.mutex.Unlock()

	fmt.Printf("Created topic: %s\n", topic.GetTopicName())
	return topic
}

// Subscribe adds a subscriber to a topic
func (q *Queue) Subscribe(sub model.Subscriber, topic *model.Topic) {
	topicSubscriber := model.NewTopicSubscriber(sub)
	topic.AddSubscriber(topicSubscriber)
	fmt.Printf("%s subscribed to topic: %s\n", sub.GetID(), topic.GetTopicName())
}

// Publish sends a message to a topic
func (q *Queue) Publish(topic *model.Topic, message *model.Message) {
	topic.AddMessage(message)
	fmt.Printf("%s published to topic: %s\n", message.GetMsg(), topic.GetTopicName())

	q.mutex.RLock()
	topicHandler := q.topicHandlers[topic.GetTopicID()]
	q.mutex.RUnlock()

	go topicHandler.Publish()
}

// ResetOffset resets a subscriber's offset position in a topic
func (q *Queue) ResetOffset(topic *model.Topic, sub model.Subscriber, newOffset int) {
	for _, topicSubscriber := range topic.GetSubscribers() {
		if topicSubscriber.GetSubscriber().GetID() == sub.GetID() {
			topicSubscriber.GetOffset().Store(int32(newOffset))
			fmt.Printf("%s offset reset to: %d\n", sub.GetID(), newOffset)

			q.mutex.RLock()
			topicHandler := q.topicHandlers[topic.GetTopicID()]
			q.mutex.RUnlock()

			go func() {
				topicHandler.StartSubscriberWorker(topicSubscriber)
			}()
			break
		}
	}
}
