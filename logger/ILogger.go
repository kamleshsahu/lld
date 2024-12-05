package logger

type ILogger interface {
	Log(level LogLevel, s string)
	Next(iLogger *ILogger)
}
