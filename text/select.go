package text

// SelectParagraph returns the Region of the paragraph the Location
// is on.
func (txt *Text) SelectParagraph(loc Location) Region {
	txt.check(loc)
	begin, end := index(txt.paragraphs, int(loc))
	return Region{Begin: Location(begin), End: Location(end)}
}

// SelectParagraphInner returns the Region of the paragraph the Location
// is on, excluding the trailing whitespace.
func (txt *Text) SelectParagraphInner(loc Location) Region {
	return txt.RemoveTrailingWhitespace(txt.SelectParagraph(loc))
}

// SelectSentence returns the Region of the sentence the Location is
// on.
func (txt *Text) SelectSentence(loc Location) Region {
	txt.check(loc)
	begin, end := index(txt.sentences, int(loc))
	return Region{Begin: Location(begin), End: Location(end)}
}

// SelectSentenceInner returns the Region of the sentence the Location is
// on, excluding the trailing whitespace.
func (txt *Text) SelectSentenceInner(loc Location) Region {
	return txt.RemoveTrailingWhitespace(txt.SelectSentence(loc))
}

// SelectWord returns the Region of the word the Location is on.
func (txt *Text) SelectWord(loc Location) Region {
	txt.check(loc)
	begin, end := index(txt.words, int(loc))
	return Region{Begin: Location(begin), End: Location(end)}
}

// SelectWordInner returns the Region of the word the Location is on,
// excluding the trailing whitespace.
func (txt *Text) SelectWordInner(loc Location) Region {
	return txt.RemoveTrailingWhitespace(txt.SelectWord(loc))
}
