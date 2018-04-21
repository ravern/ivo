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
