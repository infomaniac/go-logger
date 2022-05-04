package gcplog

import (
	"context"
	"errors"
	"fmt"
	"os"

	"cloud.google.com/go/logging"
	"github.com/infomaniac/go-logger"
)

// Logger represents a logging object.
type Logger struct {
	client *logging.Client
	l      *logging.Logger
	level  logger.Level
}

// New initializes a new logger
func New(module string, lvl logger.Level) (*Logger, error) {
	if module == "" {
		return nil, errors.New("module name is required")
	}

	project, err := CurrentProject()
	if err != nil {
		return nil, err
	}
	if project == "" {
		return nil, errors.New("project not found")
	}

	client, err := logging.NewClient(context.Background(), project)
	if err != nil {
		return nil, err
	}

	l := &Logger{
		client: client,
		l:      client.Logger(module),
		level:  lvl,
	}
	return l, nil
}

// HasLvl checks if a log statement will be printed when sent with a specific level.
// This can and should be used to see if a log statment should be prepared at all.
func (l *Logger) HasLvl(lvl logger.Level) bool {
	return l.level <= lvl
}

// SetLvl will set the global logger to a different Loglevel.
func (l *Logger) SetLvl(lvl logger.Level) {
	l.level = lvl
}

func (l *Logger) Debug(msg string) {
	if !l.HasLvl(logger.DEBUG) {
		return
	}

	e := logging.Entry{
		Severity: logging.Debug,
		Payload:  msg,
	}
	l.l.Log(e)
}

func (l *Logger) Info(msg string) {
	if !l.HasLvl(logger.INFO) {
		return
	}

	l.l.Log(logging.Entry{
		Severity: logging.Info,
		Payload:  msg,
	})
}

func (l *Logger) Warn(msg string) {
	if !l.HasLvl(logger.WARN) {
		return
	}

	l.l.Log(logging.Entry{
		Severity: logging.Warning,
		Payload:  msg,
	})
}

func (l *Logger) Error(msg string) {
	l.l.Log(logging.Entry{
		Severity: logging.Error,
		Payload:  msg,
	})
}

func (l *Logger) Fatal(msg string) {
	l.l.Log(logging.Entry{
		Severity: logging.Emergency,
		Payload:  msg,
	})
	os.Exit(1)
}

// Debugf logs the message `msg` with the loglevel error.
func (l *Logger) Debugf(format string, values ...any) {
	if !l.HasLvl(logger.DEBUG) {
		return
	}

	l.l.Log(logging.Entry{
		Severity: logging.Debug,
		Payload:  fmt.Sprintf(format, values...),
	})
}

func (l *Logger) Infof(format string, values ...any) {
	if !l.HasLvl(logger.INFO) {
		return
	}

	l.l.Log(logging.Entry{
		Severity: logging.Info,
		Payload:  fmt.Sprintf(format, values...),
	})
}

func (l *Logger) Warnf(format string, values ...any) {
	if !l.HasLvl(logger.WARN) {
		return
	}

	l.l.Log(logging.Entry{
		Severity: logging.Warning,
		Payload:  fmt.Sprintf(format, values...),
	})
}

func (l *Logger) Errorf(format string, values ...any) {
	l.l.Log(logging.Entry{
		Severity: logging.Error,
		Payload:  fmt.Sprintf(format, values...),
	})
}

func (l *Logger) Fatalf(format string, values ...any) {
	l.l.Log(logging.Entry{
		Severity: logging.Emergency,
		Payload:  fmt.Sprintf(format, values...),
	})
	os.Exit(1)
}

func (l *Logger) Write(data []byte) (int, error) {
	l.l.Log(logging.Entry{
		Severity: logging.Default,
		Payload:  string(data),
	})
	return len(data), nil
}
