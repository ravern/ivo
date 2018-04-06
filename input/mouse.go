package input

type Button int

const (
	ButtonLeft Button = iota
	ButtonMiddle
	ButtonRight
)

func (b Button) String() string {
	switch b {
	case ButtonLeft:
		return "left"
	case ButtonMiddle:
		return "middle"
	case ButtonRight:
		return "right"
	}
	return "invalid"
}

type Direction int

const (
	DirectionUp Direction = iota
	DirectionDown
)

func (d Direction) String() string {
	switch d {
	case DirectionUp:
		return "up"
	case DirectionDown:
		return "down"
	}
	return "invalid"
}
