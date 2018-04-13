package ivo

// Context contains important objects and methods for the window
// to use.
//
// When a Context is passed to a Window via its methods, some of
// the methods of the previous Context that was passed will no
// longer work. This includes Buffer and Render.
//
// For Windows acting as a proxy to other Windows, Context can
// and should be implemented for the purpose of that Window,
// passing along the appropriate information.
type Context interface {
	// Logger should be used to perform all logging.
	//
	// In the case of core Context, it is the logger assigned to
	// the Core object.
	Logger() Logger

	// Quit signals the caller to quit.
	//
	// Cleanup code should not be run prior to this, and instead
	// should be run when the Close method of the Window is called.
	Quit()

	// Command sends an arbituary command to the caller.
	//
	// This is often used more for any custom Window proxies, for
	// example to broadcast a command to other Windows.
	Command(*Command)

	// Buffer is the buffer holding the cells of the screen.
	//
	// These are usually not set directly, instead using the
	// `ivo/window` package to draw different components.
	Buffer() Buffer

	// Render updates the screen with the contents of Buffer.
	Render()
}
