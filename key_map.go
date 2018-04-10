package ivo

type KeyMap struct {
	nodes map[string]*keyMapNode // maps mode to node
}

// keyMapNode represents a node in a KeyMap (which is actually a tree).
type keyMapNode struct {
	nodes  map[Key]*keyMapNode
	action func(*Context)
}

func NewKeyMap() *KeyMap {
	return &KeyMap{
		nodes: make(map[string]*keyMapNode),
	}
}

func (km *KeyMap) Set(mode string, kk []Key, action func(*Context)) {
	node := km.node(mode)
	for len(kk) > 0 {
		k := kk[0]
		kk = kk[1:]
		subnode, ok := node.nodes[k]
		if !ok {
			node.nodes[k] = newKeyMapNode()
			subnode = node.nodes[k]
		}
		node = subnode
	}
	node.action = action
}

func (km *KeyMap) SetFallback(mode string, action func(*Context)) {
	node := km.node(mode)
	node.action = action
}

func (km *KeyMap) Get(mode string, kk []Key) (func(*Context), bool, bool) {
	node := km.node(mode)
	for len(kk) > 0 {
		k := kk[0]
		kk = kk[1:]
		subnode, ok := node.nodes[k]
		if !ok {
			node := km.node(mode)
			if node.action == nil {
				return nil, false, false
			}
			return node.action, false, true
		}
		node = subnode
	}
	more := len(node.nodes) != 0
	if node.action == nil {
		return nil, more, false
	}
	return node.action, more, true
}

// node returns the corresponding keyMapNode for that mode. It creates a new keyMapNode for
// that mode if it doesn't exist.
func (km *KeyMap) node(mode string) *keyMapNode {
	node, ok := km.nodes[mode]
	if !ok {
		km.nodes[mode] = newKeyMapNode()
		node = km.nodes[mode]
	}
	return node
}

// newKeyMapNode creates an empty keyMapNode.
func newKeyMapNode() *keyMapNode {
	return &keyMapNode{
		nodes: make(map[Key]*keyMapNode),
	}
}
