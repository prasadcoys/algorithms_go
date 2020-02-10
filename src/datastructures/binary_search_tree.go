package datastructures

type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

type BinarySearchTree struct {
	root *BSTNode
}

func (b *BinarySearchTree) Add(node *BSTNode) {
	if b.root == nil {
		b.root = node
	}
}
