package key

import "ivoeditor.com/ivo"

// Map is a map of keys to actions, with support for modes.
type Map struct {
	modes map[string]*node
}

// node forms an action tree with other nodes.
type node struct {
	children map[ivo.Key]*node
	action   func(ivo.Context)
}

// NewMap creates an empty map.
func NewMap() *Map {
	return &Map{
		modes: make(map[string]*node),
	}
}

// Set sets the action for the given keys, for the specified mode.
func (m *Map) Set(mode string, kk []ivo.Key, action func(ivo.Context)) {
	node := m.mode(mode)
	for _, k := range kk {
		child, ok := node.children[k]
		if !ok {
			child = newNode()
			node.children[k] = child
		}
		node = child
	}
	node.action = action
}

// Get returns the action for the given keys, for the specified mode.
func (m *Map) Get(mode string, kk []ivo.Key) (func(ivo.Context), bool, bool) {
	node, ok := m.modes[mode]
	if !ok {
		return nil, false, false
	}
	for _, k := range kk {
		child, ok := node.children[k]
		if !ok {
			return nil, false, false
		}
		node = child
	}
	return node.action, len(node.children) != 0, true
}

// mode returns the node for the specified mode, creating it if it
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
