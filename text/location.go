package text

// Location represents a location within some Text.
type Location int

// Region represents two locations (and the region between them)
// within some Text.
type Region struct {
	Start Location
	End   Location
}

// locationSlice implements sort.Interface for a slice of Locations.
type locationSlice []Location

func (l locationSlice) Len() int {
	return len(l)
}

func (l locationSlice) Less(i int, j int) bool {
	return l[i] < l[j]
}

func (l locationSlice) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}
