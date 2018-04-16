package text

import (
	"sort"
)

func (t *Text) Insert(l Location, s string) bool {
	loc, ok := l.int(len(t.rr))
	if !ok {
		return false
	}

	tmp := append([]rune(s), t.rr[loc:]...)
	t.rr = append(t.rr[:loc], tmp...)
	return true
}

func (t *Text) InsertMultiple(ll []Location, s string) bool {
	sort.Sort(locationSlice(ll))

	var (
		offset  int
		success bool
	)
	for _, l := range ll {
		if t.Insert(Location(int(l)+offset), s) {
			success = true
			offset += len(s)
		}
	}
	return success
}
