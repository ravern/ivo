package text

import (
	"sort"
)

// Insert inserts the given string after the Location provided.
func (t *Text) Insert(l Location, s string) {
	t.check(l)

	tmp := append([]rune(s), t.rr[int(l):]...)
	t.rr = append(t.rr[:int(l)], tmp...)
}

// InsertMultiple performs the inserts at multiple Locations, taking
// into account the offset as a result of previous insertions.
func (t *Text) InsertMultiple(ll []Location, s string) {
	sort.Sort(locationSlice(ll))

	var offset int
	for _, l := range ll {
		t.check(l)

		org := len(t.rr)
		t.Insert(Location(int(l)+offset), s)
		offset += len(t.rr) - org
	}
}
