package service

type IObserver interface {
	Notify(msgtype string, data interface{})
}
