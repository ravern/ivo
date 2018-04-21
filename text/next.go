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
