package logger

type ILogger interface {
	Log(level int, s string)
	Next(level int)
}
