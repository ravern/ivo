package mock

import "ivoeditor.com/ivo"

// logger is an empty, no-op logger.
type logger struct{}

func (l *logger) Info(v ...interface{}) {
}

func (l *logger) Infof(format string, v ...interface{}) {
}

func (l *logger) Error(v ...interface{}) {
}

func (l *logger) Errorf(format string, v ...interface{}) {
}

// NewLogger creates a new mock logger.
func NewLogger() ivo.Logger {
	return &logger{}
}
