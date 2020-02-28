package datastructures

type NodeQueue struct {
	dataqueue []*BSTNode
}

func (n *NodeQueue) Enqueue(node *BSTNode) {
	n.dataqueue = append(n.dataqueue, node)
}

func (n *NodeQueue) Dequeue() *BSTNode {
	headNode := n.dataqueue[0]
	n.dataqueue = n.dataqueue[1:]
	return headNode
}
