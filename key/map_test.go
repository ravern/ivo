package key_test

import (
	"testing"

	"ivoeditor.com/ivo"
)

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
				{Code: ivo.KeyCodeRune, Rune: 'a'},
				{Code: ivo.KeyCodeRune, Rune: 'b'},
				{Code: ivo.KeyCodeRune, Rune: 'c'},
			},
			wantMore: false,
			wantOK:   true,
		},
		{
			mode: "insert",
			keys: []ivo.Key{
				{Code: ivo.KeyCodeRune, Rune: 'a'},
			},
			wantMore: true,
			wantOK:   true,
		},
		{
			mode: "insert",
			keys: []ivo.Key{
				{Code: ivo.KeyCodeRune, Rune: 'd'},
			},
			wantMore: true,
			wantOK:   true,
		},
		{
			mode: "insert",
			keys: []ivo.Key{
				{Code: ivo.KeyCodeRune, Rune: 'd'},
				{Code: ivo.KeyCodeEnter},
			},
			wantMore: false,
			wantOK:   true,
		},
		{
			mode: "normal",
			keys: []ivo.Key{
				{Code: ivo.KeyCodeRune, Rune: 'm'},
				{Code: ivo.KeyCodeRune, Rune: 'n'},
			},
			wantMore: true,
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
