package ivo

// Buffer represents the main content.
type Buffer interface {
	// Quit is called when the buffer is forced to quit. All code should be run before
	// this method returns.
	Quit(*Context)

	Command(*Context, Command)
	Key(*Context, Key)
	Mouse(*Context, Mouse)
}
