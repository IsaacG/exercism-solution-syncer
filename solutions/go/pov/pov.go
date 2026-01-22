package pov

// Tree holds a tree structure.
type Tree struct {
	value  string
	parent *Tree
	// adjacencies = parent + children
	adjacencies []*Tree
}

// New creates and returns a new Tree with the given root value and children.
func New(value string, children ...*Tree) *Tree {
	t := &Tree{value, nil, children}
	for _, c := range children {
		c.adjacencies = append(c.adjacencies, t)
		c.parent = t
	}
	return t
}

// findNode searches downwards (ignoring the parent) for a node and returns the path and if the node exists.
func (tr *Tree) findNode(value string) ([]*Tree, bool) {
	if tr.Value() == value {
		return []*Tree{tr}, true
	}
	for _, child := range tr.Children() {
		if found, ok := child.findNode(value); ok {
			path := make([]*Tree, 1, len(found)+1)
			path[0] = tr
			return append(path, found...), true
		}
	}
	return nil, false
}

// setParent recursively updates the parent of a tree.
func (tr *Tree) setParent(parent *Tree) {
	tr.parent = parent
	for _, child := range tr.Children() {
		child.setParent(tr)
	}
}

// Value returns the value at the root of a tree.
func (tr *Tree) Value() string {
	return tr.value
}

// Children returns a slice containing the children of a tree.
func (tr *Tree) Children() []*Tree {
	var children []*Tree
	for _, node := range tr.adjacencies {
		if node != tr.parent {
			children = append(children, node)
		}
	}
	return children
}

// FromPov returns the pov from the node specified in the argument.
func (tr *Tree) FromPov(from string) *Tree {
	path, ok := tr.findNode(from)
	if !ok {
		return nil
	}
	root := path[len(path)-1]
	root.setParent(nil)
	return root
}

// PathTo returns the shortest path between two nodes in the tree.
// O(n^2) time.
func (tr *Tree) PathTo(from, to string) []string {
	// Validate the from and to exist in the tree.
	if _, ok := tr.findNode(from); !ok {
		return nil
	}
	if _, ok := tr.findNode(to); !ok {
		return nil
	}

	tr = tr.FromPov(from)
	path, ok := tr.findNode(to)
	if !ok {
		return nil
	}
	out := make([]string, len(path))
	for i, t := range path {
		out[i] = t.Value()
	}
	return out
}
