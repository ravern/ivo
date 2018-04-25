package text_test

import (
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
