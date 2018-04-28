package handler

import "ivoeditor.com/ivo"

// Text represents a handler for text-related actions.
type Text interface {
	Cut(ivo.Context)            // cut to clipboard
	Copy(ivo.Context)           // copy to clipboard
	Paste(ivo.Context)          // paste from clipboard
	Raw(ivo.Context, []ivo.Key) // handle raw runes
}
