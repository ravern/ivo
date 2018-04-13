package ivo

import (
	termbox "github.com/nsf/termbox-go"
)

var (
	log Logger   // main logger
	win Window   // main window
	ctx *context // current context
	cmd *Command // command to be sent
)

var (
	started bool // whether main loop has started
	quit    bool // whether main loop should end
)

func init() {
	log = newLogger()
}

// SetLogger sets the logger.
//
// The logger set will be passed into the Window via the Context.
// This will be the logger used for all logging in the core, and
// it should also be used by the Window.
//
// If the logger is not set, then a default log will be used,
// which logs to os.Stderr. Once the main loop is started, this
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
	if err := termbox.Init(); err != nil {
		log.Errorf("termbox: could not initialize: %v", err)
		return
	}
	defer termbox.Close()

	if win == nil {
		log.Errorf("core: win is nil")
		return
	}
	defer win.Close(newContext())

	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)
	termbox.SetOutputMode(termbox.OutputNormal)

	started = true
	for !quit {
		if cmd != nil {
			win.Command(newContext(), *cmd)
			cmd = nil
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
				win.Key(newContext(), newKey(e))
			case termbox.EventMouse:
				win.Mouse(newContext(), newMouse(e))
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
