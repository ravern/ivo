package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Prev(t *testing.T) {
	tests := []struct {
		text string
		loc  text.Location
		want text.Location
	}{
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  5,
			want: 4,
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  70,
			want: 69,
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  0,
			want: 0,
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		got := txt.Prev(test.loc)

		if test.want != got {
			t.Errorf("test %d: wanted %d got %d", i, test.want, got)
		}
	}
}

func TestText_PrevRune(t *testing.T) {
	tests := []struct {
		text string
		loc  text.Location
		r    rune
		want text.Location
	}{
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  24,
			r:    'e',
			want: 1,
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  1,
			r:    'l',
			want: 1,
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		got := txt.PrevRune(test.loc, test.r)

		if test.want != got {
			t.Errorf("test %d: wanted %d got %d", i, test.want, got)
		}
	}
}

func TestText_PrevLine(t *testing.T) {
	tests := []struct {
		text string
		loc  text.Location
		want text.Location
	}{
		{
			text: "I am writing some text.\nThis is for a test.\nOne more line.",
			loc:  29,
			want: 5,
		},
		{
			text: "Very basic text",
			loc:  5,
			want: 5,
		},
		{
			text: "I am writing some text.\nThis is for a test.",
			loc:  29,
			want: 5,
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		got := txt.PrevLine(test.loc)

		if test.want != got {
			t.Errorf("test %d: wanted %d got %d", i, test.want, got)
		}
	}
}
