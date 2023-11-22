## Systemd formatter for logrus
Formats logrus entries into systemd journald style log lines.

For example:
```go
logger.Info("executing something")
logger.Error("an error happened")
```

    <6>executing something
    <3>an error happened

So info record will be printed as `<6>Executing something` and the `<6>` here is a syslog prefix for INFO level.

### Log level mapping
It maps:

    PanicLevel -> EMER(0)
    FatalLevel -> CRIT(2)
    ErrorLevel -> ERR(3)
    WarnLevel -> WARNING(4)
    InfoLevel -> INFO(6)
    DebugLevel -> DEBUG(7)
    TraceLevel -> DEBUG(7)

## Install

    go get -u github.com/quillaja/logrus-systemd-formatter/systemdfmt

## Usage

```go
package main

import (
	"github.com/quillaja/logrus-systemd-formatter/systemdfmt"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&systemdfmt.Formatter{})
	logger.SetLevel(logrus.DebugLevel)
	l := logger.WithField("TraceID", "42")
	l.Debug("Debug message")
	l.Info("Describe execution step or the app sends/received a request from external system, minor error occurred like a timeout")
	l.Warn("Something suspicious happened, used deprecated API or an error occurred because a request is invalid")
	// Multiline errors would be linearized into one string with line endings replaced with space
	err := "Multi\nline\nstacktrace"
	l.Error("Unexpected internal error occurred: invalid request format", err)
	l.Fatal("App can't do something: a port is already taken, missing config etc, fatal panic")
	l.Panic("The app doesn't work: out of memory, CPU is burning")
}
```

See [example](systemdfmt/examples/main.go)

## License
[0BSD](https://opensource.org/licenses/0BSD) (effectively Public Domain)

## See also
 * [go-logger-noalloc](https://github.com/stokito/go-logger-noalloc) more low level library but without dependencies