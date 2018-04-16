package key_test

import (
	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/key"
)

func newMap() *key.Map {
	m := key.NewMap()
	m.Set("insert", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'a'},
		{Code: ivo.KeyCodeRune, Rune: 'b'},
		{Code: ivo.KeyCodeRune, Rune: 'c'},
	}, nil)
	m.Set("insert", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'd'},
		{Code: ivo.KeyCodeEnter},
	}, nil)
	m.Set("insert", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'f'},
		{Code: ivo.KeyCodeRune, Rune: 'g'},
		{Code: ivo.KeyCodeRune, Rune: 'h', Mod: ivo.KeyModCtrl},
		{Code: ivo.KeyCodeRune, Rune: 'i'},
	}, nil)
	m.Set("normal", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'j'},
		{Code: ivo.KeyCodeEnter, Rune: 'k', Mod: ivo.KeyModAlt},
		{Code: ivo.KeyCodeRune, Rune: 'l'},
	}, nil)
	m.Set("normal", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'm'},
		{Code: ivo.KeyCodeRune, Rune: 'n'},
		{Code: ivo.KeyCodePgdn},
	}, nil)
	m.Set("", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'p'},
		{Code: ivo.KeyCodeRune, Rune: 'q'},
	}, nil)
	return m
}
