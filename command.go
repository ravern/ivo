package ivo

// Command is an arbituary command sent to windows.
//
// Command is mainly used to communicate between windows and for user commands
// (e.g. through the command bar). Commands should be used only when necessary,
// since they inherently break the type system.
type Command struct {
	Name    string
	Payload map[string]interface{}
}
