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
