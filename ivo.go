package ivo

import (
	termbox "github.com/nsf/termbox-go"
)

var (
	log Logger
	win Window
	ctx *context
	cmd *Command
)

var (
	started bool
	quit    bool
)

func init() {
	// TODO set logger
}

// SetLogger sets the logger.
//
// If the logger is not set, then a default log will be used,
// which logs to os.Stdout. Once the main loop is started, this
// won't do anything.
func SetLogger(l Logger) {
	if started {
		return
	}
	log = l
}

// SetWindow sets the window.
//
// If the window is not set, Run will fail. Once the main loop
// is started, this won't do anything.
func SetWindow(w Window) {
	if started {
		return
	}
	win = w
}

// Run runs the main loop.
//
// Run will block until the ivo or the Window quits.
func Run() {
	if win == nil {
		log.Errorf("core: win is nil")
		return
	}
	defer win.Close(newContext())

	if err := termbox.Init(); err != nil {
		log.Errorf("termbox: could not initialize: %v", err)
		return
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	termbox.SetOutputMode(termbox.OutputNormal)

	for !quit {
		if cmd != nil {
			// TODO perform command
		}

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
				// TODO perform key
			case termbox.EventMouse:
				// TODO perform mouse
			}
		case termbox.EventResize:
			break
		case termbox.EventInterrupt:
			break
		case termbox.EventError:
			log.Errorf("termbox: polled error event: %v", e.Err)
		default:
			log.Errorf("termbox: polled unknown event")
		}
	}
}
