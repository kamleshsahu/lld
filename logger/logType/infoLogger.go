package logType

import (
	"fmt"
	"lld/logger"
)

type infoLogger struct {
	next *logger.ILogger
}

func (i *infoLogger) Log(level logger.LogLevel, s string) {
	if level >= logger.INFO {
		fmt.Println(logger.INFO.String(), s)
	}
	if i.next != nil {
		(*i.next).Log(level, s)
	}
}

func (i *infoLogger) Next(logger *logger.ILogger) {
	i.next = logger
}

func NewInfoLogger() logger.ILogger {
	return &infoLogger{}
}
