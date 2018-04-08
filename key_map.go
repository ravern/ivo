package ivo

type KeyMap struct {
	root *keyMapNode
}

type keyMapNode struct {
	action   func()
	children map[string]*keyMapNode
}
