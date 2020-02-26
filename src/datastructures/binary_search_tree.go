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

func (b *BinarySearchTree) Height() int {
	return calculateHeight(b.root)
}

func calculateHeight(rootNode *BSTNode) int {
	if rootNode == nil {
		return 0
	} else {
		leftSubtreeHeight := calculateHeight(rootNode.left)
		rightSubtreeHeight := calculateHeight(rootNode.right)
		if leftSubtreeHeight >= rightSubtreeHeight {
			return 1 + leftSubtreeHeight
		} else {
			return 1 + rightSubtreeHeight
		}
	}

}

func (b *BinarySearchTree) TraversePreOrder() []int {
	elements := []int{}
	b.PreOrderTraverseFromNodeAndUpdateSlice(b.root, &elements)
	return elements
}

func (b *BinarySearchTree) PreOrderTraverseFromNodeAndUpdateSlice(currentRoot *BSTNode, elementsTraversedSoFar *[]int) {
	if currentRoot == nil {
	} else {
		*elementsTraversedSoFar = append(*elementsTraversedSoFar, currentRoot.data)
		b.PreOrderTraverseFromNodeAndUpdateSlice(currentRoot.left, elementsTraversedSoFar)
		b.PreOrderTraverseFromNodeAndUpdateSlice(currentRoot.right, elementsTraversedSoFar)
	}
}

func (b *BinarySearchTree) TraverseInOrder() []int {
	elements := []int{}
	b.InOrderTraverseFromNodeAndUpdateSlice(b.root, &elements)
	return elements
}

func (b *BinarySearchTree) InOrderTraverseFromNodeAndUpdateSlice(currentRoot *BSTNode, elementsTraversedSoFar *[]int) {
	if currentRoot == nil {
	} else {
		b.InOrderTraverseFromNodeAndUpdateSlice(currentRoot.left, elementsTraversedSoFar)
		*elementsTraversedSoFar = append(*elementsTraversedSoFar, currentRoot.data)
		b.InOrderTraverseFromNodeAndUpdateSlice(currentRoot.right, elementsTraversedSoFar)
	}
}

func (b *BinarySearchTree) TraversePostOrder() []int {
	elements := []int{}
	b.PostOrderTraverseFromNodeAndUpdateSlice(b.root, &elements)
	return elements
}

func (b *BinarySearchTree) PostOrderTraverseFromNodeAndUpdateSlice(currentRoot *BSTNode, elementsTraversedSoFar *[]int) {
	if currentRoot == nil {

	} else {
		b.PostOrderTraverseFromNodeAndUpdateSlice(currentRoot.left, elementsTraversedSoFar)
		b.PostOrderTraverseFromNodeAndUpdateSlice(currentRoot.right, elementsTraversedSoFar)
		*elementsTraversedSoFar = append(*elementsTraversedSoFar, currentRoot.data)

	}

}
