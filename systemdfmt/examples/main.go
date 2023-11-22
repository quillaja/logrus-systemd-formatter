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
