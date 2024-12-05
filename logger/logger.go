package logger

const (
	INFO  = 1
	DEBUG = 2
	ERROR = 3
	FATAL = 4
)

type Logger interface {
	Info(message string)
	Debug(message string)
	Error(message string)
	Fatal(message string)
}

type logger struct {
}

func (l *logger) Info(message string) {

}

func (l *logger) Debug(message string) {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Error(message string) {
	//TODO implement me
	panic("implement me")
}

func (l *logger) Fatal(message string) {
	//TODO implement me
	panic("implement me")
}

func New() Logger {
	return &logger{}
}
