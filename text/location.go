package text

// Location represents a position in some Text.
type Location int

// Region represents a region of some Text.
type Region struct {
	Begin Location
	End   Location
}
