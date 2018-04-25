package text_test

import (
	"reflect"
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Insert(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		rrr  []rune
		want text.Location
	}{
		{
			rr:   []rune{},
			loc:  0,
			rrr:  []rune("Hello world!"),
			want: 12,
		},
		{
			rr:   []rune("Hello world!\n \nMy is Jeff!"),
			loc:  18,
			rrr:  []rune("name "),
			want: 23,
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.Insert(test.loc, test.rrr)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_InsertMultiple(t *testing.T) {
	tests := []struct {
		rr   []rune
		locs []text.Location
		rrr  []rune
		want []text.Location
	}{
		{
			rr:   []rune{},
			locs: []text.Location{0},
			rrr:  []rune("Hello world!"),
			want: []text.Location{12},
		},
		{
			rr:   []rune("Hello world!\n \nMy is Jeff!"),
			locs: []text.Location{18, 20},
			rrr:  []rune("name "),
			want: []text.Location{23, 30},
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.InsertMultiple(test.locs, test.rrr)

		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}
