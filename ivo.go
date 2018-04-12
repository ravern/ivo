package ivo

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

var (
	logger *log.Logger
	buffer Buffer
)

// SetLogger sets the logger. If the logger is not set, then a default logger will
// be used, which logs to os.Stdout.
func SetLogger(l *log.Logger) {
	logger = l
}

// SetBuffer sets the buffer. If the buffer is not set, Run will fail.
func SetBuffer(b Buffer) {
	buffer = b
}

// Run performs the main loop and blocks until the editor is quit.
func Run() {
	if buffer == nil {
		logger.Printf("core: buffer is nil")
		return
	}

	if err := termbox.Init(); err != nil {
		logger.Printf("termbox: could not initialize: %v", err)
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
			logger.Printf("termbox: polled error event: %v", e.Err)
		default:
			logger.Print("termbox: polled unknown event")
		}
	}
}
