package text

func (t *Text) Delete(loc int, n int) bool {
	if loc <= 0 || loc >= len(t.rr) {
		return false
	}
	if loc < n {
		n = loc
	}
	t.rr = append(t.rr[:loc-n], t.rr[loc:]...)
	return true
}
