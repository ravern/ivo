package ivo

import (
	"strings"

	termbox "github.com/nsf/termbox-go"
)

// Key is a key press event.
type Key struct {
	// Code is the key that was pressed. If the key is not special, then it will be
	// set to KeyCodeRune and the raw value is set on Rune.
	Code KeyCode

	// Rune is the raw value of the key press. It will not be set unless Code is set to
	// KeyCodeRune.
	Rune rune

	// Mod contains any modifier keys that were pressed. The modifiers are applied as
	// masks. If no modifiers were pressed then the KeyModNone value is assigned.
	Mod KeyMod
}

// newKey creates a new Key based on the values found in the termbox.Event. It does not
// handle the aliasing problem.
func newKey(e termbox.Event) Key {
	var k Key
	if e.Ch == 0 {
		switch e.Key {
		case termbox.KeyF1:
			k.Code = KeyCodeF1
		case termbox.KeyF2:
			k.Code = KeyCodeF2
		case termbox.KeyF3:
			k.Code = KeyCodeF3
		case termbox.KeyF4:
			k.Code = KeyCodeF4
		case termbox.KeyF5:
			k.Code = KeyCodeF5
		case termbox.KeyF6:
			k.Code = KeyCodeF6
		case termbox.KeyF7:
			k.Code = KeyCodeF7
		case termbox.KeyF8:
			k.Code = KeyCodeF8
		case termbox.KeyF9:
			k.Code = KeyCodeF9
		case termbox.KeyF10:
			k.Code = KeyCodeF10
		case termbox.KeyF11:
			k.Code = KeyCodeF11
		case termbox.KeyF12:
			k.Code = KeyCodeF12
		case termbox.KeyInsert:
			k.Code = KeyCodeInsert
		case termbox.KeyDelete:
			k.Code = KeyCodeDelete
		case termbox.KeyHome:
			k.Code = KeyCodeHome
		case termbox.KeyEnd:
			k.Code = KeyCodeEnd
		case termbox.KeyPgup:
			k.Code = KeyCodePgup
		case termbox.KeyPgdn:
			k.Code = KeyCodePgdn
		case termbox.KeyArrowUp:
			k.Code = KeyCodeArrowUp
		case termbox.KeyArrowDown:
			k.Code = KeyCodeArrowDown
		case termbox.KeyArrowLeft:
			k.Code = KeyCodeArrowLeft
		case termbox.KeyArrowRight:
			k.Code = KeyCodeArrowRight
		case termbox.KeyCtrlSpace:
			k.Code = KeyCodeSpace
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlA:
			k.Rune = 'a'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlB:
			k.Rune = 'b'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlC:
			k.Rune = 'c'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlD:
			k.Rune = 'd'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlE:
			k.Rune = 'e'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlF:
			k.Rune = 'f'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlG:
			k.Rune = 'g'
			k.Mod = KeyModCtrl
		case termbox.KeyBackspace:
			k.Code = KeyCodeBackspace
		case termbox.KeyTab:
			k.Code = KeyCodeTab
		case termbox.KeyCtrlJ:
			k.Rune = 'j'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlK:
			k.Rune = 'k'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlL:
			k.Rune = 'l'
			k.Mod = KeyModCtrl
		case termbox.KeyEnter:
			k.Code = KeyCodeEnter
		case termbox.KeyCtrlN:
			k.Rune = 'n'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlO:
			k.Rune = 'o'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlP:
			k.Rune = 'p'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlQ:
			k.Rune = 'q'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlR:
			k.Rune = 'r'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlS:
			k.Rune = 's'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlT:
			k.Rune = 't'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlU:
			k.Rune = 'u'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlV:
			k.Rune = 'v'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlW:
			k.Rune = 'w'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlX:
			k.Rune = 'x'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlY:
			k.Rune = 'y'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrlZ:
			k.Rune = 'z'
			k.Mod = KeyModCtrl
		case termbox.KeyEsc:
			k.Code = KeyCodeEsc
		case termbox.KeyCtrl4:
			k.Rune = '4'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrl5:
			k.Rune = '5'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrl6:
			k.Rune = '6'
			k.Mod = KeyModCtrl
		case termbox.KeyCtrl7:
			k.Rune = '7'
			k.Mod = KeyModCtrl
		case termbox.KeySpace:
			k.Code = KeyCodeSpace
		case termbox.KeyBackspace2:
			k.Code = KeyCodeBackspace
		}
	} else {
		k.Rune = e.Ch
	}
	if e.Mod&termbox.ModAlt != 0 {
		k.Mod |= KeyModAlt
	}
	return k
}

// KeyCode represents a special key.
type KeyCode int

// Supported special keys.
const (
	KeyCodeRune KeyCode = iota
	KeyCodeF1
	KeyCodeF2
	KeyCodeF3
	KeyCodeF4
	KeyCodeF5
	KeyCodeF6
	KeyCodeF7
	KeyCodeF8
	KeyCodeF9
	KeyCodeF10
	KeyCodeF11
	KeyCodeF12
	KeyCodeInsert
	KeyCodeDelete
	KeyCodeHome
	KeyCodeEnd
	KeyCodePgup
	KeyCodePgdn
	KeyCodeArrowUp
	KeyCodeArrowDown
	KeyCodeArrowLeft
	KeyCodeArrowRight
	KeyCodeEsc
	KeyCodeEnter
	KeyCodeBackspace
	KeyCodeTab
	KeyCodeSpace
)

func (kc KeyCode) String() string {
	switch kc {
	case KeyCodeF1:
		return "f1"
	case KeyCodeF2:
		return "f2"
	case KeyCodeF3:
		return "f3"
	case KeyCodeF4:
		return "f4"
	case KeyCodeF5:
		return "f5"
	case KeyCodeF6:
		return "f6"
	case KeyCodeF7:
		return "f7"
	case KeyCodeF8:
		return "f8"
	case KeyCodeF9:
		return "f9"
	case KeyCodeF10:
		return "f10"
	case KeyCodeF11:
		return "f11"
	case KeyCodeF12:
		return "f12"
	case KeyCodeInsert:
		return "insert"
	case KeyCodeDelete:
		return "delete"
	case KeyCodeHome:
		return "home"
	case KeyCodeEnd:
		return "end"
	case KeyCodePgup:
		return "pgup"
	case KeyCodePgdn:
		return "pgdn"
	case KeyCodeArrowUp:
		return "arrowUp"
	case KeyCodeArrowDown:
		return "arrowDown"
	case KeyCodeArrowLeft:
		return "arrowLeft"
	case KeyCodeArrowRight:
		return "arrowRight"
	case KeyCodeEsc:
		return "esc"
	case KeyCodeEnter:
		return "enter"
	case KeyCodeBackspace:
		return "backspace"
	case KeyCodeTab:
		return "tab"
	case KeyCodeSpace:
		return "space"
	}
	return "invalid"
}

// KeyMod represents a modifier key.
type KeyMod int

// Supported modifier keys.
const (
	KeyModNone KeyMod = 0
	KeyModCtrl KeyMod = 1 << (iota - 1)
	KeyModAlt
)

func (km KeyMod) String() string {
	if km == KeyModNone {
		return "none"
	}
	var mods []string
	if km&KeyModCtrl != 0 {
		mods = append(mods, "ctrl")
	}
	if km&KeyModAlt != 0 {
		mods = append(mods, "alt")
	}
	return strings.Join(mods, ", ")
}
