package text_test

import (
	"reflect"
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Delete(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		n    int
		want text.Location
	}{
		{
			rr:   []rune{},
			loc:  0,
			n:    5,
			want: 0,
		},
		{
			rr:   []rune("Hello world!\n\nMy name is Jeff!"),
			loc:  18,
			n:    5,
			want: 13,
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.Delete(test.loc, test.n)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_DeleteMultiple(t *testing.T) {
	tests := []struct {
		rr   []rune
		locs []text.Location
		n    int
		want []text.Location
	}{
		{
			rr:   []rune{},
			locs: []text.Location{0},
			n:    200,
			want: []text.Location{0},
		},
		{
			rr:   []rune("Hello world!\n \nMy is Jeff!"),
			locs: []text.Location{10, 20},
			n:    8,
			want: []text.Location{2, 4},
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.DeleteMultiple(test.locs, test.n)

		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}
