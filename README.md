# clog – coop logging for golang
`clog` or "coop-log" is a convenience wrapper around `github.com/rs/zerolog` that makes it easy and intuitive to use.
The default `stdout` logging produces nice and colourful log statements on linux consoles.
The logger accepts multiple `io.Writer`s that will output log statements in JSON form.
The fields are
- **level** the log-level (`TRACE`, `DEBUG`, `INFO`, `WARN`, `ERROR`)
- **module** the name of the "module" that was given when the logger was initialized
- **host** hostname of the machine, running the application
- **time** timestamp
- **message** the log message, whatever that may be.

```json
{
	"level":   "info",
	"module":  "MyModule",
	"host":    "svrl1midtoolt51",
	"time":    "2022-03-28T09:35:14.769433+02:00",
	"message": "This is a INFO logging statement."
}
```
## Usage
### Dedicated logger
The dedicated logger is the default way to use clog. A logger can be initialized with a "Module" parameter, that helps to separate different loggers from different modules.
```go
out1 := &bytes.Buffer{} // or any other IO.writer
out2 := &bytes.Buffer{} // or any other IO.writer
import "gitlab.hs.coop.ch/middleware/golibs/clog"

ll := clog.New("MyModule", clog.DEBUG, true, out1, out2)
```

### Global Logger (stdout)
The "global" logger can be used on a package level without special initialisation. The output is only on stdout and in plaintext, the log level can be set.
```go
import "gitlab.hs.coop.ch/middleware/golibs/clog"

clog.Debugf("log this debug statement to stdout at %v", time.Now())
clog.SetLvl(clog.INFO)
```
