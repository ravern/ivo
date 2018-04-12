package ivo

import (
	termbox "github.com/nsf/termbox-go"
)

var (
	logger Logger
	window Window
)

// SetLogger sets the logger. If the logger is not set, then a default logger will
// be used, which logs to os.Stdout.
func SetLogger(l Logger) {
	logger = l
}

// SetWindow sets the window. If the window is not set, Run will fail.
func SetWindow(w Window) {
	window = w
}

// Run performs the main loop and blocks until the editor quits.
func Run() {
	if logger == nil {
		logger = defaultLogger
	}
	if window == nil {
		logger.Errorf("core: window is nil")
		return
	}

	if err := termbox.Init(); err != nil {
		logger.Errorf("termbox: could not initialize: %v", err)
		return
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	termbox.SetOutputMode(termbox.OutputNormal)

	for {
		data := make([]byte, 32)
		switch e := termbox.PollRawEvent(data); e.Type {
		case termbox.EventRaw:
			data := data[:e.N]
			e := termbox.ParseEvent(data)
			if e.Type == termbox.EventNone {
				e.Type = termbox.EventKey
				e.Key = termbox.KeyEsc
			}
			switch e.Type {
			case termbox.EventKey:
				newKey(e)
			case termbox.EventMouse:
				break
			}
		case termbox.EventResize:
			break
		case termbox.EventInterrupt:
			break
		case termbox.EventError:
			logger.Errorf("termbox: polled error event: %v", e.Err)
		default:
			logger.Errorf("termbox: polled unknown event")
		}
	}
}
