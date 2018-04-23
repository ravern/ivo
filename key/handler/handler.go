// Package handler provides common handler types.
package handler

import "ivoeditor.com/ivo"

// Prompt represents a handler for prompt-related actions.
type Prompt interface {
	Confirm(ivo.Context, []ivo.Key) // confirm prompt
	Cancel(ivo.Context, []ivo.Key)  // cancel prompt
	Raw(ivo.Context, []ivo.Key)     // insert raw runes
}

// Proxy represents a handler for proxy-related actions.
type Proxy interface {
	Forward(ivo.Context, []ivo.Key) // forward raw runes
}

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

// Text represents a handler for text-related actions.
type Text interface {
	Cut(ivo.Context, []ivo.Key)   // cut to clipboard
	Copy(ivo.Context, []ivo.Key)  // copy to clipboard
	Paste(ivo.Context, []ivo.Key) // paste from clipboard
	Raw(ivo.Context, []ivo.Key)   // handle raw runes
}
