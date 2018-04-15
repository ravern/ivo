package text

func (t *Text) Delete(l Location, n int) bool {
	loc, ok := l.int(len(t.rr))
	if !ok {
		return false
	}

	if loc < n {
		n = loc
	}
	t.rr = append(t.rr[:loc-n], t.rr[loc:]...)
	return true
}
