package handler

import "ivoeditor.com/ivo"

// Cursor represents a handler for cursor-related actions.
type Cursor interface {
	MoveBeginning(ivo.Context, []ivo.Key) // move to beginning
	MoveEnd(ivo.Context, []ivo.Key)       // move to end

	MoveNext(ivo.Context, []ivo.Key)     // move to next rune
	MoveNextWord(ivo.Context, []ivo.Key) // move to next word
	MoveNextLine(ivo.Context, []ivo.Key) // move to next line

	MovePrev(ivo.Context, []ivo.Key)     // move to previous rune
	MovePrevWord(ivo.Context, []ivo.Key) // move to previous word
	MovePrevLine(ivo.Context, []ivo.Key) // move to previous line

	SelectNext(ivo.Context, []ivo.Key)     // select next rune
	SelectNextWord(ivo.Context, []ivo.Key) // select next word
	SelectNextLine(ivo.Context, []ivo.Key) // select next line

	SelectPrev(ivo.Context, []ivo.Key)     // select previous rune
	SelectPrevWord(ivo.Context, []ivo.Key) // select previous word
	SelectPrevLine(ivo.Context, []ivo.Key) // select previous line
}
