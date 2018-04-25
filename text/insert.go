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
