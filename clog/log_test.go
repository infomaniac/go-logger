package clog

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/infomaniac/go-logger"
	"github.com/stretchr/testify/assert"
)

func TestCreateLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	globallogger = New("TestCreateLogger", logger.INFO, false, buf)

	e := &struct {
		Module  string
		Level   string
		Message string
	}{}

	Info("Hi")
	err := json.Unmarshal(buf.Bytes(), e)
	assert.NoError(t, err)
	assert.Equal(t, "TestCreateLogger", e.Module)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "Hi", e.Message)

	buf.Reset()
	Infof("Hello, %s!", "Mr. Anderson")
	err = json.Unmarshal(buf.Bytes(), e)
	assert.NoError(t, err)
	assert.Equal(t, "TestCreateLogger", e.Module)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "Hello, Mr. Anderson!", e.Message)
}

func TestCreateLoggerConsole(t *testing.T) {
	buf := &bytes.Buffer{}
	globallogger = New("TestCreateLoggerConsole", logger.DEBUG, true, buf)

	e := &struct {
		Module  string
		Level   string
		Message string
	}{}

	Info("Hi")
	err := json.Unmarshal(buf.Bytes(), e)
	assert.NoError(t, err)
	assert.Equal(t, "TestCreateLoggerConsole", e.Module)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "Hi", e.Message)

	buf.Reset()
	Infof("Hello, %s!", "Mr. Anderson")
	err = json.Unmarshal(buf.Bytes(), e)
	assert.NoError(t, err)
	assert.Equal(t, "TestCreateLoggerConsole", e.Module)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "Hello, Mr. Anderson!", e.Message)
}

func TestLogLevelOutput(t *testing.T) {
	testCases := []struct {
		lvl  logger.Level
		name string
	}{
		{lvl: logger.TRACE, name: "TRACE"},
		{lvl: logger.DEBUG, name: "DEBUG"},
		{lvl: logger.INFO, name: "INFO"},
		{lvl: logger.WARN, name: "WARN"},
		{lvl: logger.ERROR, name: "ERROR"},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			globallogger = New(tc.name, tc.lvl, false, buf)

			Trace("tada")
			switch HasLvl(logger.TRACE) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			Debug("tada")
			switch HasLvl(logger.DEBUG) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			Info("tada")
			switch HasLvl(logger.INFO) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			Warn("tada")
			switch HasLvl(logger.WARN) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			Error("tada")
			switch HasLvl(logger.ERROR) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}
		})
	}
}

func TestLogLevelOutputf(t *testing.T) {
	testCases := []struct {
		lvl  logger.Level
		name string
	}{
		{lvl: logger.TRACE, name: "TRACE"},
		{lvl: logger.DEBUG, name: "DEBUG"},
		{lvl: logger.INFO, name: "INFO"},
		{lvl: logger.WARN, name: "WARN"},
		{lvl: logger.ERROR, name: "ERROR"},
	}
	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			globallogger = New(tc.name, tc.lvl, false, buf)

			buf.Reset()
			Tracef("%s", "tada")
			switch HasLvl(logger.TRACE) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			Debugf("%s", "tada")
			switch HasLvl(logger.DEBUG) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			Infof("%s", "tada")
			switch HasLvl(logger.INFO) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			Warnf("%s", "tada")
			switch HasLvl(logger.WARN) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			Errorf("%s", "tada")
			switch HasLvl(logger.ERROR) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}
		})
	}
}

func TestLogLevelOutputDefault(t *testing.T) {
	testCases := []struct {
		name        string
		method      func(msg string)
		setLevel    logger.Level
		expectEmpty bool
	}{
		{name: "TRACE", setLevel: logger.INFO, method: Trace, expectEmpty: true},
		{name: "DEBUG", setLevel: logger.INFO, method: Debug, expectEmpty: true},
		{name: "INFO ", setLevel: logger.INFO, method: Info, expectEmpty: false},
		{name: "INFO ", setLevel: logger.WARN, method: Info, expectEmpty: true},
		{name: "WARN ", setLevel: logger.WARN, method: Warn, expectEmpty: false},
		{name: "ERROR", setLevel: logger.INFO, method: Error, expectEmpty: false},
	}

	buf := &bytes.Buffer{}
	globallogger = New("TestLogLevelOutputDefault", logger.INFO, true, buf)
	msg := "tada"

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			SetLvl(tc.setLevel)
			tc.method(msg)
			switch tc.expectEmpty {
			case true:
				assert.Equal(t, buf.String(), "", tc.name)
			case false:
				assert.Contains(t, buf.String(), msg, tc.name)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	buf := &bytes.Buffer{}
	globallogger = New("TestPrint", logger.DEBUG, true, buf)

	Print("tada")
	assert.Contains(t, buf.String(), "tada", "Print")

	buf.Reset()
	Printf("%s", "tada")
	assert.Contains(t, buf.String(), "tada", "Printf")
}

func TestNoInit(t *testing.T) {
	globallogger = New("global", logger.DEBUG, true)
	msg := "tada"
	Trace(msg)
	Debug(msg)
	Info(msg)
	Warn(msg)
	Error(msg)
	Print(msg)

	Tracef("%s", msg)
	Debugf("%s", msg)
	Infof("%s", msg)
	Warnf("%s", msg)
	Errorf("%s", msg)
	Printf("%s", msg)
}

func TestGlobalWrite(t *testing.T) {
	buf := &bytes.Buffer{}
	globallogger = New("global", logger.DEBUG, false, buf)

	type evt struct {
		Level   string
		Message string
	}
	e := &evt{}

	buf.Reset()
	SetLvl(logger.TRACE)
	n, err := Write([]byte("HelloT"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "trace", e.Level)
	assert.Equal(t, "HelloT", e.Message)

	buf.Reset()
	SetLvl(logger.DEBUG)
	n, err = Write([]byte("HelloD"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "debug", e.Level)
	assert.Equal(t, "HelloD", e.Message)

	buf.Reset()
	SetLvl(logger.INFO)
	n, err = Write([]byte("HelloI"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "HelloI", e.Message)

	buf.Reset()
	SetLvl(logger.WARN)
	n, err = Write([]byte("HelloW"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "warn", e.Level)
	assert.Equal(t, "HelloW", e.Message)

	buf.Reset()
	SetLvl(logger.ERROR)
	n, err = Write([]byte("HelloE"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "error", e.Level)
	assert.Equal(t, "HelloE", e.Message)

	buf.Reset()
	SetLvl(12)
	n, err = Write([]byte("HelloE"))
	assert.Equal(t, n, 0)
	assert.Error(t, err)
	assert.Empty(t, buf.Bytes())
}
