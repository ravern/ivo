package handler

import "ivoeditor.com/ivo"

// Text represents a handler for text-related actions.
type Text interface {
	Cut(ivo.Context, []ivo.Key)   // cut to clipboard
	Copy(ivo.Context, []ivo.Key)  // copy to clipboard
	Paste(ivo.Context, []ivo.Key) // paste from clipboard
	Raw(ivo.Context, []ivo.Key)   // handle raw runes
}
