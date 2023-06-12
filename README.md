# Logging for Console and GCP

## Common Interface `logger.ILogger`

The `ILogger` interface is a common interface for all loggers in this package.
It provides the ability, to use e.g. `clog` for unit test and `gcplog` for production.


## clog

`clog` or "console-log" is a convenience wrapper around `github.com/rs/zerolog` that makes it easy and intuitive to use.
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
### Usage
### Dedicated logger
The dedicated logger is the default way to use clog. A logger can be initialized with a "Module" parameter, that helps to separate different loggers from different modules, meaning each module can use it's own logger and output will be prefixed accodingly.


```go
out1 := &bytes.Buffer{} // or any other IO.writer
out2 := &bytes.Buffer{} // or any other IO.writer

ll := clog.New("MyTestApplication", clog.DEBUG, true, out1, out2)
defer ll.Close()

ll.Info("This is an Info Statement")
```


## gcplog

### Usage

```go
ll, err := gcplog.New("MyTestApplication", logger.DEBUG)
if err != nil {
	panic(err)
}
defer ll.Close()

ll.Info("This is an Info Statement")

```