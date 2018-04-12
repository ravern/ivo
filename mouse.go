package ivo

import termbox "github.com/nsf/termbox-go"

// Mouse is a mouse event.
type Mouse struct {
	// Action is the action that occured.
	Action MouseAction

	Col int // x coordinate
	Row int // y coordinate
}

// newMouse creates a new Mouse based on the values found in the termbox.Event.
func newMouse(e termbox.Event) Mouse {
	var m Mouse
	m.Col = e.MouseX
	m.Row = e.MouseY
	switch e.Key {
	case termbox.MouseLeft:
		m.Action = MouseButtonLeft
	case termbox.MouseMiddle:
		m.Action = MouseButtonMiddle
	case termbox.MouseRight:
		m.Action = MouseButtonRight
	case termbox.MouseRelease:
		m.Action = MouseButtonRelease
	case termbox.MouseWheelUp:
		m.Action = MouseWheelUp
	case termbox.MouseWheelDown:
		m.Action = MouseWheelDown
	}
	return m
}

// MouseAction represents a mouse button, scroll or release.
type MouseAction int

// Supported mouse actions.
const (
	MouseButtonLeft MouseAction = iota
	MouseButtonMiddle
	MouseButtonRight
	MouseButtonRelease
	MouseWheelUp
	MouseWheelDown
)

func (ma MouseAction) String() string {
	switch ma {
	case MouseButtonLeft:
		return "left"
	case MouseButtonMiddle:
		return "middle"
	case MouseButtonRight:
		return "right"
	case MouseButtonRelease:
		return "release"
	case MouseWheelUp:
		return "wheelUp"
	case MouseWheelDown:
		return "wheelDown"
	}
	return "invalid"
}
