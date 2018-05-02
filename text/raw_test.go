package text_test

import (
	"reflect"
	"testing"

	"ivoeditor.com/ivo/text"
)

func TestText_Raw(t *testing.T) {
	tests := [][]rune{
		[]rune{},
		[]rune("Hello world!"),
		[]rune("This is the first sentence. This is the second."),
		[]rune("This is the first paragraph.\n\nThis is the second."),
		[]rune("Boom clap sound of my heart"),
		[]rune("I'm in love with the shape of you"),
		[]rune("Oh I oh I oh I oh I\nOh I oh I oh I oh I"),
	}

	for i, test := range tests {
		txt := text.NewText(test)
		got := txt.Raw()

		if !reflect.DeepEqual(test, got) {
			t.Errorf("test %d: want %s got %s", i, string(test), string(got))
		}
	}
}

func TestText_RegionRaw(t *testing.T) {
	tests := []struct {
		rr   []rune
		reg  text.Region
		want []rune
	}{
		{
			rr:   []rune{},
			reg:  text.Region{Begin: 0, End: 0},
			want: []rune{},
		},
		{
			rr:   []rune("Hello world!"),
			reg:  text.Region{Begin: 3, End: 12},
			want: []rune("lo world!"),
		},
		{
			rr:   []rune("This is the first sentence. This is the second."),
			reg:  text.Region{Begin: 5, End: 19},
			want: []rune("is the first s"),
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.RegionRaw(test.reg)

		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("test %d: want %s got %s", i, string(test.want), string(got))
		}
	}
}

func TestText_RawLines(t *testing.T) {
	tests := []struct {
		rr   []rune
		want [][]rune
	}{
		{
			rr: []rune("Hello world!\nMy name is Ravern."),
			want: [][]rune{
				[]rune("Hello world!\n"),
				[]rune("My name is Ravern."),
			},
		},
		{
			rr: []rune("This is the first sentence.\nThis is the second."),
			want: [][]rune{
				[]rune("This is the first sentence.\n"),
				[]rune("This is the second."),
			},
		},
	}

	for i, test := range tests {
		txt := text.NewText(test.rr)
		got := txt.RawLines()

		if !reflect.DeepEqual(test.want, got) {
			wantStrs := make([]string, len(test.want))
			for _, rr := range test.want {
				wantStrs = append(wantStrs, string(rr))
			}

			gotStrs := make([]string, len(got))
			for _, rr := range test.want {
				gotStrs = append(gotStrs, string(rr))
			}

			t.Errorf("test %d: want %v got %v", i, wantStrs, gotStrs)
		}
	}
}
