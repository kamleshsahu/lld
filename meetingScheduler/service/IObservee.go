package service

type IObservee interface {
	Fire(event string)
}
