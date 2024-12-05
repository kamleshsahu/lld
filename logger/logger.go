package logger

type LogLevel int

const (
	DEBUG LogLevel = 1
	INFO  LogLevel = 2
	ERROR LogLevel = 3
	FATAL LogLevel = 4
)

func (l LogLevel) String() string {
	return [...]string{"", "DEBUG", "INFO", "ERROR", "FATAL"}[l]
}

type Logger struct {
	LoggerType map[LogLevel]ILogger
}

func GetLogger() Logger {
	return Logger{}
}

func (l *Logger) Info(s string) {
	l.LoggerType[DEBUG].Log(INFO, s)
}

func (l *Logger) Debug(s string) {
	l.LoggerType[DEBUG].Log(DEBUG, s)
}

func (l *Logger) Error(s string) {
	l.LoggerType[DEBUG].Log(ERROR, s)
}

func (l *Logger) Fatal(s string) {
	l.LoggerType[DEBUG].Log(FATAL, s)
}
