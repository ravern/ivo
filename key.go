package ivo

import (
	"strconv"
	"strings"
)

type Key struct {
	Code KeyCode
	Rune rune
	Mod  KeyMod
}

func (k Key) hash() string {
	if k.Code == KeyCodeRune {
		return string(k.Rune) + strconv.Itoa(int(k.Mod))
	}
	return strconv.Itoa(int(k.Code)) + strconv.Itoa(int(k.Mod))
}

type KeyCode int

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

type KeyMod int

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
