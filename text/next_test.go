package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Next(t *testing.T) {
	tests := []struct {
		text string
		loc  text.Location
		want text.Location
	}{
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  5,
			want: 6,
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  71,
			want: 71,
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  30,
			want: 31,
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		got := txt.Next(test.loc)

		if test.want != got {
			t.Errorf("test %d: wanted %d got %d", i, test.want, got)
		}
	}
}

func TestText_NextRune(t *testing.T) {
	tests := []struct {
		text string
		loc  text.Location
		r    rune
		want text.Location
	}{
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  5,
			r:    'e',
			want: 24,
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  70,
			r:    'l',
			want: 70,
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		got := txt.NextRune(test.loc, test.r)

		if test.want != got {
			t.Errorf("test %d: wanted %d got %d", i, test.want, got)
		}
	}
}

func TestText_NextLine(t *testing.T) {
	tests := []struct {
		text string
		loc  text.Location
		want text.Location
	}{
		{
			text: "I am writing some text.\nThis is for a test.\nOne more line.",
			loc:  5,
			want: 29,
		},
		{
			text: "Very basic text",
			loc:  5,
			want: 5,
		},
		{
			text: "I am writing some text.\nThis is for a test.",
			loc:  5,
			want: 29,
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		got := txt.NextLine(test.loc)

		if test.want != got {
			t.Errorf("test %d: wanted %d got %d", i, test.want, got)
		}
	}
}
