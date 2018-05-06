package key

import (
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
	// 'k' key still isn't pressed, the handler for 'j' will be used.
	//
	// The default value of Timeout is 2 seconds.
	Timeout time.Duration

	// Mode is the mode that is used to get keys from Map.
	//
	// The default value of mode is an empty string.
	Mode string

	m *Map // contains all the key combinations

	// events is the channel to send new key events to.
	//
	// If events is nil, it means there is no processing goroutine
	// running in the background, waiting for events. Otherwise, the
	// goroutine should be started.
	events chan *event

	keys    []ivo.Key                    // current key buffer
	ctx     ivo.Context                  // current context
	handler func(ivo.Context, []ivo.Key) // current handler
}

// event represents a key press along with its ivo.Context.
type event struct {
	key ivo.Key
	ctx ivo.Context
}

// NewProcessor creates a new Processor.
func NewProcessor(m *Map) *Processor {
	return &Processor{
		Timeout: 2 * time.Second,
		m:       m,
		keys:    []ivo.Key{},
	}
}

// Process processes the next key.
func (p *Processor) Process(ctx ivo.Context, key ivo.Key) {
	// Start the next processing
	if p.events == nil {
		p.events = make(chan *event)
		go p.process()
	}

	// Send the next event
	p.events <- &event{
		ctx: ctx,
		key: key,
	}
}

// process runs a key event loop until a handler is either
// found or not found.
func (p *Processor) process() {
	// When the processing for this combination ends, reset
	// the processing state
	defer p.reset()

	// Loop and keep receiving keys until success or failure.
	for p.events != nil {
		// Wait for a key or timeout.
		select {
		case e := <-p.events:
			// Add the new key to the buffer.
			p.keys = append(p.keys, e.key)
			p.ctx = e.ctx
		case <-time.After(p.Timeout):
			// Call the handler and end the loop.
			p.events = nil
			p.execute()
			continue
		}

		// Get the corresponding handler for the keys in the
		// current buffer.
		var more, ok bool
		p.handler, more, ok = p.m.Get(p.Mode, p.keys)

		// If no handler is found, log the failure and reset.
		if !ok {
			p.events = nil
			p.ctx.Logger().Errorf("key: failed to find mapping for %v", p.keys)
			continue
		}

		// Since there are more possible handlers, poll the
		// next key.
		if more {
			continue
		}

		// Since a handler is found and there are no more
		// possible handlers, call the current one and reset.
		p.events = nil
		p.execute()
	}
}

// execute calls the current handler if it isn't nil.
func (p *Processor) execute() {
	if p.handler != nil {
		p.handler(p.ctx, p.keys)
	}
}

// reset resets the processing state.
func (p *Processor) reset() {
	p.ctx = nil
	p.keys = []ivo.Key{}
	p.handler = nil
}
