package ivo

// Context contains important objects and methods for the buffer to use.
//
// For Buffers acting as a proxy to other Buffers, Context can and should
// be implemented for the purpose of that Buffer, passing along the appropriate
// information.
type Context interface {
	// Logger should be used to perform all logging. In the case of core Context,
	// it is the logger assigned to the Core object.
	Logger() Logger

	// Cells is the individual cells of the screen. These are usually not set
	// directly, instead using the `ivo/buffer` package to draw different
	// components.
	Cells() Cells

	// Render updates the screen with the contents of Cells.
	Render()

	// Quit signals the caller to quit. Cleanup code should not be run prior to
	// this, and instead should be run when the Quit(Context) method of the
	// Buffer is called.
	Quit()
}
