package ivo

type KeyMap struct {
	nodes map[string]*keyMapNode // maps mode to node
}

type keyMapNode struct {
	nodes  map[string]*keyMapNode
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
		subnode, ok := node.nodes[k.hash()]
		if !ok {
			node.nodes[k.hash()] = newKeyMapNode()
			subnode = node.nodes[k.hash()]
		}
		node = subnode
	}
	node.action = action
}

func (km *KeyMap) Get(mode string, kk []Key) (func(*Context), bool, bool) {
	node := km.node(mode)
	for len(kk) > 0 {
		k := kk[0]
		kk = kk[1:]
		subnode, ok := node.nodes[k.hash()]
		if !ok {
			return nil, false, false
		}
		node = subnode
	}
	more := len(node.nodes) != 0
	if node.action == nil {
		return nil, more, false
	}
	return node.action, more, true
}

func (km *KeyMap) node(mode string) *keyMapNode {
	node, ok := km.nodes[mode]
	if !ok {
		km.nodes[mode] = newKeyMapNode()
		node = km.nodes[mode]
	}
	return node
}

func newKeyMapNode() *keyMapNode {
	return &keyMapNode{
		nodes: make(map[string]*keyMapNode),
	}
}
