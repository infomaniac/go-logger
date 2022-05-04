package logger

type ILogger interface {
	Debug(msg string)
	Debugf(format string, values ...any)

	Info(msg string)
	Infof(format string, values ...any)

	Warn(msg string)
	Warnf(format string, values ...any)

	Error(msg string)
	Errorf(format string, values ...any)

	Fatal(msg string)
	Fatalf(format string, values ...any)

	// Compatiblity with io.Writer
	Write(data []byte) (int, error)

	SetLevel(Level)
}
