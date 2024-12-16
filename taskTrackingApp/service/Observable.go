package service

import "container/list"

type Observable struct {
	subscribers *list.List
}

func (o *Observable) Subscribe(subscriber IObserver) {
	o.subscribers.PushBack(subscriber)
}

func (o *Observable) Unsubscribe(subscriber IObserver) {
	for e := o.subscribers.Front(); e != nil; e = e.Next() {
		if e.Value.(IObserver) == subscriber {
			o.subscribers.Remove(e)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for e := o.subscribers.Front(); e != nil; e = e.Next() {
		err := e.Value.(IObserver).Notify(data)
		if err != nil {
			return
		}
	}
}
