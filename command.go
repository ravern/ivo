package ivo

// Command is an arbituary command sent to buffers.
type Command struct {
	Name    string
	Payload map[string]interface{}
}
