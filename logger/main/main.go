package main

import (
	"lld/logger"
	"lld/logger/logType"
)

func main() {
	mainLogger := logger.GetLogger()

	infoLogger := logType.NewInfoLogger()
	debugLogger := logType.NewDebugLogger()
	errorLogger := logType.NewErrorLogger()
	debugLogger.Next(&infoLogger)
	infoLogger.Next(&errorLogger)

	mainLogger.LoggerType = map[logger.LogLevel]logger.ILogger{
		logger.INFO:  infoLogger,
		logger.DEBUG: debugLogger,
		logger.ERROR: errorLogger,
	}

	mainLogger.Info("printing info log")
	mainLogger.Debug("printing info debug")
	mainLogger.Error("printing info error")
}
