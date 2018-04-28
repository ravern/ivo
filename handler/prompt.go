package handler

import "ivoeditor.com/ivo"

// Prompt represents a handler for prompt-related actions.
type Prompt interface {
	Confirm(ivo.Context)        // confirm prompt
	Cancel(ivo.Context)         // cancel prompt
	Raw(ivo.Context, []ivo.Key) // insert raw runes
}
