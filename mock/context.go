package mock

import "ivoeditor.com/ivo"

type context struct{}

func (c *context) Logger() ivo.Logger {
	return NewLogger()
}

func (c *context) Quit() {
}

func (c *context) Command(ivo.Command) {
}

func (c *context) Buffer() *ivo.Buffer {
	return nil
}

func (c *context) Render() {
}

func NewContext() ivo.Context {
	return &context{}
}