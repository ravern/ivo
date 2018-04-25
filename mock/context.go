package mock

import "ivoeditor.com/ivo"

// context is an empty, no-op context.
type context struct{}

// NewContext creates a new mock context.
func NewContext() ivo.Context {
	return &context{}
}

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
