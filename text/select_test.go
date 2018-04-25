package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_SelectParagraph(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Region
	}{
		{
			rr:   []rune("Hello world!\n \nMy name is Jeff!"),
			loc:  13,
			want: text.Region{Begin: 0, End: 15},
		},
		{
			rr:   []rune("Boom clap sound of my heart.\n\nThe beat goes on and on and on and on and."),
			loc:  32,
			want: text.Region{Begin: 30, End: 72},
		},
		{
			rr:   []rune("I'm in love with the shape of you"),
			loc:  13,
			want: text.Region{Begin: 0, End: 33},
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.SelectParagraph(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_SelectParagraphInner(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Region
	}{
		{
			rr:   []rune("Hello world!\n \nMy name is Jeff!"),
			loc:  13,
			want: text.Region{Begin: 0, End: 12},
		},
		{
			rr:   []rune("Boom clap sound of my heart.\n\nThe beat goes on and on and on and on and."),
			loc:  32,
			want: text.Region{Begin: 30, End: 72},
		},
		{
			rr:   []rune("I'm in love with the shape of you        "),
			loc:  13,
			want: text.Region{Begin: 0, End: 33},
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.SelectParagraphInner(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_SelectSentence(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Region
	}{
		{
			rr:   []rune("Hello world!\n \nMy name is Jeff!"),
			loc:  13,
			want: text.Region{Begin: 0, End: 15},
		},
		{
			rr:   []rune("Boom clap sound of my heart.\n\nThe beat goes on and on and on and on and."),
			loc:  32,
			want: text.Region{Begin: 30, End: 72},
		},
		{
			rr:   []rune("I'm in love with the shape of you. We push and pull like a magnet do."),
			loc:  13,
			want: text.Region{Begin: 0, End: 35},
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.SelectSentence(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_SelectSentenceInner(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Region
	}{
		{
			rr:   []rune("Hello world!\n \nMy name is Jeff!"),
			loc:  13,
			want: text.Region{Begin: 0, End: 12},
		},
		{
			rr:   []rune("Boom clap sound of my heart.\n\nThe beat goes on and on and on and on and."),
			loc:  32,
			want: text.Region{Begin: 30, End: 72},
		},
		{
			rr:   []rune("I'm in love with the shape of you. We push and pull like a magnet do."),
			loc:  13,
			want: text.Region{Begin: 0, End: 34},
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.SelectSentenceInner(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_SelectWord(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Region
	}{
		{
			rr:   []rune("Hello world!\n \nMy name is Jeff!"),
			loc:  6,
			want: text.Region{Begin: 6, End: 15},
		},
		{
			rr:   []rune("Boom clap sound of my heart.\n\nThe beat goes on and on and on and on and."),
			loc:  32,
			want: text.Region{Begin: 30, End: 34},
		},
		{
			rr:   []rune("I'm in love with the shape of you. We push and pull like a magnet do."),
			loc:  41,
			want: text.Region{Begin: 38, End: 43},
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.SelectWord(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}

func TestText_SelectWordInner(t *testing.T) {
	tests := []struct {
		rr   []rune
		loc  text.Location
		want text.Region
	}{
		{
			rr:   []rune("Hello world!\n \nMy name is Jeff!"),
			loc:  6,
			want: text.Region{Begin: 6, End: 12},
		},
		{
			rr:   []rune("Boom clap sound of my heart.\n\nThe beat goes on and on and on and on and."),
			loc:  32,
			want: text.Region{Begin: 30, End: 33},
		},
		{
			rr:   []rune("I'm in love with the shape of you. We push and pull like a magnet do."),
			loc:  41,
			want: text.Region{Begin: 38, End: 42},
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.SelectWordInner(test.loc)

		if test.want != got {
			t.Errorf("test %d: want %d got %d", i, test.want, got)
		}
	}
}
