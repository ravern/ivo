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
		{Rune: 'a'},
		{Rune: 'b'},
		{Rune: 'c'},
	}, key.Handler(func(ctx ivo.Context, kk []ivo.Key) {
		t.insertABC = true
	}))

	m.Set("insert", []ivo.Key{
		{Rune: 'd'},
		{Code: ivo.KeyCodeEnter},
	}, key.Handler(func(ctx ivo.Context, kk []ivo.Key) {
		t.insertDEnter = true
	}))

	m.Set("insert", []ivo.Key{
		{Rune: 'f'},
		{Rune: 'g'},
		{Rune: 'h', Mod: ivo.KeyModCtrl},
		{Rune: 'i'},
	}, key.Handler(func(ctx ivo.Context, kk []ivo.Key) {
		t.insertFGCtrlHI = true
	}))

	m.Set("normal", []ivo.Key{
		{Rune: 'j'},
		{Rune: 'k', Mod: ivo.KeyModAlt},
		{Rune: 'l'},
	}, key.Handler(func(ctx ivo.Context, kk []ivo.Key) {
		t.normalJAltKL = true
	}))

	m.Set("normal", []ivo.Key{
		{Rune: 'm'},
		{Rune: 'n'},
		{Code: ivo.KeyCodePgdn},
	}, key.Handler(func(ctx ivo.Context, kk []ivo.Key) {
		t.normalMNPgdn = true
	}))

	m.Set("", []ivo.Key{
		{Rune: 'p'},
		{Rune: 'q'},
	}, key.Handler(func(ctx ivo.Context, kk []ivo.Key) {
		t.rootPQ = true
	}))

	m.SetFallback("", key.Handler(func(ctx ivo.Context, kk []ivo.Key) {
		t.rootFallback = true
	}))

	return m, t
}

func newMapper() (*key.Mapper, *tracker) {
	m, t := newMap()
	mr := key.NewMapper(m)
	mr.Timeout = 100 * time.Millisecond
	return mr, t
}
