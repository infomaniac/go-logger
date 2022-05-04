package logger

type ILogger interface {
	Trace(msg string) error
	Tracef(format string, values ...interface{}) error
	Debug(msg string) error
	Debugf(format string, values ...interface{}) error
	Info(msg string) error
	Infof(format string, values ...interface{}) error
	Warn(msg string) error
	Warnf(format string, values ...interface{}) error
	Error(msg string) error
	Errorf(format string, values ...interface{}) error
	Fatal(msg string)
	Fatalf(format string, values ...interface{})

	// Compatibilty with log.Logger
	Print(values ...interface{})
	Printf(format string, values ...interface{})
	Println(values ...interface{})

	// Compatiblity with io.Writer
	Write(data []byte) (int, error)

	SetLevel(Level)
}
