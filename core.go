package ivo

import (
	"log"

	termbox "github.com/nsf/termbox-go"
)

type Core struct {
	Logger *log.Logger
	KeyMap *KeyMap

	data []byte // data of raw event polling
}

func NewCore() *Core {
	return &Core{
		data: make([]byte, 32),
	}
}

func (c *Core) Run() {
	if err := termbox.Init(); err != nil {
		c.Logger.Printf("termbox: could not initialize: %v", err)
		return
	}
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputAlt | termbox.InputMouse)

	for {
		switch e := c.pollEvent(); e.Type {
		case termbox.EventKey:
			newKey(e)
		case termbox.EventMouse:
			break
		}
	}
}

func (c *Core) pollEvent() termbox.Event {
	for {
		switch e := termbox.PollRawEvent(c.data); e.Type {
		case termbox.EventRaw:
			data := c.data[:e.N]
			e := termbox.ParseEvent(data)
			if e.Type == termbox.EventNone {
				e.Type = termbox.EventKey
				e.Key = termbox.KeyEsc
			}
			return e
		case termbox.EventResize:
			break
		case termbox.EventError:
			c.Logger.Printf("termbox: polled error event: %v", e.Err)
		default:
			c.Logger.Print("termbox: polled unknwon event")
		}
	}
}
