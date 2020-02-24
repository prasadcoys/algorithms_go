package datastructures

type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

type BinarySearchTree struct {
	root *BSTNode
	size int
}

func (b *BinarySearchTree) Add(node *BSTNode) {
	if b.root == nil {
		b.root = node
		b.size = b.size + 1
		return
	}
	b.addNode(node, b.root)

}

func (b *BinarySearchTree) addNode(node *BSTNode, currentRoot *BSTNode) {
	if node.data < currentRoot.data {
		if currentRoot.left == nil {
			currentRoot.left = node
			b.size = b.size + 1
		} else {
			b.addNode(node, currentRoot.left)
		}

	} else if node.data > currentRoot.data {
		if currentRoot.right == nil {
			currentRoot.right = node
			b.size = b.size + 1
		} else {
			b.addNode(node, currentRoot.right)
		}
	}
}

func (b *BinarySearchTree) Contains(value int) bool {
	if b.root == nil {
		return false
	} else {
		return b.containsNode(value, b.root)
	}
}

func (b *BinarySearchTree) containsNode(value int, currentRoot *BSTNode) bool {
	if currentRoot == nil {
		return false
	} else if value == currentRoot.data {
		return true
	} else if value < currentRoot.data {
		return b.containsNode(value, currentRoot.left)
	} else {
		return b.containsNode(value, currentRoot.right)
	}
}

func (b *BinarySearchTree) Size() int {
	return b.size
}
