package ivo

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

type Core struct {
	Logger *log.Logger
	Buffer Buffer
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) Run() {
	if err := termbox.Init(); err != nil {
		c.Logger.Printf("termbox: could not initialize: %v", err)
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
			c.Logger.Printf("termbox: polled error event: %v", e.Err)
		default:
			c.Logger.Print("termbox: polled unknown event")
		}
	}
}
