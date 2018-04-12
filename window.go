package ivo

// Window represents the main content.
//
// In all Window's methods, a Context is passed. This Context may be used after the
// methods return. However, the Context should only be used for actions related to the
// method call is was passed under.
type Window interface {
	// Quit is called when the buffer is forced to quit. Cleanup code should be run before
	// this method returns.
	Quit(Context)

	// Command is called when an arbituary command is to be executed.
	Command(Context, Command)

	// Key is called when a key is pressed.
	Key(Context, Key)

	// Mouse is called when a mouse action is performed.
	Mouse(Context, Mouse)
}
