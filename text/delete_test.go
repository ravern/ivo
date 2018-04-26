package text_test

import (
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
