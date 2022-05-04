package clog

import (
	"bytes"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/infomaniac/go-logger"
	"github.com/stretchr/testify/assert"
)

func TestGlobalCreateLogger(t *testing.T) {
	buf := &bytes.Buffer{}

	l := New("TestCreateLogger", logger.DEBUG, false, buf)
	e := &struct {
		Module  string
		Host    string
		Level   string
		Message string
	}{}

	l.Info("Hi")
	err := json.Unmarshal(buf.Bytes(), e)
	assert.NoError(t, err)
	assert.Equal(t, "TestCreateLogger", e.Module)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "Hi", e.Message)

	buf.Reset()
	l.Infof("Hello, %s!", "Mr. Anderson")
	err = json.Unmarshal(buf.Bytes(), e)
	assert.NoError(t, err)
	assert.Equal(t, "TestCreateLogger", e.Module)
	host, _ := os.Hostname()
	assert.Equal(t, host, e.Host)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "Hello, Mr. Anderson!", e.Message)
}

func TestGlobalCreateLoggerConsole(t *testing.T) {
	buf := &bytes.Buffer{}

	l := New("TestCreateLogger", logger.DEBUG, true, buf)
	e := &struct {
		Module  string
		Level   string
		Message string
	}{}

	l.Info("Hi")
	err := json.Unmarshal(buf.Bytes(), e)
	assert.NoError(t, err)
	assert.Equal(t, "TestCreateLogger", e.Module)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "Hi", e.Message)

	buf.Reset()
	l.Infof("Hello, %s!", "Mr. Anderson")
	err = json.Unmarshal(buf.Bytes(), e)
	assert.NoError(t, err)
	assert.Equal(t, "TestCreateLogger", e.Module)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "Hello, Mr. Anderson!", e.Message)
}

func TestCompareDates(t *testing.T) {
	var err error
	const interval time.Duration = 100 * time.Millisecond

	buf1 := &bytes.Buffer{}
	l1 := New("1", logger.INFO, false, buf1)

	buf2 := &bytes.Buffer{}
	l2 := New("2", logger.INFO, false, buf2)

	l1.Info("Hello1")
	time.Sleep(interval)
	l2.Info("Hello2")
	time.Sleep(interval)

	type event struct {
		Timestamp time.Time `json:"time"`
		Message   string    `json:"message"`
	}

	e1 := &event{}
	e2 := &event{}

	err = json.Unmarshal(buf1.Bytes(), e1)
	assert.NoError(t, err, buf1.String())
	err = json.Unmarshal(buf2.Bytes(), e2)
	assert.NoError(t, err, buf2.String())

	assert.Equal(t, "Hello1", e1.Message)
	assert.Equal(t, "Hello2", e2.Message)

	assert.NotEqual(t, e1.Timestamp, e2.Timestamp)

	dur := e2.Timestamp.Sub(e1.Timestamp)

	if dur < interval {
		t.Errorf("Timestamps are less than the sleep interval, %v and %v", e1.Timestamp, e2.Timestamp)
	}
	if dur-interval > 10*time.Millisecond {
		t.Errorf("Timestamps are too far apart %v and %v", e1.Timestamp, e2.Timestamp)
	}
}

func TestGlobalLogLevelOutput(t *testing.T) {
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
			l := New(tc.name, tc.lvl, false, buf)

			l.Trace("tada")
			switch l.HasLvl(logger.TRACE) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			l.Debug("tada")
			switch l.HasLvl(logger.DEBUG) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			l.Info("tada")
			switch l.HasLvl(logger.INFO) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			l.Warn("tada")
			switch l.HasLvl(logger.WARN) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			l.Error("tada")
			switch l.HasLvl(logger.ERROR) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}
		})
	}
}

func TestGlobalLogLevelOutputf(t *testing.T) {
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
			l := New(tc.name, tc.lvl, false, buf)

			buf.Reset()
			l.Tracef("%s", "tada")
			switch l.HasLvl(logger.TRACE) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			l.Debugf("%s", "tada")
			switch l.HasLvl(logger.DEBUG) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			l.Infof("%s", "tada")
			switch l.HasLvl(logger.INFO) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			l.Warnf("%s", "tada")
			switch l.HasLvl(logger.WARN) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}

			buf.Reset()
			l.Errorf("%s", "tada")
			switch l.HasLvl(logger.ERROR) {
			case true:
				assert.NotEmpty(t, buf, tc.name, tc.lvl)
			case false:
				assert.Empty(t, buf, tc.name, tc.lvl)
			}
		})
	}
}

func TestLogLevelOutputDefault2(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New("TestLogLevelOutputDefault", logger.INFO, true, buf)

	testCases := []struct {
		name        string
		method      func(msg string) error
		setLevel    logger.Level
		expectEmpty bool
	}{
		{name: "TRACE1", setLevel: logger.INFO, method: l.Trace, expectEmpty: true},
		{name: "TRACE2", setLevel: logger.TRACE, method: l.Trace, expectEmpty: false},
		{name: "DEBUG1", setLevel: logger.TRACE, method: l.Debug, expectEmpty: false},
		{name: "DEBUG2", setLevel: logger.INFO, method: l.Debug, expectEmpty: true},
		{name: "INFO1", setLevel: logger.INFO, method: l.Info, expectEmpty: false},
		{name: "INFO2", setLevel: logger.WARN, method: l.Info, expectEmpty: true},
		{name: "WARN1", setLevel: logger.ERROR, method: l.Warn, expectEmpty: true},
		{name: "WARN2", setLevel: logger.WARN, method: l.Warn, expectEmpty: false},
		{name: "ERROR1", setLevel: logger.INFO, method: l.Error, expectEmpty: false},
		{name: "ERROR2", setLevel: logger.ERROR, method: l.Error, expectEmpty: false},
	}

	msg := "tada"
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			buf.Reset()
			l.SetLvl(tc.setLevel)
			tc.method(msg)
			switch tc.expectEmpty {
			case true:
				assert.Equal(t, "", buf.String(), tc.name, l.level)
			case false:
				assert.Contains(t, buf.String(), msg, tc.name, l.level)
			}
		})
	}
}

func TestPrintFunctions(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New("TestLogLevelOutputDefault", logger.DEBUG, true, buf)

	l.Print("Hello", "World")
	assert.Contains(t, buf.String(), "HelloWorld")

	buf.Reset()
	l.Printf("%s, %s", "Hello", "World")
	assert.Contains(t, buf.String(), "Hello, World")

	buf.Reset()
	l.Println("HelloWorld")
	assert.Contains(t, buf.String(), "HelloWorld")

}

func TestNewWithNil(t *testing.T) {
	buf := &bytes.Buffer{}

	var f *os.File = nil
	l := New("Tada", logger.INFO, true, buf, nil, f)
	l.Info("Hello")
	assert.Contains(t, buf.String(), "Hello")
}

func TestWrite(t *testing.T) {
	buf := &bytes.Buffer{}
	l := New("TestWrite", logger.TRACE, false, buf)

	type evt struct {
		Level   string
		Message string
	}
	e := &evt{}

	buf.Reset()
	l.SetLvl(logger.TRACE)
	n, err := l.Write([]byte("HelloT"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "trace", e.Level)
	assert.Equal(t, "HelloT", e.Message)

	buf.Reset()
	l.SetLvl(logger.DEBUG)
	n, err = l.Write([]byte("HelloD"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "debug", e.Level)
	assert.Equal(t, "HelloD", e.Message)

	buf.Reset()
	l.SetLvl(logger.INFO)
	n, err = l.Write([]byte("HelloI"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "info", e.Level)
	assert.Equal(t, "HelloI", e.Message)

	buf.Reset()
	l.SetLvl(logger.WARN)
	n, err = l.Write([]byte("HelloW"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "warn", e.Level)
	assert.Equal(t, "HelloW", e.Message)

	buf.Reset()
	l.SetLvl(logger.ERROR)
	n, err = l.Write([]byte("HelloE"))
	assert.Equal(t, n, 6)
	assert.NoError(t, err)
	json.Unmarshal(buf.Bytes(), e)
	assert.Equal(t, "error", e.Level)
	assert.Equal(t, "HelloE", e.Message)

	buf.Reset()
	l.SetLvl(12)
	n, err = l.Write([]byte("HelloE"))
	assert.Equal(t, n, 0)
	assert.Error(t, err)
	assert.Empty(t, buf.Bytes())
}
