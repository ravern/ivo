package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Len(t *testing.T) {
	tests := []struct {
		rr   []rune
		want int
	}{
		{
			rr:   []rune("Hello world!"),
			want: 12,
		},
		{
			rr:   []rune("Boom clap sound of my heart"),
			want: 27,
		},
		{
			rr:   []rune("I'm in love with the shape of you"),
			want: 33,
		},
		{
			rr:   []rune("Oh I oh I oh I oh I"),
			want: 19,
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.Len()

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}
