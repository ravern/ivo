package handler

import "ivoeditor.com/ivo"

// Cursor represents a handler for cursor-related actions.
type Cursor interface {
	Beginning(ivo.Context, []ivo.Key) // go to beginning
	End(ivo.Context, []ivo.Key)       // go to end
	Next(ivo.Context, []ivo.Key)      // next rune
	NextWord(ivo.Context, []ivo.Key)  // next word
	NextLine(ivo.Context, []ivo.Key)  // next line
	Prev(ivo.Context, []ivo.Key)      // previous rune
	PrevWord(ivo.Context, []ivo.Key)  // previous word
	PrevLine(ivo.Context, []ivo.Key)  // previous line
}
