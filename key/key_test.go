package key_test

import (
	"time"

	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/key"
)

type tracker struct {
	insertABC      bool
	insertDEnter   bool
	insertFGCtrlHI bool
	normalJAltKL   bool
	normalMNPgdn   bool
	rootPQ         bool
	rootFallback   bool
}

func newMap() (*key.Map, *tracker) {
	t := &tracker{}
	m := key.NewMap()

	m.Set("insert", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'a'},
		{Code: ivo.KeyCodeRune, Rune: 'b'},
		{Code: ivo.KeyCodeRune, Rune: 'c'},
	}, func(ctx ivo.Context) {
		t.insertABC = true
	})

	m.Set("insert", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'd'},
		{Code: ivo.KeyCodeEnter},
	}, func(ctx ivo.Context) {
		t.insertDEnter = true
	})

	m.Set("insert", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'f'},
		{Code: ivo.KeyCodeRune, Rune: 'g'},
		{Code: ivo.KeyCodeRune, Rune: 'h', Mod: ivo.KeyModCtrl},
		{Code: ivo.KeyCodeRune, Rune: 'i'},
	}, func(ctx ivo.Context) {
		t.insertFGCtrlHI = true
	})

	m.Set("normal", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'j'},
		{Code: ivo.KeyCodeEnter, Rune: 'k', Mod: ivo.KeyModAlt},
		{Code: ivo.KeyCodeRune, Rune: 'l'},
	}, func(ctx ivo.Context) {
		t.normalJAltKL = true
	})

	m.Set("normal", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'm'},
		{Code: ivo.KeyCodeRune, Rune: 'n'},
		{Code: ivo.KeyCodePgdn},
	}, func(ctx ivo.Context) {
		t.normalMNPgdn = true
	})

	m.Set("", []ivo.Key{
		{Code: ivo.KeyCodeRune, Rune: 'p'},
		{Code: ivo.KeyCodeRune, Rune: 'q'},
	}, func(ctx ivo.Context) {
		t.rootPQ = true
	})

	m.SetFallback("", func(ivo.Context) {
		t.rootFallback = true
	})

	return m, t
}

func newMapper() (*key.Mapper, *tracker) {
	m, t := newMap()
	mr := key.NewMapper(m)
	mr.Timeout = 100 * time.Millisecond
	return mr, t
}
