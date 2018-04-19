package handler

import "ivoeditor.com/ivo"

// Proxy represents a handler for proxy-related actions.
type Proxy interface {
	Forward(ivo.Context, []ivo.Key) // forward raw runes
}
