package ivo

import termbox "github.com/nsf/termbox-go"

// Context contains important objects and methods for the window
// to use.
//
// When a Context is passed to a Window via its methods, the
// previous Context expires. Once a Context is expired, Buffer
// and Render will no longer do anything.
//
// For Windows acting as a proxy to other Windows, Context can
// and should be implemented for the purpose of that Window,
// passing along the appropriate information.
type Context interface {
	// Logger should be used to perform all logging.
	//
	// In the case of core Context, it is the logger assigned to
	// the Core object.
	Logger() Logger

	// Quit signals the caller to quit.
	//
	// Cleanup code should not be run prior to this, and instead
	// should be run when the Close method of the Window is called.
	Quit()

	// Command sends an arbituary command to the caller.
	//
	// This is often used more for any custom Window proxies, for
	// example to broadcast a command to other Windows.
	Command(Command)

	// Buffer is the buffer holding the cells of the screen.
	//
	// These are usually not set directly, instead using the
	// `ivo/window` package to draw different components.
	Buffer() *Buffer

	// Render updates the screen with the contents of Buffer.
	Render()
}

// context is the core implementation of Context.
type context struct {
	buf     *Buffer
	expired bool
}

// newContext creates a new context, expiring the previous one.
func newContext() *context {
	if ctx != nil {
		ctx.expired = true
	}

	cols, rows := termbox.Size()
	buf := newBuffer(cols, rows)

	return &context{
		buf: buf,
	}
}

// Logger should be used to perform all logging.
func (ctx *context) Logger() Logger {
	return log
}

// Quit signals the caller to quit.
func (ctx *context) Quit() {
	quit = true
	go termbox.Interrupt()
}

// Command sends an arbituary command to the caller.
func (ctx *context) Command(c Command) {
	cmd = &c
	go termbox.Interrupt()
}

// Buffer is the buffer holding the cells of the screen.
func (ctx *context) Buffer() *Buffer {
	if ctx.expired {
		return nil
	}
	return ctx.buf
}

// Render updates the screen with the contents of Buffer.
func (ctx *context) Render() {
	if ctx.expired {
		return
	}

	for row := 0; row < ctx.buf.Rows; row++ {
		for col := 0; col < ctx.buf.Cols; col++ {
			c, ok := ctx.buf.Get(col, row)
			if !ok {
				continue
			}

			fg := termbox.Attribute(c.Fore)
			if c.Attr&CellAttrBold != 0 {
				fg |= termbox.AttrBold
			}
			if c.Attr&CellAttrUnderline != 0 {
				fg |= termbox.AttrUnderline
			}
			bg := termbox.Attribute(c.Back)
			termbox.SetCell(col, row, c.Rune, fg, bg)
		}
	}

	if err := termbox.Flush(); err != nil {
		log.Errorf("termbox: failed to flush: %v", err)
	}
}
