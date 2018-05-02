package text

// Delete deletes n runes before the specified Location within
// the Text.
func (txt *Text) Delete(loc Location, n int) Location {
	txt.check(loc)

	if int(loc)-n < 0 {
		n = int(loc)
	}

	txt.rr = append(txt.rr[:int(loc)-n], txt.rr[int(loc):]...)

	i, begin, _ := index(txt.lines, int(loc)-n)
	j, _, end := index(txt.lines, int(loc))
	lines := buildLines(txt.rr[begin : end-n])
	txt.lines = append(txt.lines[:i], append(lines, txt.lines[j+1:]...)...)

	i, begin, _ = index(txt.paragraphs, int(loc)-n)
	j, _, end = index(txt.paragraphs, int(loc))
	paragraphs := buildParagraphs(txt.rr[begin : end-n])
	txt.paragraphs = append(txt.paragraphs[:i], append(paragraphs, txt.paragraphs[j+1:]...)...)

	i, begin, _ = index(txt.sentences, int(loc)-n)
	j, _, end = index(txt.sentences, int(loc))
	sentences := buildSentences(txt.rr[begin : end-n])
	txt.sentences = append(txt.sentences[:i], append(sentences, txt.sentences[j+1:]...)...)

	i, begin, _ = index(txt.words, int(loc)-n)
	j, _, end = index(txt.words, int(loc))
	words := buildWords(txt.rr[begin : end-n])
	txt.lines = append(txt.words[:i], append(words, txt.words[j+1:]...)...)

	return Location(int(loc) - n)
}

// DeleteMultiple n runes before all the specified Locations within
// the Text.
//
// The given Location slice should be sorted in ascending order. This
// is important, or the offset will not be tracked properly. The distance
// between each Location should not be smaller than n.
//
// DeleteMultiple takes into account the offset caused by previous
// deletions.  For example, deleting 1 from 'hello' at 0 and 3 will
// result in 'elo'.
func (txt *Text) DeleteMultiple(locs []Location, n int) []Location {
	newLocs := make([]Location, len(locs))

	offset := 0
	for i, loc := range locs {
		loc = Location(int(loc) - offset)
		txt.check(loc)

		newLoc := txt.Delete(loc, n)
		newLocs[i] = newLoc

		offset += int(loc) - int(newLoc)
	}

	return newLocs
}
