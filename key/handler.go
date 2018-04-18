package key

import "ivoeditor.com/ivo"

// Handler represents a key handler that handles a successful key
// combination.
type Handler interface {
	Handle(ivo.Context, []ivo.Key)
}

// HandlerFunc is a convenience type to convert regular functions
// into Handlers.
type HandlerFunc func(ivo.Context, []ivo.Key)

func (f HandlerFunc) Handle(ctx ivo.Context, kk []ivo.Key) {
	f(ctx, kk)
}
