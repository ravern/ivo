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
		lines:      buildLines(rr),
		paragraphs: buildParagraphs(rr),
		sentences:  buildSentences(rr),
		words:      buildWords(rr),
	}
}

// Subtext returns the text in the given Region as a new Text.
func (txt *Text) Subtext(reg Region) *Text {
	txt.check(reg.Begin)
	txt.check(reg.End)

	begin, _, _ := index(txt.lines, int(reg.Begin))
	end, _, _ := index(txt.lines, int(reg.End))
	lines := txt.lines[begin:end]

	begin, _, _ = index(txt.paragraphs, int(reg.Begin))
	end, _, _ = index(txt.paragraphs, int(reg.End))
	paragraphs := txt.paragraphs[begin:end]

	begin, _, _ = index(txt.sentences, int(reg.Begin))
	end, _, _ = index(txt.sentences, int(reg.End))
	sentences := txt.sentences[begin:end]

	begin, _, _ = index(txt.words, int(reg.Begin))
	end, _, _ = index(txt.words, int(reg.End))
	words := txt.words[begin:end]

	return &Text{
		rr:         append(txt.rr[reg.Begin:reg.End], 0),
		lines:      lines,
		paragraphs: paragraphs,
		sentences:  sentences,
		words:      words,
	}
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
		if !isWhitespace(txt.rr[i]) {
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

// buildLines splits the rune slice into buildLines, returning the indices of
// the first rune in each line.
func buildLines(rr []rune) []int {
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

// buildParagraphs splits the rune slice into buildParagraphs, returning the
// indices of the sentences.
func buildParagraphs(rr []rune) []int {
	ii := make([]int, 1)
	count := 0

	for _, r := range rr {
		if count > 1 {
			ii = append(ii, 0)
			count = 0
		}

		ii[len(ii)-1]++

		if isWhitespace(r) {
			if r == '\n' {
				count++
			}
			continue
		}
	}

	return ii
}

// buildSentences splits the rune slice into buildSentences, returning the
// indices of the words.
func buildSentences(rr []rune) []int {
	ii := make([]int, 1)
	ended := false

	for _, r := range rr {
		if !isWhitespace(r) && ended {
			ii = append(ii, 0)
			ended = false
		}

		ii[len(ii)-1]++

		if isEnding(r) {
			ended = true
			continue
		}
	}

	return ii
}

// buildWords splits the rune slice into buildWords, returning the indices
// of the first rune in each word.
func buildWords(rr []rune) []int {
	ii := make([]int, 1)
	ended := false

	for _, r := range rr {
		whitespace := isWhitespace(r)

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

// isWhitespace returns whether the given rune is whitespace.
func isWhitespace(r rune) bool {
	return unicode.IsSpace(r)
}

// isEnding returns whether the given rune is a sentence ending.
func isEnding(r rune) bool {
	if r == '.' ||
		r == '\n' ||
		r == ';' ||
		r == '?' ||
		r == '!' {
		return true
	}
	return false
}

// index returns the index of the region containing the given
// location, and the beginning and end of that region.
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
