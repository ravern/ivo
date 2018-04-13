package ivo

// Window represents the main content.
//
// In the methods where a Context is passed, the Context may be used after the methods return.
// However, the Context should only be used for actions related to the method call is was
// passed via.
//
// For example, for in a call to Key, only actions directly resulting from the key (e.g.
// autocompletion or syntax highlighting) should have access to that Context. Actions that
// result from a call to Mouse in the future should not have access to this Context.
//
// An easy way to understand this is that the Context should not be stored directly in the Window
// or its descendants. If the Context is to be used after the method it was passed via returns,
// then it should be captured in closures or goroutines.
type Window interface {
	// Close is called to run any cleanup code the Window has before the main loop is
	// exited. All code should be run before this method returns.
	Close(Context)

	// Command is called when an arbituary command is to be executed.
	Command(Context, Command)

	// Key is called when a key is pressed.
	Key(Context, Key)

	// Mouse is called when a mouse action is performed.
	Mouse(Context, Mouse)
}
