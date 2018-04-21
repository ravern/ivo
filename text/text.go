package text

import "strings"

// Text is some text, stored internally as a rune slice.
//
// The operations on Text will panic with an out of bounds message
// if invalid Locations are provided. Fortunately, all of Text's
// methods will never return an invalid Location, so as long as no
// Location is created manually, panics should not occur.
type Text struct {
	rr []rune
}

// New creates a new Text containing the string value provided.
func New(s string) *Text {
	return &Text{
		rr: append([]rune(s), 0),
	}
}

// Len returns the length of the text.
func (t *Text) Len() int {
	return len(t.rr) - 1
}

// String returns the value of the text.
func (t *Text) String() string {
	return string(t.rr[:len(t.rr)-1])
}

// RegionString returns the value of a region of the text.
func (t *Text) RegionString(r Region) string {
	t.check(r.Start)
	t.check(r.End)
	if int(r.End) == len(t.rr)-1 {
		r.End = Location(len(t.rr) - 2)
	}
	return string(t.rr[r.Start:r.End])
}

// Lines returns the value of the text, split into lines.
func (t *Text) Lines() []string {
	return strings.Split(t.String(), "\n")
}

// RegionLines returns the value of a region of the text, split into
// lines.
func (t *Text) RegionLines(r Region) []string {
	return strings.Split(t.RegionString(r), "\n")
}

// check checks whether the Location provided is within the bounds.
//
// If the Location is within the bounds, nothing will happen. If it
// isn't, then it will panic with an out of bounds message.
func (t *Text) check(l Location) {
	if int(l) < 0 || int(l) >= len(t.rr) {
		panic("runtime error: index out of bounds")
	}
}
