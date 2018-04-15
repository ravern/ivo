package text

func (t *Text) Insert(l Location, s string) bool {
	loc, ok := l.int(len(t.rr))
	if !ok {
		return false
	}

	tmp := append([]rune(s), t.rr[loc:]...)
	t.rr = append(t.rr[:loc], tmp...)
	return true
}
