package text_test

import (
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Next(t *testing.T) {
	tests := []struct {
		loc  text.Location
		want text.Location
	}{
		{
			loc:  5,
			want: 6,
		},
		{
			loc:  70,
			want: 70,
		},
		{
			loc:  30,
			want: 31,
		},
	}

	for i, test := range tests {
		txt := text.New(textStr)
		got := txt.Next(test.loc)

		if test.want != got {
			t.Errorf("test %d: wanted %d got %d", i, test.want, got)
		}
	}
}

func TestText_Prev(t *testing.T) {
	tests := []struct {
		loc  text.Location
		want text.Location
	}{
		{
			loc:  5,
			want: 4,
		},
		{
			loc:  70,
			want: 69,
		},
		{
			loc:  0,
			want: 0,
		},
	}

	for i, test := range tests {
		txt := text.New(textStr)
		got := txt.Prev(test.loc)

		if test.want != got {
			t.Errorf("test %d: wanted %d got %d", i, test.want, got)
		}
	}
}
