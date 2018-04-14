package text

func (t *Text) Insert(loc int, s string) bool {
	if loc < 0 || loc >= len(t.rr) {
		return false
	}
	tmp := append([]rune(s), t.rr[loc:]...)
	t.rr = append(t.rr[:loc], tmp...)
	return true
}
