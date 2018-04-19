package key

import (
	"sync"
	"time"

	"ivoeditor.com/ivo"
)

// Mapper provides state and timeouts for processing keys via Map.
//
// Mapper runs its own loop in the background to process incoming
// key presses. For full use in a ivo.Window, all key events should
// be forwarded to the Mapper, which will then process and call the
// handlers with the correct ivo.Contexts.
type Mapper struct {
	// Timeout is how long before the handler of the current key
	// combination is used, instead of waiting for more keys.
	//
	// For example, if the Map contains handler for 'j' and 'jk',
	// and the key 'j' is processed, the Mapper will then wait for
	// the duration of Timeout for the 'k' key. After waiting, if the
	// 'k' key still isn't pressed, the 'j' handler will be used.
	//
	// The default value of Timeout is 2 seconds.
	Timeout time.Duration

	// Mode is the mode that is used to get keys from Map.
	//
	// The default value of mode is an empty string.
	Mode string

	m    *Map
	init sync.Once
	ctxs chan ivo.Context
	keys chan ivo.Key
}

// NewMapper creates a new Mapper.
func NewMapper(m *Map) *Mapper {
	return &Mapper{
		Timeout: 2 * time.Second,
		m:       m,
		ctxs:    make(chan ivo.Context),
		keys:    make(chan ivo.Key),
	}
}

// Process sends the key to the background loop for processing.
func (mr *Mapper) Process(ctx ivo.Context, k ivo.Key) {
	mr.init.Do(func() {
		go mr.process()
	})
	mr.ctxs <- ctx
	mr.keys <- k
}

// process is the key event loop.
func (mr *Mapper) process() {
	var (
		kk      []ivo.Key
		ctx     ivo.Context
		handler func(ivo.Context, []ivo.Key)
	)

	reset := func() {
		kk = make([]ivo.Key, 0)
		ctx = nil
		handler = nil
	}

	for {
		var k ivo.Key

		if len(kk) > 0 {
			// More keys awaiting
			select {
			case ctx = <-mr.ctxs:
				k = <-mr.keys
			case <-time.After(mr.Timeout):
				if handler != nil {
					handler(ctx, kk)
				}
				reset()
				continue
			}
		} else {
			// New key
			ctx = <-mr.ctxs
			k = <-mr.keys
		}

		kk = append(kk, k)
		handler, more, ok := mr.m.Get(mr.Mode, kk)

		if !ok {
			ctx.Logger().Errorf("key: failed to find mapping for %v", kk)
			reset()
			continue
		}
		if more {
			continue
		}

		if handler != nil {
			handler(ctx, kk)
		}
		reset()
	}
}
