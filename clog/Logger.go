package clog

import (
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

func (l *Logger) Close() error {
	return nil
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

func (l *Logger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *Logger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *Logger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

func (l *Logger) Error(msg string) {
	l.log.Error().Msg(msg)
}

func (l *Logger) Fatal(msg string) {
	l.log.Error().Msg(msg)
	os.Exit(1)
}

func (l *Logger) Debugf(format string, values ...any) {
	l.log.Debug().Msgf(format, values...)
}

func (l *Logger) Infof(format string, values ...any) {
	l.log.Info().Msgf(format, values...)
}

func (l *Logger) Warnf(format string, values ...any) {
	l.log.Warn().Msgf(format, values...)
}

func (l *Logger) Errorf(format string, values ...any) {
	l.log.Error().Msgf(format, values...)
}

func (l *Logger) Fatalf(format string, values ...any) {
	l.log.Error().Msgf(format, values...)
	os.Exit(1)
}

func (l *Logger) Print(values ...any) {
	l.log.Print(values...)
}

func (l *Logger) Printf(format string, values ...any) {
	l.log.Printf(format, values...)
}

func (l *Logger) Println(values ...any) {
	l.log.Print(values...)
}

func (l *Logger) Write(data []byte) (int, error) {
	l.Print(string(data))
	return len(data), nil
}
