package logType

import (
	"fmt"
	"lld/logger"
)

type errorLogger struct {
	next *logger.ILogger
}

func (i *errorLogger) Log(level logger.LogLevel, s string) {
	if level >= logger.ERROR {
		fmt.Println(logger.ERROR.String(), s)
	}

	if i.next != nil {
		(*i.next).Log(level, s)
	}
}

func (i *errorLogger) Next(logger *logger.ILogger) {
	i.next = logger
}

func NewErrorLogger() logger.ILogger {
	return &errorLogger{}
}
