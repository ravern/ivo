package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_MoveNext(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Location
	}{
		{
			rr:   []rune{},
			loc:  0,
			want: 0,
		},
		{
			rr:   []rune("Hello world!\n \nMy is Jeff!"),
			loc:  18,
			want: 19,
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.MoveNext(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_MoveNextWord(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Location
	}{
		{
			rr:   []rune{},
			loc:  0,
			want: 0,
		},
		{
			rr:   []rune("Hello world!\n \nMy is Jeff!"),
			loc:  18,
			want: 21,
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.MoveNextWord(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_MoveNextSentence(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Location
	}{
		{
			rr:   []rune{},
			loc:  0,
			want: 0,
		},
		{
			rr:   []rune("Hello world!\n \nMy is Jeff!"),
			loc:  5,
			want: 15,
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.MoveNextSentence(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_MoveNextParagraph(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Location
	}{
		{
			rr:   []rune{},
			loc:  0,
			want: 0,
		},
		{
			rr:   []rune("Hello world!\n \nMy is Jeff!"),
			loc:  5,
			want: 15,
		},
		{
			rr:   []rune("Hello world!"),
			loc:  3,
			want: 12,
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.MoveNextParagraph(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}
