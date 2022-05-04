package logger

type Level int8

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
	TRACE Level = -1
)
