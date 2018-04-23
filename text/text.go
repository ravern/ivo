package text

// Text is some text stored as a rune slice.
//
// The operations on Text will panic with an out of bounds message
// if invalid Locations are provided. Fortunately, all of Text's
// methods will never return an invalid Location, so as long as no
// Location is created manually, panics should not occur.
type Text struct {
	rr    []rune
	words [][][]int
	lines []int
}

// NewText creates a new Text containing the raw rune slice provided.
func NewText(rr []rune) *Text {
	t := &Text{rr: append(rr, 0)}
	// build indexes
	return t
}

// Raw returns the raw rune slice contained in Text.
func (t *Text) Raw() []rune {
	return t.rr[:len(t.rr)-1]
}

// RegionRaw returns the raw rune slice contained in Text, within the
// Region provided.
func (t *Text) RegionRaw(reg Region) []rune {
	t.check(reg.Begin)
	t.check(reg.End)
	return t.rr[reg.Begin:reg.End]
}

// Len returns the length of the rune slice.
func (t *Text) Len() int {
	return len(t.rr) - 1
}

// check checks whether the Location provided is within the bounds.
//
// If the Location is within the bounds, nothing will happen. If it
// isn't, then it will panic with an out of bounds message.
func (t *Text) check(loc Location) {
	if int(loc) < 0 || int(loc) >= len(t.rr) {
		panic("runtime error: index out of bounds")
	}
}
