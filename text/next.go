package text

// Next returns the Location of a rune to the right of the Location
// provided, or the same Location if the rune is unavailable.
func (t *Text) Next(l Location) Location {
	t.check(l)

	if len(t.rr)-1 == int(l) {
		return l
	}
	return l + 1
}

// NextRune returns the Location of the next occurrence of the given
// rune to the right of the Location provided, or the same Location
// if the rune does not exist.
func (t *Text) NextRune(l Location, r rune) Location {
	t.check(l)

	for i := int(l) + 1; i < len(t.rr); i++ {
		if t.rr[i] == r {
			return Location(i)
		}
	}
	return l
}

// NextLine returns the Location on the next line with the same
// horizontal offset of the Location provided, the same Location
// if the next line does not exist, or the last Location of the
// next line if it's shorter than the current one.
func (t *Text) NextLine(l Location) Location {
	t.check(l)

	minLoc := int(t.NextRune(l, '\n'))
	if minLoc == int(l) {
		return l
	}
	minLoc++

	maxLoc := int(t.NextRune(Location(minLoc), '\n'))
	if maxLoc == minLoc {
		maxLoc = len(t.rr) - 1
	}

	prevLoc := int(t.PrevRune(l, '\n'))
	if prevLoc == int(l) {
		prevLoc = 0
	} else {
		prevLoc++
	}

	return Location(min(minLoc+int(l)-prevLoc, maxLoc))
}
