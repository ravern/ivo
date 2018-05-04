package key

import (
	"ivoeditor.com/ivo"
)

// Map is a map of keys to handlers, with support for modes.
//
// Map uses a tree of handlers internally to represent key combinations.
// Therefore, when requesting the handler for a certain combination of
// keys, Map will also return further information about the node.
type Map struct {
	modes map[string]*node
}

// node forms an handler tree with other nodes.
type node struct {
	children map[ivo.Key]*node
	handler  func(ivo.Context, []ivo.Key)
}

// NewMap creates an empty key map.
func NewMap() *Map {
	return &Map{
		modes: make(map[string]*node),
	}
}

// Set sets the handler for the given key combination, for the specified
// mode.
func (m *Map) Set(mode string, keys []ivo.Key, handler func(ivo.Context, []ivo.Key)) {
	node := m.mode(mode)

	// Keep trying to find child node, creating new node if not found.
	for _, key := range keys {
		child, ok := node.children[key]
		if !ok {
			child = newNode()
			node.children[key] = child
		}

		node = child
	}

	node.handler = handler
}

// SetFallback sets the fallback handler for the given mode (See Get).
func (m *Map) SetFallback(mode string, handler func(ivo.Context, []ivo.Key)) {
	node := m.mode(mode)
	node.handler = handler
}

// Get returns the handler for the given key combination, for the specified
// mode.
//
// The first bool is the more flag. It represents whether there are
// additional key combinations that start with the given keys. The
// second bool is the ok flag. It represents whether a handler exists
// for the given key combination.
//
// If the given mode has a fallback handler, and no handlers are found
// for the given key combination, the fallback handler will be returned,
// with the more flag set to false and the ok flag set to true.
func (m *Map) Get(mode string, keys []ivo.Key) (func(ivo.Context, []ivo.Key), bool, bool) {
	node, ok := m.modes[mode]
	if !ok {
		return nil, false, false
	}

	// Keep a reference to the fallback in case it's needed.
	fallback := node.handler

	// Keep trying to find child node until found or unavailable.
	for _, key := range keys {
		child, ok := node.children[key]
		if !ok {
			return fallback, false, true
		}

		node = child
	}

	return node.handler, len(node.children) != 0, true
}

// mode returns the node for the given mode, creating it if it
// doesn't exist.
func (m *Map) mode(mode string) *node {
	node, ok := m.modes[mode]

	if !ok {
		node = newNode()
		m.modes[mode] = node
	}

	return node
}

// newNode creates an empty node.
func newNode() *node {
	return &node{
		children: make(map[ivo.Key]*node),
	}
}
