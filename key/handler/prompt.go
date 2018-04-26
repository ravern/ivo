package handler

import "ivoeditor.com/ivo"

// Prompt represents a handler for prompt-related actions.
type Prompt interface {
	Confirm(ivo.Context, []ivo.Key) // confirm prompt
	Cancel(ivo.Context, []ivo.Key)  // cancel prompt
	Raw(ivo.Context, []ivo.Key)     // insert raw runes
}
