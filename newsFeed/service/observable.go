package service

type IObservable interface {
	Subscribe(observable IObserver)
	Unsubscribe(observable IObserver)
	Fire(msgtype string, data interface{})
}
type Observable struct {
	Subscribers []IObserver
}

func (o *Observable) Subscribe(observable IObserver) {
	o.Subscribers = append(o.Subscribers, observable)
}

func (o *Observable) Unsubscribe(observable IObserver) {
	for i, subscriber := range o.Subscribers {
		if subscriber == observable {
			o.Subscribers = append(o.Subscribers[:i], o.Subscribers[i+1:]...)
			break
		}
	}
}

func (o *Observable) Fire(msgtype string, data interface{}) {
	for _, subscriber := range o.Subscribers {
		subscriber.Notify(msgtype, data)
	}
}

func NewObservable() IObservable {
	return &Observable{Subscribers: make([]IObserver, 0)}
}
