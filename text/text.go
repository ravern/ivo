package text

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
		rr: []rune(s),
	}
}

// Len returns the length of the text.
func (t *Text) Len() int {
	return len(t.rr)
}

// String returns the value of the text as a string.
func (t *Text) String() string {
	return string(t.rr)
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
