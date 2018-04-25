package text

import (
	"unicode"
)

// Text is some text stored as a rune slice.
//
// The operations on Text will panic with an out of bounds message
// if invalid Locations are provided. Fortunately, all of Text's
// methods will never return an invalid Location, so as long as no
// Location is created manually, panics should not occur.
type Text struct {
	rr         []rune
	lines      []int
	paragraphs []int
	sentences  []int
	words      []int
}

// NewText creates a new Text containing the raw rune slice provided.
func NewText(rr []rune) *Text {
	return &Text{
		rr:         append(rr, 0),
		lines:      lines(rr),
		paragraphs: paragraphs(rr),
		sentences:  sentences(rr),
		words:      words(rr),
	}
}

// Raw returns the raw rune slice contained in Textxt.
func (txt *Text) Raw() []rune {
	return txt.rr[:len(txt.rr)-1]
}

// RegionRaw returns the raw rune slice contained in Text, within the
// Region provided.
func (txt *Text) RegionRaw(reg Region) []rune {
	txt.check(reg.Begin)
	txt.check(reg.End)

	return txt.rr[reg.Begin:reg.End]
}

// Len returns the length of the rune slice.
func (txt *Text) Len() int {
	return len(txt.rr) - 1
}

// RemoveTrailingWhitespace returns a new Region without the trailing
// whitespace.
func (txt *Text) RemoveTrailingWhitespace(reg Region) Region {
	txt.check(reg.Begin)
	txt.check(reg.End)

	for i := reg.End - 1; i >= 0; i-- {
		if !whitespace(txt.rr[i]) {
			return Region{Begin: reg.Begin, End: Location(i + 1)}
		}
	}
	return Region{Begin: reg.Begin, End: reg.End}
}

// check checks whether the Location provided is within the bounds.
//
// If the Location is within the bounds, nothing will happen. If it
// isn't, then it will panic with an out of bounds message.
func (txt *Text) check(loc Location) {
	if int(loc) < 0 || int(loc) >= len(txt.rr) {
		panic("runtime error: index out of bounds")
	}
}

// lines splits the rune slice into lines, returning the indices of
// the first rune in each line.
func lines(rr []rune) []int {
	ii := make([]int, 1)

	// FIXME handle \r runes
	for _, r := range rr {
		ii[len(ii)-1]++
		if r == '\n' {
			ii = append(ii, 0)
		}
	}

	return ii
}

// paragraphs splits the rune slice into paragraphs, returning the
// indices of the sentences.
func paragraphs(rr []rune) []int {
	ii := make([]int, 1)
	count := 0

	for _, r := range rr {
		if count > 1 {
			ii = append(ii, 0)
			count = 0
		}

		ii[len(ii)-1]++

		if whitespace(r) {
			if r == '\n' {
				count++
			}
			continue
		}
	}

	return ii
}

// sentences splits the rune slice into sentences, returning the
// indices of the words.
func sentences(rr []rune) []int {
	ii := make([]int, 1)
	ended := false

	for _, r := range rr {
		if !whitespace(r) && ended {
			ii = append(ii, 0)
			ended = false
		}

		ii[len(ii)-1]++

		if ending(r) {
			ended = true
			continue
		}
	}

	return ii
}

// words splits the rune slice into words, returning the indices of
// the first rune in each word.
func words(rr []rune) []int {
	ii := make([]int, 1)
	ended := false

	for _, r := range rr {
		whitespace := whitespace(r)

		if !whitespace && ended {
			ii = append(ii, 0)
			ended = false
		}

		ii[len(ii)-1]++

		if whitespace {
			ended = true
			continue
		}
	}

	return ii
}

// whitespace returns whether the given rune is whitespace.
func whitespace(r rune) bool {
	return unicode.IsSpace(r)
}

// ending returns whether the given rune is a sentence ending.
func ending(r rune) bool {
	if r == '.' ||
		r == '\n' ||
		r == ';' ||
		r == '?' ||
		r == '!' {
		return true
	}
	return false
}

// index returns the beginning and end of the region of a location
// in the given index and the index of that region.
func index(ii []int, loc int) (int, int, int) {
	sum := 0
	for index, i := range ii {
		sum += i
		if sum <= int(loc) {
			continue
		}
		return index, sum - i, sum
	}
	return 0, 0, 0
}
