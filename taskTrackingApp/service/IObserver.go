package service

type IObserver interface {
	Notify(data interface{}) error
}
