package key_test

import (
	"testing"

	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/key"
)

type mapping struct {
	mode string
	keys []ivo.Key
}

func TestMap_Get(t *testing.T) {
	tests := []struct {
		mappings  []mapping
		mode      string
		keys      []ivo.Key
		wantIndex int
		wantMore  bool
		wantOK    bool
	}{
		{
			mappings: []mapping{
				{
					mode: "insert",
					keys: []ivo.Key{
						{Rune: 'a'},
						{Rune: 'b'},
						{Rune: 'c'},
					},
				},
			},
			mode: "insert",
			keys: []ivo.Key{
				{Rune: 'a'},
				{Rune: 'b'},
			},
			wantIndex: -1,
			wantMore:  true,
			wantOK:    true,
		},
		{
			mappings: []mapping{
				{
					mode: "normal",
					keys: []ivo.Key{
						{Rune: 'a'},
						{Rune: 'b'},
						{Rune: 'c'},
						{Rune: 'd'},
					},
				},
				{
					mode: "normal",
					keys: []ivo.Key{
						{Rune: 'a'},
						{Rune: 'b'},
						{Rune: 'c'},
						{Rune: 'd'},
						{Rune: 'e'},
					},
				},
			},
			mode: "normal",
			keys: []ivo.Key{
				{Rune: 'a'},
				{Rune: 'b'},
				{Rune: 'c'},
				{Rune: 'd'},
			},
			wantIndex: 0,
			wantMore:  true,
			wantOK:    true,
		},
		{
			mappings: []mapping{
				{
					mode: "",
					keys: []ivo.Key{
						{Rune: 'a'},
					},
				},
				{
					mode: "normal",
					keys: []ivo.Key{
						{Rune: 'a'},
					},
				},
			},
			mode: "normal",
			keys: []ivo.Key{
				{Rune: 'a'},
			},
			wantIndex: 1,
			wantMore:  false,
			wantOK:    true,
		},
	}

	for i, test := range tests {
		gotIndex := -1

		m := key.NewMap()
		for i, mapping := range test.mappings {
			j := i
			m.Set(mapping.mode, mapping.keys, func(ivo.Context, []ivo.Key) {
				gotIndex = j
			})
		}

		gotHandler, gotMore, gotOK := m.Get(test.mode, test.keys)
		if test.wantIndex != -1 && gotHandler != nil {
			gotHandler(nil, nil)
		}

		if test.wantIndex != gotIndex || test.wantMore != gotMore || test.wantOK != gotOK {
			t.Errorf(
				"test %d: wanted %d, %v, %v got %d, %v, %v",
				i,
				test.wantIndex,
				test.wantMore,
				test.wantOK,
				gotIndex,
				gotMore,
				gotOK,
			)
		}
	}
}
