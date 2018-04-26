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
	lines := lines(txt.rr[begin : end-n])
	txt.lines = append(txt.lines[:i], append(lines, txt.lines[j+1:]...)...)

	i, begin, _ = index(txt.paragraphs, int(loc)-n)
	j, _, end = index(txt.paragraphs, int(loc))
	paragraphs := paragraphs(txt.rr[begin : end-n])
	txt.paragraphs = append(txt.paragraphs[:i], append(paragraphs, txt.paragraphs[j+1:]...)...)

	i, begin, _ = index(txt.sentences, int(loc)-n)
	j, _, end = index(txt.sentences, int(loc))
	sentences := sentences(txt.rr[begin : end-n])
	txt.sentences = append(txt.sentences[:i], append(sentences, txt.sentences[j+1:]...)...)

	i, begin, _ = index(txt.words, int(loc)-n)
	j, _, end = index(txt.words, int(loc))
	words := words(txt.rr[begin : end-n])
	txt.lines = append(txt.words[:i], append(words, txt.words[j+1:]...)...)

	return Location(int(loc) - n)
}
