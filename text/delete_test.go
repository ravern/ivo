package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Delete(t *testing.T) {
	tests := []struct {
		text string
		loc  text.Location
		n    int
		want string
	}{
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  5,
			n:    3,
			want: "He world! This is some sample text for the testing package for text.",
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  20,
			n:    10,
			want: "Hello worl some sample text for the testing package for text.",
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  0,
			n:    2748923473,
			want: "Hello world! This is some sample text for the testing package for text.",
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		txt.Delete(test.loc, test.n)
		got := txt.String()

		if test.want != got {
			t.Errorf("test %d: wanted %s got %s", i, test.want, got)
		}
	}
}

func TestText_DeleteMultiple(t *testing.T) {
	tests := []struct {
		text string
		locs []text.Location
		n    int
		want string
	}{
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			locs: []text.Location{5, 25},
			n:    5,
			want: " world! This is sample text for the testing package for text.",
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			locs: []text.Location{49, 22},
			n:    10,
			want: "Hello world!ome sample text fting package for text.",
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			locs: []text.Location{2, 20},
			n:    10,
			want: "llo worl some sample text for the testing package for text.",
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		txt.DeleteMultiple(test.locs, test.n)
		got := txt.String()

		if test.want != got {
			t.Errorf("test %d: wanted %s got %s", i, test.want, got)
		}
	}
}
