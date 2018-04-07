package ivo

type Mouse struct {
	Button MouseButton
	Col    int
	Row    int
}

type MouseButton int

const (
	MouseButtonLeft MouseButton = iota
	MouseButtonMiddle
	MouseButtonRight
	MouseButtonRelease
	MouseWheelUp
	MouseWheelDown
)

func (mb MouseButton) String() string {
	switch mb {
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
