package service

type INotificationService interface {
	Notify(message string, emailId string) error
}
