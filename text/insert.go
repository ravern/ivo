package text

// Insert inserts the given rune slice into the specified Location
// within the Text.
func (txt *Text) Insert(loc Location, rr []rune) Location {
	txt.check(loc)

	txt.rr = append(txt.rr[:loc], append(rr, txt.rr[loc:]...)...)

	i, begin, end := index(txt.lines, int(loc))
	lines := lines(txt.rr[begin : end+len(rr)])
	txt.lines = append(txt.lines[:i], append(lines, txt.lines[i+1:]...)...)

	i, begin, end = index(txt.paragraphs, int(loc))
	paragraphs := paragraphs(txt.rr[begin : end+len(rr)])
	txt.paragraphs = append(txt.paragraphs[:i], append(paragraphs, txt.paragraphs[i+1:]...)...)

	i, begin, end = index(txt.sentences, int(loc))
	sentences := sentences(txt.rr[begin : end+len(rr)])
	txt.sentences = append(txt.sentences[:i], append(sentences, txt.sentences[i+1:]...)...)

	i, begin, end = index(txt.words, int(loc))
	words := words(txt.rr[begin : end+len(rr)])
	txt.words = append(txt.words[:i], append(words, txt.words[i+1:]...)...)

	return Location(int(loc) + len(rr))
}

// InsertMultiple inserts the given rune slice into all the specified
// Locations within the Text.
//
// The given Location slice should be sorted in ascending order. This
// is important, or the offset will not be tracked properly.
//
// InsertMultiple takes into account the offset caused by previous
// insertions. For example, inserting 'h' to 'hello' at 0 and 3 will
// result in 'hhelhlo'.
func (txt *Text) InsertMultiple(locs []Location, rr []rune) []Location {
	newLocs := make([]Location, len(locs))

	offset := 0
	for i, loc := range locs {
		loc = Location(int(loc) + offset)
		txt.check(loc)

		newLoc := txt.Insert(loc, rr)
		newLocs[i] = newLoc

		offset += int(newLoc) - int(loc)
	}

	return newLocs
}
