// Package systemdfmt implements a formatter for logrus
// that produces log lines for use in systemd services.
//
// See https://github.com/sirupsen/logrus
package systemdfmt

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// SystemdFormatter prepares logrus entries for output to systemd.
// It maps:
//   PanicLevel -> EMER(0)
//   FatalLevel -> CRIT(2)
//   ErrorLevel -> ERR(3)
//   WarnLevel -> WARNING(4)
//   InfoLevel -> INFO(6)
//   DebugLevel -> DEBUG(7)
//   TraceLevel -> DEBUG(7)
type SystemdFormatter struct{}

// Format formats entry to a systemd log line.
func (f *SystemdFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	var level uint32 = 6 // default systemd lvl 6 "info"
	switch entry.Level {
	case logrus.PanicLevel: // 0
		level = emerg
	case logrus.FatalLevel: // 1
		level = crit
	case logrus.ErrorLevel: // 2
		level = err
	case logrus.WarnLevel: // 3
		level = warning
	case logrus.InfoLevel: // 4
		level = info
	case logrus.DebugLevel: // 5
		level = debug
	case logrus.TraceLevel: // 6
		level = debug
	}

	msg := fmt.Sprintf("%s %+v", entry.Message, entry.Data)
	msg = strings.ReplaceAll(msg, "\n", " ") // systemd log lines contain no newlines
	return []byte(fmt.Sprintf("<%d>%s\n", level, msg)), nil
}

// systemd log levels
const (
	emerg = iota
	alert
	crit
	err
	warning
	notice
	info
	debug
)
