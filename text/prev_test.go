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
