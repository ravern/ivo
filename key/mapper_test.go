package key_test

import (
	"testing"
	"time"

	"ivoeditor.com/ivo"
	"ivoeditor.com/ivo/key"
	"ivoeditor.com/ivo/mock"
)

type event struct {
	key   ivo.Key
	sleep time.Duration
}

func TestMapper_Process(t *testing.T) {
	tests := []struct {
		mappings []mapping
		mode     string
		events   []event
		timeout  time.Duration
		want     int
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
			events: []event{
				{
					key:   ivo.Key{Rune: 'a'},
					sleep: 1 * time.Millisecond,
				},
				{
					key:   ivo.Key{Rune: 'b'},
					sleep: 1 * time.Millisecond,
				},
			},
			timeout: 10 * time.Millisecond,
			want:    -1,
		},
		{
			mappings: []mapping{
				{
					mode: "insert",
					keys: []ivo.Key{
						{Rune: 'a'},
						{Rune: 'b'},
					},
				},
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
			events: []event{
				{
					key:   ivo.Key{Rune: 'a'},
					sleep: 1 * time.Millisecond,
				},
				{
					key:   ivo.Key{Rune: 'b'},
					sleep: 1 * time.Millisecond,
				},
				{
					key:   ivo.Key{Rune: 'c'},
					sleep: 1 * time.Millisecond,
				},
			},
			timeout: 10 * time.Millisecond,
			want:    1,
		},
		{
			mappings: []mapping{
				{
					mode: "insert",
					keys: []ivo.Key{
						{Rune: 'a'},
						{Rune: 'b'},
					},
				},
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
			events: []event{
				{
					key:   ivo.Key{Rune: 'a'},
					sleep: 1 * time.Millisecond,
				},
				{
					key:   ivo.Key{Rune: 'b'},
					sleep: 15 * time.Millisecond,
				},
				{
					key:   ivo.Key{Rune: 'c'},
					sleep: 1 * time.Millisecond,
				},
			},
			timeout: 10 * time.Millisecond,
			want:    0,
		},
	}

	for i, test := range tests {
		got := -1

		m := key.NewMap()
		for i, mapping := range test.mappings {
			j := i
			m.Set(mapping.mode, mapping.keys, func(ivo.Context, []ivo.Key) {
				got = j
			})
		}

		mr := key.NewMapper(m)
		mr.Timeout = test.timeout
		mr.Mode = test.mode

		for _, event := range test.events {
			mr.Process(mock.NewContext(), event.key)
			time.Sleep(event.sleep)
		}

		if test.want != got {
			t.Errorf("test %d: want %d, got %d", i, test.want, got)
		}
	}
}
