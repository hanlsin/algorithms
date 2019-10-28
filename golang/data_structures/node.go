package data_structures

type Node struct {
	object interface{}	// generic type object
}

func (n *Node) Value() interface{} {
	return n.object
}
