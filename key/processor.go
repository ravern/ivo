package key

import (
	"sync"
	"time"

	"ivoeditor.com/ivo"
)

// Processor provides state and timeouts for processing keys via Map.
//
// Processor runs its own loop in the background to process incoming
// key presses. For full use in a ivo.Window, all key events should
// be forwarded to the Processor, which will then process and call the
// handlers with the correct ivo.Contexts.
type Processor struct {
	// Timeout is how long before the handler of the current key
	// combination is used, instead of waiting for more keys.
	//
	// For example, if the Map contains handler for 'j' and 'jk',
	// and the key 'j' is processed, the Processor will then wait for
	// the duration of Timeout for the 'k' key. After waiting, if the
	// 'k' key still isn't pressed, the 'j' handler will be used.
	//
	// The default value of Timeout is 2 seconds.
	Timeout time.Duration

	// Mode is the mode that is used to get keys from Map.
	//
	// The default value of mode is an empty string.
	Mode string

	m      *Map
	init   sync.Once
	events chan *event
}

// event represents a key press along with its ivo.Context.
type event struct {
	ctx ivo.Context
	key ivo.Key
}

// NewProcessor creates a new Processor.
func NewProcessor(m *Map) *Processor {
	return &Processor{
		Timeout: 2 * time.Second,
		m:       m,
		events:  make(chan *event),
	}
}

// Process processes the key.
func (p *Processor) Process(ctx ivo.Context, k ivo.Key) {
	p.init.Do(func() {
		go p.process()
	})

	e := &event{
		ctx: ctx,
		key: k,
	}
	p.events <- e
}

// process is the key event loop.
func (p *Processor) process() {
	var (
		// kk is the current key buffer
		kk []ivo.Key

		// ctx is the latest context (to use on timeout)
		ctx ivo.Context

		// handler is latest handler (to use on timeout)
		handler func(ivo.Context, []ivo.Key)
	)

	// reset resets everything to their original values
	reset := func() {
		kk = make([]ivo.Key, 0)
		ctx = nil
		handler = nil
	}

	for {
		var k ivo.Key // current key

		// Wait for a key or timeout.
		select {
		case e := <-p.events:
			ctx = e.ctx
			k = e.key
		case <-time.After(p.Timeout):
			if handler != nil {
				handler(ctx, kk)
			}
			reset()
			continue
		}

		// Add the new key to the buffer.
		kk = append(kk, k)

		// Get the corresponding handler for the keys in the
		// current buffer.
		var more, ok bool
		handler, more, ok = p.m.Get(p.Mode, kk)

		// If no handler is found, log it and reset.
		if !ok {
			ctx.Logger().Errorf("key: failed to find mapping for %v", kk)
			reset()
			continue
		}

		// Since there are more possible handlers, poll the
		// next key.
		if more {
			continue
		}

		// Since a handler is found and there are no more
		// possible handlers, run the current one and reset.
		if handler != nil {
			handler(ctx, kk)
		}
		reset()
	}
}
