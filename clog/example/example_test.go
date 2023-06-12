package example

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/infomaniac/go-logger"
	"github.com/infomaniac/go-logger/clog"
)

func TestDedicatedLogger(t *testing.T) {
	out1 := &bytes.Buffer{}
	out2 := &bytes.Buffer{}
	ll := clog.New("MyModule", logger.DEBUG, true, out1, out2)
	defer ll.Close()

	ll.Infof("This is a INFO logging statement %v", time.Now())

	ll.Warn("This is a warning")

	ll.SetLvl(logger.INFO)
	ll.Debug("Now, this will also be ignored.")

	// out1 and out2 contains the log statements in JSON format, separated by newline
	fmt.Print(out1.String())
	fmt.Print(out2.String())
}
