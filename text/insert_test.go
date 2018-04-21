package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Insert(t *testing.T) {
	tests := []struct {
		text string
		loc  text.Location
		s    string
		want string
	}{
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  5,
			s:    " great",
			want: "Hello great world! This is some sample text for the testing package for text.",
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  20,
			s:    " only",
			want: "Hello world! This is only some sample text for the testing package for text.",
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			loc:  0,
			s:    "Some frontmatter. ",
			want: "Some frontmatter. Hello world! This is some sample text for the testing package for text.",
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		txt.Insert(test.loc, test.s)
		got := txt.String()

		if test.want != got {
			t.Errorf("test %d: wanted %s got %s", i, test.want, got)
		}
	}
}

func TestText_InsertMultiple(t *testing.T) {
	tests := []struct {
		text string
		locs []text.Location
		s    string
		want string
	}{
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			locs: []text.Location{5, 25},
			s:    " great",
			want: "Hello great world! This is some great sample text for the testing package for text.",
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			locs: []text.Location{20, 70, 41},
			s:    " only",
			want: "Hello world! This is only some sample text for only the testing package for text only.",
		},
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			locs: []text.Location{0},
			s:    "Some frontmatter. ",
			want: "Some frontmatter. Hello world! This is some sample text for the testing package for text.",
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		txt.InsertMultiple(test.locs, test.s)
		got := txt.String()

		if test.want != got {
			t.Errorf("test %d: wanted %s got %s", i, test.want, got)
		}
	}
}
