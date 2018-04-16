package text

import (
	"sort"
)

func (t *Text) Insert(l Location, s string) {
	tmp := append([]rune(s), t.rr[int(l):]...)
	t.rr = append(t.rr[:int(l)], tmp...)
}

func (t *Text) InsertMultiple(ll []Location, s string) {
	sort.Sort(locationSlice(ll))
	var offset int
	for _, l := range ll {
		org := len(t.rr)
		t.Insert(Location(int(l)+offset), s)
		offset += len(t.rr) - org
	}
}
