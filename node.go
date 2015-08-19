package linkedList

type Node struct {
	next, prev *Node
	Value interface{}
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) NewNode(value interface{}) *Node{
	var new *Node
	new = &Node{Value: value}
	return new
}
