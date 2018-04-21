package text_test

import (
	"reflect"
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_RegionString(t *testing.T) {
	tests := []struct {
		text string
		r    text.Region
		want string
	}{
		{
			text: "Hello world! This is some sample text for the testing package for text.",
			r:    text.Region{Start: 3, End: 16},
			want: "lo world! Thi",
		},
		{
			text: "Boom clap, sound of my heart, the beat goes on and on and on and on and...",
			r:    text.Region{Start: 30, End: 40},
			want: "the beat g",
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		got := txt.RegionString(test.r)

		if test.want != got {
			t.Errorf("test %d: wanted %s got %s", i, test.want, got)
		}
	}
}

func TestText_RegionLines(t *testing.T) {
	tests := []struct {
		text string
		r    text.Region
		want []string
	}{
		{
			text: "Hello\nDarkness\nMy\nOld\nFriend",
			r:    text.Region{Start: 3, End: 16},
			want: []string{"lo", "Darkness", "M"},
		},
		{
			text: "How many shrimps do you have to eat\nTo make your skin turn pink",
			r:    text.Region{Start: 30, End: 50},
			want: []string{"o eat", "To make your s"},
		},
	}

	for i, test := range tests {
		txt := text.New(test.text)
		got := txt.RegionLines(test.r)

		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("test %d: wanted %s got %s", i, test.want, got)
		}
	}
}
