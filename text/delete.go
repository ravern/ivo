package text

import "sort"

func (t *Text) Delete(l Location, n int) {
	t.check(l)

	if int(l) < n {
		n = int(l)
	}
	t.rr = append(t.rr[:int(l)-n], t.rr[int(l):]...)
}

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
