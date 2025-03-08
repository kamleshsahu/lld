package main

import (
	"lld/go-pubsub-queue/pubsub/model"
	"os"
	"os/signal"
	"syscall"
	"time"

	"lld/go-pubsub-queue/pubsub/queue"
)

func main() {
	queue := queue.NewQueue()

	topic1 := queue.CreateTopic("t1")
	//topic2 := queue.CreateTopic("t2")

	sub1 := NewSleepingSubscriber("sub1", 10)
	//sub2 := NewSleepingSubscriber("sub2", 10000)
	queue.Subscribe(sub1, topic1)
	//queue.Subscribe(sub2, topic1)

	//sub3 := NewSleepingSubscriber("sub3", 5000)
	//queue.Subscribe(sub3, topic1)

	queue.Publish(topic1, model.NewMessage("m1"))
	queue.Publish(topic1, model.NewMessage("m2"))

	//queue.Publish(topic2, model.NewMessage("m3"))

	time.Sleep(30 * time.Second)
	//queue.Publish(topic2, model.NewMessage("m4"))
	queue.Publish(topic1, model.NewMessage("m5"))

	//queue.ResetOffset(topic1, sub1, 0)

	// Keep main thread alive to allow worker goroutines to process
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}
