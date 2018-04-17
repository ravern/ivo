package mock

import "ivoeditor.com/ivo"

type logger struct{}

func (l *logger) Info(v ...interface{}) {
}

func (l *logger) Infof(format string, v ...interface{}) {
}

func (l *logger) Error(v ...interface{}) {
}

func (l *logger) Errorf(format string, v ...interface{}) {
}

func NewLogger() ivo.Logger {
	return &logger{}
}
