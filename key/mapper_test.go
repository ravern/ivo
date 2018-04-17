package key_test

import (
	"testing"
	"time"

	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/mock"
)

func TestMapper_Process(t *testing.T) {
	tests := []struct {
		mode     string
		keys     []ivo.Key
		interval time.Duration
		check    func(*tracker) bool
	}{
		{
			mode: "insert",
			keys: []ivo.Key{
				{Code: ivo.KeyCodeRune, Rune: 'a'},
				{Code: ivo.KeyCodeRune, Rune: 'b'},
				{Code: ivo.KeyCodeRune, Rune: 'c'},
			},
			interval: time.Millisecond,
			check: func(t *tracker) bool {
				return t.insertABC
			},
		},
		{
			mode: "insert",
			keys: []ivo.Key{
				{Code: ivo.KeyCodeRune, Rune: 'd'},
				{Code: ivo.KeyCodeEnter},
			},
			interval: 200 * time.Millisecond,
			check: func(t *tracker) bool {
				return !t.insertDEnter
			},
		},
	}

	for i, test := range tests {
		mr, tr := newMapper()
		mr.Mode = test.mode

		for _, key := range test.keys {
			mr.Process(mock.NewContext(), key)
			time.Sleep(test.interval)
		}

		if !test.check(tr) {
			t.Errorf("test %d: failed check", i)
		}
	}
}
