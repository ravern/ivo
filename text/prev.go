package text

// Prev returns the Location of a rune to the left of the Location
// provided, or the same Location if the rune is unavailable.
func (t *Text) Prev(l Location) Location {
	t.check(l)

	if 0 == int(l) {
		return l
	}
	return l - 1
}

// PrevRune returns the Location of the next occurrence of the given
// rune to the left of the Location provided, or the same Location
// if the rune does not exist.
func (t *Text) PrevRune(l Location, r rune) Location {
	t.check(l)

	for i := int(l) - 1; i >= 0; i-- {
		if t.rr[i] == r {
			return Location(i)
		}
	}
	return l
}

// PrevLine returns the Location on the previous line with the same
// horizontal offset of the Location provided, the same Location
// if the previous line does not exist, or the last Location of the
// previous line if it's shorter than the current one.
func (t *Text) PrevLine(l Location) Location {
	t.check(l)

	maxLoc := int(t.PrevRune(l, '\n'))
	if maxLoc == int(l) {
		return l
	}

	minLoc := int(t.PrevRune(Location(maxLoc), '\n'))
	if minLoc == maxLoc {
		minLoc = 0
	}

	prevLoc := int(t.PrevRune(l, '\n'))
	if prevLoc == int(l) {
		prevLoc = 0
	} else {
		prevLoc++
	}

	return Location(min(minLoc+int(l)-prevLoc, maxLoc))
}
