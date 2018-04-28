package handler

import "ivoeditor.com/ivo"

// KeyFunc converts a simple handler function into a key based one.
func KeyFunc(f func(ivo.Context)) func(ivo.Context, []ivo.Key) {
	return func(ctx ivo.Context, keys []ivo.Key) {
		f(ctx)
	}
}
