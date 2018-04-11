package ivo

// Mouse is a mouse event.
type Mouse struct {
	// Action is the action that occured.
	Action MouseAction
	Col    int
	Row    int
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
