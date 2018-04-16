package text

func (t *Text) Next(l Location) Location {
	t.check(l)

	if len(t.rr)-1 == int(l) {
		return l
	}
	return l + 1
}

func (t *Text) Prev(l Location) Location {
	t.check(l)

	if 0 == int(l) {
		return l
	}
	return l - 1
}
