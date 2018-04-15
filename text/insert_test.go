package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Insert(t *testing.T) {
	tests := []struct {
		loc  text.Location
		s    string
		want string
	}{
		{
			loc:  5,
			s:    " great",
			want: "Hello great world! This is some sample text for the testing package for text.",
		},
		{
			loc:  20,
			s:    " only",
			want: "Hello world! This is only some sample text for the testing package for text.",
		},
		{
			loc:  -1,
			s:    "invalid",
			want: "Hello world! This is some sample text for the testing package for text.",
		},
		{
			loc:  0,
			s:    "Some frontmatter. ",
			want: "Some frontmatter. Hello world! This is some sample text for the testing package for text.",
		},
	}

	for i, test := range tests {
		txt := text.New(textStr)
		txt.Insert(test.loc, test.s)
		got := txt.String()

		if test.want != got {
			t.Errorf("test %d: wanted %s got %s", i, test.want, got)
		}
	}
}
