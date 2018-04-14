package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Delete(t *testing.T) {
	tests := []struct {
		loc  int
		n    int
		want string
	}{
		{
			loc:  5,
			n:    3,
			want: "He world! This is some sample text for the testing package for text.",
		},
		{
			loc:  20,
			n:    10,
			want: "Hello worl some sample text for the testing package for text.",
		},
		{
			loc:  -1,
			n:    4238478947,
			want: "Hello world! This is some sample text for the testing package for text.",
		},
	}

	for i, test := range tests {
		txt := text.New(textStr)
		txt.Delete(test.loc, test.n)
		got := txt.String()

		if test.want != got {
			t.Errorf("test %d: wanted %s got %s", i, test.want, got)
		}
	}
}
