package key

import "ivoeditor.com/ivo"

type Map struct {
	modes map[string]*node
}

type node struct {
	children map[ivo.Key]*node
	action   func(ivo.Context)
}

func NewMap() *Map {
	return &Map{
		modes: make(map[string]*node),
	}
}

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

func (m *Map) mode(mode string) *node {
	node, ok := m.modes[mode]
	if !ok {
		node = newNode()
		m.modes[mode] = node
	}
	return node
}

func newNode() *node {
	return &node{
		children: make(map[ivo.Key]*node),
	}
}
