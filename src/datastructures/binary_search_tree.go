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

func (b *BinarySearchTree) TraverseLevelOrder() []int {
	elements := []int{}
	nodesQueue := NodeQueue{}
	if b.root == nil {
		return elements
	}
	nodesQueue.Enqueue(b.root)
	b.LevelOrderTraverseFromNodeAndUpdateSlice(&elements, &nodesQueue)
	return elements
}

func (b *BinarySearchTree) LevelOrderTraverseFromNodeAndUpdateSlice(elementsTraversedSoFar *[]int, nodesQueue *NodeQueue) {
	if len(nodesQueue.dataqueue) == 0 {
		return
	}
	currentNode := nodesQueue.Dequeue()
	if currentNode.left != nil {
		nodesQueue.Enqueue(currentNode.left)
	}
	if currentNode.right != nil {
		nodesQueue.Enqueue(currentNode.right)
	}
	*elementsTraversedSoFar = append(*elementsTraversedSoFar, currentNode.data)
	b.LevelOrderTraverseFromNodeAndUpdateSlice(elementsTraversedSoFar, nodesQueue)
}

func (b *BinarySearchTree) Remove(element int) bool {
	if !b.Contains(element) {
		return false
	}
	b.removeElement(element, b.root, nil)

	return false
}

func (b *BinarySearchTree) removeElement(element int, currentNode *BSTNode, parentNode *BSTNode) {
	if element == currentNode.data {
		// actual remove code here.
		b.removeNodeFromTree(currentNode, parentNode)
	} else if element > currentNode.data {
		b.removeElement(element, currentNode.right, currentNode)
	} else {
		b.removeElement(element, currentNode.left, currentNode)
	}
}

func (b *BinarySearchTree) removeNodeFromTree(currentNode *BSTNode, parentNode *BSTNode) {

	if currentNode.left == nil {
		b.removeNodeFromParent(currentNode, parentNode, currentNode.right)
	} else if currentNode.right == nil {
		b.removeNodeFromParent(currentNode, parentNode, currentNode.left)
	} else {
		// find the largest element of the left sub tree and then swap
		largestNode, parentOfLargestNode := b.FindLargestNodeInTheLeftSubtree(currentNode.left, currentNode)
		parentOfLargestNode.right = nil
		if parentNode == nil {
			b.root = largestNode
			largestNode.left = currentNode.left
			largestNode.right = currentNode.right
			return
		}
		if parentNode.data > currentNode.data {
			parentNode.left = largestNode
		} else {
			parentNode.right = largestNode
		}
		largestNode.left = currentNode.left
		largestNode.right = currentNode.right

	}
}

func (b *BinarySearchTree) FindLargestNodeInTheLeftSubtree(currentNode *BSTNode, parentNode *BSTNode) (*BSTNode, *BSTNode) {
	if currentNode.right == nil {
		return currentNode, parentNode
	}
	return b.FindLargestNodeInTheLeftSubtree(currentNode.right, currentNode)
}

func (b *BinarySearchTree) removeNodeFromParent(currentNode *BSTNode, parentNode *BSTNode, replacementNode *BSTNode) {
	if parentNode == nil {
		b.root = replacementNode
		return
	}
	if parentNode.left == currentNode {
		parentNode.left = replacementNode
	} else {
		parentNode.right = replacementNode
	}
}

func (b *BinarySearchTree) FetchAllElementsLessThan(higherLimit int) []int {
	matchingElements := []int{}

	inorderTraversal := b.TraverseInOrder()
	for _, element := range inorderTraversal {
		if element >= higherLimit {
			continue
		}
		matchingElements = append(matchingElements, element)
	}
	return matchingElements
}

func (b *BinarySearchTree) FetchAllElementsGreaterThan(lowerLimit int) []int {
	matchingElements := []int{}

	inorderTraversal := b.TraverseInOrder()
	for _, element := range inorderTraversal {
		if element <= lowerLimit {
			continue
		}
		matchingElements = append(matchingElements, element)
	}
	return matchingElements
}

func (b *BinarySearchTree) FindElementsInRange(lowerLimit int, higherLimit int) []int {
	matchingElements := []int{}
	if b.root != nil {
		b.findElementsInRange(b.root, &matchingElements, lowerLimit, higherLimit)
	}
	return matchingElements
}

func (b *BinarySearchTree) findElementsInRange(currentRoot *BSTNode, matchingElements *[]int, lowerLimit int, higherLimit int) {
	if currentRoot == nil {
		return
	}
	if currentRoot.data > lowerLimit && currentRoot.data < higherLimit {
		*matchingElements = append(*matchingElements, currentRoot.data)
		b.findElementsInRange(currentRoot.left, matchingElements, lowerLimit, higherLimit)
		b.findElementsInRange(currentRoot.right, matchingElements, lowerLimit, higherLimit)
	} else if currentRoot.data > higherLimit {
		b.findElementsInRange(currentRoot.left, matchingElements, lowerLimit, higherLimit)
	} else if currentRoot.data < lowerLimit {
		b.findElementsInRange(currentRoot.right, matchingElements, lowerLimit, higherLimit)
	}
}
