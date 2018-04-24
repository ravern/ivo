package text

// SelectParagraph returns the Region of the paragraph the Location
// is on.
func (t *Text) SelectParagraph(loc Location) Region {
	t.check(loc)

	begin, end, ok := index(t.paragraphs, int(loc))
	if !ok {
		panic("runtime error: paragraphs index broken")
	}
	return Region{Begin: Location(begin), End: Location(end)}
}

// SelectSentence returns the Region of the sentence the Location is
// on.
func (t *Text) SelectSentence(loc Location) Region {
	t.check(loc)

	begin, end, ok := index(t.sentences, int(loc))
	if !ok {
		panic("runtime error: sentences index broken")
	}
	return Region{Begin: Location(begin), End: Location(end)}
}

// SelectWord returns the Region of the word the Location is on.
func (t *Text) SelectWord(loc Location) Region {
	t.check(loc)

	begin, end, ok := index(t.words, int(loc))
	if !ok {
		panic("runtime error: words index broken")
	}
	return Region{Begin: Location(begin), End: Location(end)}
}
