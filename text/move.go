package text

// MoveNext returns the Location of the next rune or the end
// Location if the end has been reached.
func (txt *Text) MoveNext(loc Location) Location {
	txt.check(loc)
	if int(loc) == len(txt.rr)-1 {
		return loc
	}
	return Location(int(loc) + 1)
}

// MoveNextWord returns the Location of the first rune of the
// next word, or the end Location if the end has been reached.
func (txt *Text) MoveNextWord(loc Location) Location {
	txt.check(loc)
	_, _, end := index(txt.words, int(loc))
	if int(end) == len(txt.rr)-1 {
		return Location(end)
	}
	return Location(end)
}

// MoveNextSentence returns the Location of the first rune of the
// next word, or the end Location if the end has been reached.
func (txt *Text) MoveNextSentence(loc Location) Location {
	txt.check(loc)
	_, _, end := index(txt.sentences, int(loc))
	if int(end) == len(txt.rr)-1 {
		return Location(end)
	}
	return Location(end)
}

// MoveNextParagraph returns the Location of the first rune of
// the next word, or the end Location if the end has been reached.
func (txt *Text) MoveNextParagraph(loc Location) Location {
	txt.check(loc)
	_, _, end := index(txt.paragraphs, int(loc))
	if int(end) == len(txt.rr)-1 {
		return Location(end)
	}
	return Location(end)
}
