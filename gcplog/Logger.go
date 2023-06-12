package gcplog

import (
	"context"
	"errors"
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

func (l *Logger) Close() error {
	return l.client.Close()
}

// hasLvl checks if a log statement will be printed when sent with a specific level.
// This can and should be used to see if a log statment should be prepared at all.
func (l *Logger) hasLvl(lvl logger.Level) bool {
	return l.level <= lvl
}

// SetLvl will set the global logger to a different Loglevel.
func (l *Logger) SetLvl(lvl logger.Level) {
	l.level = lvl
}

func (l *Logger) Debug(msg string) {
	if !l.hasLvl(logger.DEBUG) {
		return
	}
	l.l.StandardLogger(logging.Debug).Print(msg)
}

func (l *Logger) Info(msg string) {
	if !l.hasLvl(logger.INFO) {
		return
	}
	l.l.StandardLogger(logging.Info).Print(msg)
}

func (l *Logger) Warn(msg string) {
	if !l.hasLvl(logger.WARN) {
		return
	}
	l.l.StandardLogger(logging.Warning).Print(msg)
}

func (l *Logger) Error(msg string) {
	l.l.StandardLogger(logging.Error).Print(msg)
}

func (l *Logger) Fatal(msg string) {
	l.l.StandardLogger(logging.Emergency).Print(msg)
	os.Exit(1)
}

func (l *Logger) Debugf(format string, values ...any) {
	if !l.hasLvl(logger.DEBUG) {
		return
	}
	l.l.StandardLogger(logging.Debug).Printf(format, values...)
}

func (l *Logger) Infof(format string, values ...any) {
	if !l.hasLvl(logger.INFO) {
		return
	}
	l.l.StandardLogger(logging.Info).Printf(format, values...)
}

func (l *Logger) Warnf(format string, values ...any) {
	if !l.hasLvl(logger.WARN) {
		return
	}
	l.l.StandardLogger(logging.Warning).Printf(format, values...)
}

func (l *Logger) Errorf(format string, values ...any) {
	l.l.StandardLogger(logging.Error).Printf(format, values...)
}

func (l *Logger) Fatalf(format string, values ...any) {
	l.l.StandardLogger(logging.Alert).Printf(format, values...)
	os.Exit(1)
}

func (l *Logger) Print(msg ...any) {
	l.l.StandardLogger(logging.Default).Print(msg...)
}

func (l *Logger) Printf(format string, values ...any) {
	l.l.StandardLogger(logging.Default).Printf(format, values...)
}

func (l *Logger) Write(data []byte) (int, error) {
	l.l.StandardLogger(logging.Default).Print(string(data))
	return len(data), nil
}
