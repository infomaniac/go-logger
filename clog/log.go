package clog

import (
	"errors"
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
	globallogger = New(app, logger.DEBUG, true)
}

// SetLvl will set the global logger to a different Loglevel.
func SetLvl(lvl logger.Level) {
	globallogger.SetLvl(lvl)
}

// HasLvl checks if a log statement will be printed when sent with a specific level.
// This can and should be used to see if a log statment should be prepared at all.
func HasLvl(lvl logger.Level) bool {
	return globallogger.HasLvl(lvl)
}

// Trace will send a log message with the corresponding loglevel
func Trace(msg string) {
	if !globallogger.HasLvl(logger.TRACE) {
		return
	}
	globallogger.Trace(msg)
}

// Tracef will send a log message with the corresponding loglevel
func Tracef(format string, values ...interface{}) {
	if !globallogger.HasLvl(logger.TRACE) {
		return
	}
	globallogger.Tracef(format, values...)
}

// Debug will send a log message with the corresponding loglevel
func Debug(msg string) {
	if !globallogger.HasLvl(logger.DEBUG) {
		return
	}
	globallogger.Debug(msg)
}

// Debugf will send a log message with the corresponding loglevel
func Debugf(format string, values ...interface{}) {
	if !globallogger.HasLvl(logger.DEBUG) {
		return
	}
	globallogger.Debugf(format, values...)
}

// Info will send a log message with the corresponding loglevel
func Info(msg string) {
	if !globallogger.HasLvl(logger.INFO) {
		return
	}
	globallogger.Info(msg)
}

// Infof will send a log message with the corresponding loglevel
func Infof(format string, values ...interface{}) {
	if !globallogger.HasLvl(logger.INFO) {
		return
	}
	globallogger.Infof(format, values...)
}

// Warn will send a log message with the corresponding loglevel
func Warn(msg string) {
	if !globallogger.HasLvl(logger.WARN) {
		return
	}
	globallogger.Warn(msg)
}

// Warnf will send a log message with the corresponding loglevel
func Warnf(format string, values ...interface{}) {
	if !globallogger.HasLvl(logger.WARN) {
		return
	}
	globallogger.Warnf(format, values...)
}

// Error will send a log message with the corresponding loglevel
func Error(msg string) {
	globallogger.Error(msg)
}

// Errorf will send a log message with the corresponding loglevel
func Errorf(format string, values ...interface{}) {
	globallogger.Errorf(format, values...)
}

// Fatal will send a log message with the corresponding loglevel
func Fatal(msg string) {
	globallogger.Error(msg)
	os.Exit(1)
}

// Fatalf will send a log message with the corresponding loglevel
func Fatalf(format string, values ...interface{}) {
	globallogger.Errorf(format, values...)
	os.Exit(1)
}

// Print will send a log message with the default loglevel
func Print(msg ...interface{}) {
	globallogger.Print(msg...)
}

// Printf will send a log message with the default loglevel
func Printf(format string, values ...interface{}) {
	globallogger.Printf(format, values...)
}

func Write(data []byte) (int, error) {
	switch globallogger.level {
	case logger.TRACE:
		return len(data), globallogger.Trace(string(data))
	case logger.DEBUG:
		return len(data), globallogger.Debug(string(data))
	case logger.INFO:
		return len(data), globallogger.Info(string(data))
	case logger.WARN:
		return len(data), globallogger.Warn(string(data))
	case logger.ERROR:
		return len(data), globallogger.Error(string(data))
	}
	return 0, errors.New("invalid log level")
}
