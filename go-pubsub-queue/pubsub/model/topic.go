package model

import (
	"sync"
)

// Topic represents a message topic
type Topic struct {
	topicName   string
	topicID     string
	messages    []*Message
	subscribers []*TopicSubscriber
	mutex       sync.RWMutex
}

// NewTopic creates a new topic
func NewTopic(topicName, topicID string) *Topic {
	return &Topic{
		topicName:   topicName,
		topicID:     topicID,
		messages:    make([]*Message, 0),
		subscribers: make([]*TopicSubscriber, 0),
	}
}

// GetTopicName returns the topic name
func (t *Topic) GetTopicName() string {
	return t.topicName
}

// GetTopicID returns the topic ID
func (t *Topic) GetTopicID() string {
	return t.topicID
}

// GetMessages returns the messages in the topic
func (t *Topic) GetMessages() []*Message {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.messages
}

// GetSubscribers returns the subscribers for this topic
func (t *Topic) GetSubscribers() []*TopicSubscriber {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	return t.subscribers
}

// AddMessage adds a message to the topic
func (t *Topic) AddMessage(message *Message) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.messages = append(t.messages, message)
}

// AddSubscriber adds a subscriber to the topic
func (t *Topic) AddSubscriber(subscriber *TopicSubscriber) {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.subscribers = append(t.subscribers, subscriber)
}
