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

func (n *NodeQueue) Push(node *BSTNode) {
	n.Enqueue(node)
}

func (n *NodeQueue) Pop() *BSTNode {
	length := len(n.dataqueue)
	tailNode := n.dataqueue[length-1]
	n.dataqueue = n.dataqueue[:length-1]
	return tailNode
}
