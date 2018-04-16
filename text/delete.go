package text

import "sort"

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

func (t *Text) DeleteMultiple(ll []Location, n int) bool {
	sort.Sort(locationSlice(ll))

	var (
		offset  int
		success bool
	)
	for _, l := range ll {
		org := len(t.rr)
		if t.Delete(Location(int(l)-offset), n) {
			success = true
			offset += org - len(t.rr)
		}
	}
	return success
}
