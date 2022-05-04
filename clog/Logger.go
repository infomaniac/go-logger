package clog

import (
	"errors"
	"io"
	"os"
	"reflect"
	"time"

	"github.com/infomaniac/go-logger"
	"github.com/rs/zerolog"
)

// Logger represents a logging object.
type Logger struct {
	log   zerolog.Logger
	level logger.Level
}

// New initializes a new logger
func New(module string, lvl logger.Level, stdoutConsole bool, uncheckedOut ...io.Writer) *Logger {
	var out []io.Writer
	host, _ := os.Hostname()

	for _, w := range uncheckedOut {
		if w == nil || (reflect.ValueOf(w).Kind() == reflect.Ptr && reflect.ValueOf(w).IsNil()) {
			// invalid writer, nil or the underlying type is nil
			continue
		}
		out = append(out, w)
	}

	if stdoutConsole {
		console := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339Nano}
		out = append(out, console)
	}
	zerolog.TimeFieldFormat = time.RFC3339Nano

	multi := zerolog.MultiLevelWriter(out...)
	logger := zerolog.New(multi).Level(zerolog.Level(lvl)).With().Timestamp().Str("module", module).Str("host", host).Logger()

	l := &Logger{log: logger}
	l.SetLvl(lvl)
	return l
}

// HasLvl checks if a log statement will be printed when sent with a specific level.
// This can and should be used to see if a log statment should be prepared at all.
func (l *Logger) HasLvl(lvl logger.Level) bool {
	return l.level <= lvl
}

// SetLvl will set the global logger to a different Loglevel.
func (l *Logger) SetLvl(lvl logger.Level) {
	l.level = lvl
	l.log = l.log.Level(zerolog.Level(lvl))
}

// Trace logs the message `msg` with the corresponding loglevel error.
func (l *Logger) Trace(msg string) error {
	l.log.Trace().Msg(msg)
	return nil
}

// Debug logs the message `msg` with the corresponding loglevel error.
func (l *Logger) Debug(msg string) error {
	l.log.Debug().Msg(msg)
	return nil
}

// Info logs the message `msg` with the corresponding loglevel error.
func (l *Logger) Info(msg string) error {
	l.log.Info().Msg(msg)
	return nil
}

// Warn logs the message `msg` with the corresponding loglevel error.
func (l *Logger) Warn(msg string) error {
	l.log.Warn().Msg(msg)
	return nil
}

// Error logs the message `msg` with the corresponding loglevel error.
func (l *Logger) Error(msg string) error {
	l.log.Error().Msg(msg)
	return nil
}

// Fatal logs the message `msg` with the corresponding loglevel error.
func (l *Logger) Fatal(msg string) {
	l.log.Error().Msg(msg)
	os.Exit(1)
}

// Tracef logs the message `msg` with the loglevel error.
func (l *Logger) Tracef(format string, values ...interface{}) error {
	l.log.Trace().Msgf(format, values...)
	return nil
}

// Debugf logs the message `msg` with the loglevel error.
func (l *Logger) Debugf(format string, values ...interface{}) error {
	l.log.Debug().Msgf(format, values...)
	return nil
}

// Infof logs the message `msg` with the loglevel error.
func (l *Logger) Infof(format string, values ...interface{}) error {
	l.log.Info().Msgf(format, values...)
	return nil
}

// Warnf logs the message `msg` with the loglevel error.
func (l *Logger) Warnf(format string, values ...interface{}) error {
	l.log.Warn().Msgf(format, values...)
	return nil
}

// Errorf logs the message `msg` with the loglevel error.
func (l *Logger) Errorf(format string, values ...interface{}) error {
	l.log.Error().Msgf(format, values...)
	return nil
}

// Fatalf logs the message `msg` with the loglevel error.
func (l *Logger) Fatalf(format string, values ...interface{}) {
	l.log.Error().Msgf(format, values...)
	os.Exit(1)
}

// Print logs the message `msg` with the DEBUG loglevel.
func (l *Logger) Print(values ...interface{}) {
	l.log.Print(values...)
}

// Printf logs with the DEBUG loglevel.
func (l *Logger) Printf(format string, values ...interface{}) {
	l.log.Printf(format, values...)
}

// Println logs with the DEBUG loglevel.
func (l *Logger) Println(values ...interface{}) {
	l.log.Print(values...)
}

func (l *Logger) Write(data []byte) (int, error) {
	switch l.level {
	case logger.TRACE:
		return len(data), l.Trace(string(data))
	case logger.DEBUG:
		return len(data), l.Debug(string(data))
	case logger.INFO:
		return len(data), l.Info(string(data))
	case logger.WARN:
		return len(data), l.Warn(string(data))
	case logger.ERROR:
		return len(data), l.Error(string(data))
	}
	return 0, errors.New("invalid log level")
}
