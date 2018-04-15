package text

type Location int

type Region struct {
	Start Location
	End   Location
}

func (l Location) int(max int) (int, bool) {
	loc := int(l)
	if loc < 0 || loc >= max {
		return 0, false
	}
	return loc, true
}
