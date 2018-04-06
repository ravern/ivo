package input

import "strings"

type Key int

const (
	KeyF1 Key = iota + 256
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyInsert
	KeyDelete
	KeyHome
	KeyEnd
	KeyPgup
	KeyPgdn
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyEsc
	KeyEnter
	KeyBackspace
	KeyTab
	KeySpace
)

func (k Key) String() string {
	switch k {
	case KeyF1:
		return "f1"
	case KeyF2:
		return "f2"
	case KeyF3:
		return "f3"
	case KeyF4:
		return "f4"
	case KeyF5:
		return "f5"
	case KeyF6:
		return "f6"
	case KeyF7:
		return "f7"
	case KeyF8:
		return "f8"
	case KeyF9:
		return "f9"
	case KeyF10:
		return "f10"
	case KeyF11:
		return "f11"
	case KeyF12:
		return "f12"
	case KeyInsert:
		return "insert"
	case KeyDelete:
		return "delete"
	case KeyHome:
		return "home"
	case KeyEnd:
		return "end"
	case KeyPgup:
		return "pgup"
	case KeyPgdn:
		return "pgdn"
	case KeyArrowUp:
		return "arrow up"
	case KeyArrowDown:
		return "arrow down"
	case KeyArrowLeft:
		return "arrow left"
	case KeyArrowRight:
		return "arrow right"
	case KeyEsc:
		return "esc"
	case KeyEnter:
		return "enter"
	case KeyBackspace:
		return "backspace"
	case KeyTab:
		return "tab"
	case KeySpace:
		return "space"
	default:
		return rune(k)
	}
}

type Mod int

const (
	ModNone Mod = 0
	ModCtrl Mod = 1 << (iota - 1)
	ModAlt
)

func (m Mod) String() string {
	if m == ModNone {
		return "none"
	}
	var mods []string
	if m&ModCtrl != 0 {
		mods = append(mods, "ctrl")
	}
	if m&ModAlt != 0 {
		mods = append(mods, "alt")
	}
	return strings.Join(mods, ", ")
}
