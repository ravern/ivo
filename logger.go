package ivo

import (
	stdlog "log"
	"os"
)

// Logger represents a generic logger.
type Logger interface {
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
}

// logger is a core implementation of Logger.
type logger struct {
	infoLogger *stdlog.Logger
	errLogger  *stdlog.Logger
}

// newLogger creates a new logger that logs to os.Stderr.
func newLogger() *logger {
	return &logger{
		infoLogger: stdlog.New(os.Stderr, "INFO", stdlog.LstdFlags),
		errLogger:  stdlog.New(os.Stderr, "ERROR", stdlog.LstdFlags),
	}
}

func (l *logger) Info(v ...interface{}) {
	l.infoLogger.Print(v)
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.infoLogger.Printf(format, v)
}

func (l *logger) Error(v ...interface{}) {
	l.errLogger.Print(v)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.errLogger.Printf(format, v)
}
