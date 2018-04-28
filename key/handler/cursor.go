package handler

import "ivoeditor.com/ivo"

// Cursor represents a handler for cursor-related actions.
type Cursor interface {
	MoveBeginning(ivo.Context, []ivo.Key) // move to beginning
	MoveEnd(ivo.Context, []ivo.Key)       // move to end

	MoveNext(ivo.Context, []ivo.Key)          // move to next rune
	MoveNextWord(ivo.Context, []ivo.Key)      // move to next word
	MoveNextSentence(ivo.Context, []ivo.Key)  // move to next sentence
	MoveNextParagraph(ivo.Context, []ivo.Key) // move to next paragraph
	MoveNextLine(ivo.Context, []ivo.Key)      // move to next line

	MovePrev(ivo.Context, []ivo.Key)          // move to previous rune
	MovePrevWord(ivo.Context, []ivo.Key)      // move to previous word
	MovePrevSentence(ivo.Context, []ivo.Key)  // move to previous sentence
	MovePrevParagraph(ivo.Context, []ivo.Key) // move to previous paragraph
	MovePrevLine(ivo.Context, []ivo.Key)      // move to previous line

	SelectRune(ivo.Context, []ivo.Key)      // select current rune
	SelectWord(ivo.Context, []ivo.Key)      // select current word
	SelectSentence(ivo.Context, []ivo.Key)  // select current sentence
	SelectParagraph(ivo.Context, []ivo.Key) // select current paragraph
	SelectLine(ivo.Context, []ivo.Key)      // select current line
}
