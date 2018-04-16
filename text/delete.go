package text

import "sort"

// Delete deletes n number of runes before the Location provided or
// up to the start of Text if n is larger than the Location.
func (t *Text) Delete(l Location, n int) {
	t.check(l)

	if int(l) < n {
		n = int(l)
	}
	t.rr = append(t.rr[:int(l)-n], t.rr[int(l):]...)
}

// DeleteMultiple performs the deletes at multiple Locations, taking
// into account the offset as a result of previous deletions.
func (t *Text) DeleteMultiple(ll []Location, n int) {
	sort.Sort(locationSlice(ll))

	var offset int
	for _, l := range ll {
		t.check(l)

		org := len(t.rr)
		t.Delete(Location(int(l)-offset), n)
		offset += org - len(t.rr)
	}
}
