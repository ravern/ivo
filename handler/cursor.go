package handler

import "ivoeditor.com/ivo"

// Cursor represents a handler for cursor-related actions.
type Cursor interface {
	MoveBeginning(ivo.Context) // move to beginning
	MoveEnd(ivo.Context)       // move to end

	MoveNext(ivo.Context)          // move to next rune
	MoveNextWord(ivo.Context)      // move to next word
	MoveNextSentence(ivo.Context)  // move to next sentence
	MoveNextParagraph(ivo.Context) // move to next paragraph
	MoveNextLine(ivo.Context)      // move to next line

	MovePrev(ivo.Context)          // move to previous rune
	MovePrevWord(ivo.Context)      // move to previous word
	MovePrevSentence(ivo.Context)  // move to previous sentence
	MovePrevParagraph(ivo.Context) // move to previous paragraph
	MovePrevLine(ivo.Context)      // move to previous line

	SelectRune(ivo.Context)      // select current rune
	SelectWord(ivo.Context)      // select current word
	SelectSentence(ivo.Context)  // select current sentence
	SelectParagraph(ivo.Context) // select current paragraph
	SelectLine(ivo.Context)      // select current line
}
