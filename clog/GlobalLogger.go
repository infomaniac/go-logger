package clog

import (
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

// Debug will send a log message with the corresponding loglevel
func Debug(msg string) {
	if !globallogger.HasLvl(logger.DEBUG) {
		return
	}
	globallogger.Debug(msg)
}

// Debugf will send a log message with the corresponding loglevel
func Debugf(format string, values ...any) {
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
func Infof(format string, values ...any) {
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
func Warnf(format string, values ...any) {
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
func Errorf(format string, values ...any) {
	globallogger.Errorf(format, values...)
}

// Fatal will send a log message with the corresponding loglevel
func Fatal(msg string) {
	globallogger.Error(msg)
	os.Exit(1)
}

// Fatalf will send a log message with the corresponding loglevel
func Fatalf(format string, values ...any) {
	globallogger.Errorf(format, values...)
	os.Exit(1)
}

// Print will send a log message with the default loglevel
func Print(msg ...any) {
	globallogger.Print(msg...)
}

// Printf will send a log message with the default loglevel
func Printf(format string, values ...any) {
	globallogger.Printf(format, values...)
}

// Printf will send a log message with the default loglevel
func Write(data []byte) (int, error) {
	globallogger.Print(string(data))
	return len(data), nil
}
