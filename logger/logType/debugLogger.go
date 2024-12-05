package logType

import (
	"fmt"
	"lld/logger"
)

type debugLogger struct {
	next *logger.ILogger
}

func (i *debugLogger) Log(level logger.LogLevel, s string) {
	if level >= logger.DEBUG {
		fmt.Println(logger.DEBUG.String(), s)
	}
	if i.next != nil {
		(*i.next).Log(level, s)
	}
}

func (i *debugLogger) Next(logger *logger.ILogger) {
	i.next = logger
}

func NewDebugLogger() logger.ILogger {
	return &debugLogger{}
}
