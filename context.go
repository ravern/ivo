package ivo

// Context contains important objects and methods for the window to use.
//
// For Windows acting as a proxy to other Windows, Context can and should
// be implemented for the purpose of that Window, passing along the appropriate
// information.
type Context interface {
	// Logger should be used to perform all logging. In the case of core Context,
	// it is the logger assigned to the Core object.
	Logger() Logger

	// Buffer is the buffer holding the cells of the screen. These are usually not
	// set directly, instead using the `ivo/window` package to draw different
	// components.
	Buffer() Buffer

	// Render updates the screen with the contents of Buffer.
	Render()

	// Quit signals the caller to quit. Cleanup code should not be run prior to
	// this, and instead should be run when the Quit(Context) method of the
	// Buffer is called.
	Quit()
}
