package gcplog

import (
	"log"
	"os"

	"github.com/infomaniac/go-logger"
)

var (
	globallogger *Logger
)

func init() {
	app, err := os.Executable()
	if err != nil {
		app = "global"
	}
	globallogger, err = New(app, logger.DEBUG)
	if err != nil {
		log.Fatal(err)
	}
}

func Debug(msg string) {
	globallogger.Debug(msg)
}
func Debugf(format string, values ...any) {
	globallogger.Debugf(format, values...)
}

func Info(msg string) {
	globallogger.Info(msg)
}
func Infof(format string, values ...any) {
	globallogger.Infof(format, values...)
}

func Warn(msg string) {
	globallogger.Warn(msg)
}
func Warnf(format string, values ...any) {
	globallogger.Warnf(format, values...)
}

func Error(msg string) {
	globallogger.Error(msg)
}
func Errorf(format string, values ...any) {
	globallogger.Errorf(format, values...)
}

func Fatal(msg string) {
	globallogger.Fatal(msg)
}
func Fatalf(format string, values ...any) {
	globallogger.Fatalf(format, values...)
}

// Compatiblity with io.Writer
func Write(data []byte) (int, error) {
	return globallogger.Write(data)
}

func SetLevel(lvl logger.Level) {
	globallogger.level = lvl
}
