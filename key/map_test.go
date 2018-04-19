package key_test

import (
	"testing"

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
	}, func(ctx ivo.Context, kk []ivo.Key) {
		t.insertABC = true
	})

	m.Set("insert", []ivo.Key{
		{Rune: 'd'},
		{Code: ivo.KeyCodeEnter},
	}, func(ctx ivo.Context, kk []ivo.Key) {
		t.insertDEnter = true
	})

	m.Set("insert", []ivo.Key{
		{Rune: 'f'},
		{Rune: 'g'},
		{Rune: 'h', Mod: ivo.KeyModCtrl},
		{Rune: 'i'},
	}, func(ctx ivo.Context, kk []ivo.Key) {
		t.insertFGCtrlHI = true
	})

	m.Set("normal", []ivo.Key{
		{Rune: 'j'},
		{Rune: 'k', Mod: ivo.KeyModAlt},
		{Rune: 'l'},
	}, func(ctx ivo.Context, kk []ivo.Key) {
		t.normalJAltKL = true
	})

	m.Set("normal", []ivo.Key{
		{Rune: 'm'},
		{Rune: 'n'},
		{Code: ivo.KeyCodePgdn},
	}, func(ctx ivo.Context, kk []ivo.Key) {
		t.normalMNPgdn = true
	})

	m.Set("", []ivo.Key{
		{Rune: 'p'},
		{Rune: 'q'},
	}, func(ctx ivo.Context, kk []ivo.Key) {
		t.rootPQ = true
	})

	m.SetFallback("", func(ctx ivo.Context, kk []ivo.Key) {
		t.rootFallback = true
	})

	return m, t
}

func TestMap_Get(t *testing.T) {
	tests := []struct {
		mode     string
		keys     []ivo.Key
		wantMore bool
		wantOK   bool
	}{
		{
			mode: "insert",
			keys: []ivo.Key{
				{Rune: 'a'},
				{Rune: 'b'},
				{Rune: 'c'},
			},
			wantMore: false,
			wantOK:   true,
		},
		{
			mode: "insert",
			keys: []ivo.Key{
				{Rune: 'a'},
			},
			wantMore: true,
			wantOK:   true,
		},
		{
			mode: "insert",
			keys: []ivo.Key{
				{Rune: 'd'},
			},
			wantMore: true,
			wantOK:   true,
		},
		{
			mode: "insert",
			keys: []ivo.Key{
				{Rune: 'd'},
				{Code: ivo.KeyCodeEnter},
			},
			wantMore: false,
			wantOK:   true,
		},
		{
			mode: "normal",
			keys: []ivo.Key{
				{Rune: 'm'},
				{Rune: 'n'},
			},
			wantMore: true,
			wantOK:   true,
		},
		{
			mode: "",
			keys: []ivo.Key{
				{Rune: 'x'},
				{Rune: 'y'},
				{Rune: 'z'},
			},
			wantMore: false,
			wantOK:   true,
		},
	}

	for i, test := range tests {
		m, _ := newMap()
		_, gotMore, gotOK := m.Get(test.mode, test.keys)

		if test.wantMore != gotMore || test.wantOK != gotOK {
			t.Errorf("test %d: wanted %v, %v got %v, %v", i, test.wantMore, test.wantOK, gotMore, gotOK)
		}
	}
}
