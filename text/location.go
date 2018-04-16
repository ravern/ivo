package text

type Location int

type Region struct {
	Start Location
	End   Location
}

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
