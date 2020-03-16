package datastructures

import (
	"strconv"
)

type BSTNode struct {
	data  int
	left  *BSTNode
	right *BSTNode
}

func (b *BSTNode) String() string {
	return strconv.Itoa(b.data)
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

func (b *BinarySearchTree) TraverseReverseInOrder() []int {
	elements := []int{}
	b.reverseInOrderTraverseFromNodeAndUpdateSlice(b.root, &elements)
	return elements
}

func (b *BinarySearchTree) reverseInOrderTraverseFromNodeAndUpdateSlice(currentRoot *BSTNode, elementsTraversedSoFar *[]int) {

	if currentRoot == nil {
	} else {
		b.reverseInOrderTraverseFromNodeAndUpdateSlice(currentRoot.right, elementsTraversedSoFar)
		*elementsTraversedSoFar = append(*elementsTraversedSoFar, currentRoot.data)
		b.reverseInOrderTraverseFromNodeAndUpdateSlice(currentRoot.left, elementsTraversedSoFar)
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

func ConvertLevelOrderTraversalToBinarySearchTree(levelOrder []int) BinarySearchTree {
	binaryTree := BinarySearchTree{}
	for _, element := range levelOrder {
		binaryTree.Add(&BSTNode{data: element})
	}
	return binaryTree
}

func (b *BinarySearchTree) GetSumOfAllLeafNodes() int {
	var sum int

	if b.root == nil {
		return sum
	}
	b.calculateSum(b.root, &sum)

	return sum
}

func (b *BinarySearchTree) calculateSum(currentRoot *BSTNode, sum *int) {
	if currentRoot == nil {
		return
	}
	if currentRoot.left == nil && currentRoot.right == nil {
		*sum = *sum + currentRoot.data
		return
	}
	b.calculateSum(currentRoot.left, sum)
	b.calculateSum(currentRoot.right, sum)

}

func (b *BinarySearchTree) BalanaceBinaryTree() {

	b.balanceSubTree(nil, b.root)
}

func (b *BinarySearchTree) balanceSubTree(parentNode *BSTNode, subtreeRoot *BSTNode) {
	numberOfNodesToSkip := b.Height() / 2
	currentNode := subtreeRoot
	nodesSkippedInStack := NodeQueue{}
	for nodesSkipped := 0; nodesSkipped < numberOfNodesToSkip; {
		nodesSkippedInStack.Push(currentNode)
		if currentNode.right != nil {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
		nodesSkipped = nodesSkipped + 1

	}
	//check if the subtree root is the root of the overall tree
	if parentNode == nil {
		b.root = currentNode

	}
	currentRoot := b.root
	for {
		if len(nodesSkippedInStack.dataqueue) == 0 {
			return
		}
		nodeFromStack := nodesSkippedInStack.Pop()
		nodeFromStack.left = nil
		nodeFromStack.right = nil
		if nodeFromStack.data < currentRoot.data {
			currentRoot.left = nodeFromStack
			currentRoot = currentRoot.left
		} else {
			currentRoot.right = nodeFromStack
			currentRoot = currentRoot.right
		}
	}
}

func (b *BinarySearchTree) GetSmallestElement() int {
	if b.root == nil {
		return -1
	}
	if b.root.left == nil {
		return b.root.data
	} else {
		return getSmallestElementOfSubTree(b.root.left)
	}

}

func getSmallestElementOfSubTree(subtreeRoot *BSTNode) int {
	if subtreeRoot.left == nil {
		return subtreeRoot.data
	}
	return getSmallestElementOfSubTree(subtreeRoot.left)
}
func (b *BinarySearchTree) CalculateLowestCommonAncestor(lowerLimit int, higherLimit int) int {
	currentNode := b.root
	for {
		if currentNode.data >= lowerLimit && currentNode.data <= higherLimit {
			return currentNode.data
		} else if currentNode.data <= lowerLimit && currentNode.data <= higherLimit {
			currentNode = currentNode.right
			continue
		} else if currentNode.data >= lowerLimit && currentNode.data > higherLimit {
			currentNode = currentNode.left
			continue
		}
	}
	return -1
}
func (b *BinarySearchTree) DeleteNodesGreaterThan(element int) {
	b.deleteElementsGreaterThan(element, nil, b.root)
}

func (b *BinarySearchTree) deleteElementsGreaterThan(element int, prevNode *BSTNode, currentNode *BSTNode) {
	if currentNode == nil {
		return
	}
	if currentNode.data <= element {
		b.deleteElementsGreaterThan(element, currentNode, currentNode.right)
	} else {
		if prevNode == nil {
			b.root = currentNode.left
			b.deleteElementsGreaterThan(element, prevNode, currentNode.left)
		} else {
			if prevNode.data < currentNode.data {
				prevNode.right = currentNode.left
				b.deleteElementsGreaterThan(element, prevNode, currentNode.left)
			} else {
				prevNode.left = currentNode.left
				b.deleteElementsGreaterThan(element, prevNode, currentNode.left)
			}
		}
	}

}

func (b *BinarySearchTree) AddAllGreaterValuesToEveryNode() {
	inorderTraversal := b.traverseInOrderAndAddNodeToStack()
	currentSum := 0
	for {
		{
			node := inorderTraversal.Pop()
			currentSum = currentSum + node.data
			node.data = currentSum
		}
		if len(inorderTraversal.dataqueue) <= 0 {
			return
		}
	}

}

//TODO: IF WE ARE NOT TO USE A STACK, ALL WE NEED TO DO IS DO A REVERSE IN ORDER TRAVERSAL AND WE WOULD BE ABLE TO DO THAT.
func (b *BinarySearchTree) traverseInOrderAndAddNodeToStack() NodeQueue {
	elements := NodeQueue{}
	b.inOrderTraverseFromNodeAndUpdateStack(b.root, &elements)
	return elements
}

func (b *BinarySearchTree) inOrderTraverseFromNodeAndUpdateStack(currentRoot *BSTNode, elementsTraversedSoFar *NodeQueue) {
	if currentRoot == nil {
	} else {
		b.inOrderTraverseFromNodeAndUpdateStack(currentRoot.left, elementsTraversedSoFar)
		// *elementsTraversedSoFar = append(*elementsTraversedSoFar, currentRoot.data)
		elementsTraversedSoFar.Push(currentRoot)
		b.inOrderTraverseFromNodeAndUpdateStack(currentRoot.right, elementsTraversedSoFar)
	}
}

func (b *BinarySearchTree) IsBST() bool {
	inorderTraversal := b.TraverseInOrder()
	for index, element := range inorderTraversal {
		if index > 1 && element < inorderTraversal[index-1] {
			return false
		}
	}
	return true
}

func (b *BinarySearchTree) HasDeadEnd() bool {
	inorderEntries := b.traverseInOrderAndAddNodeToStack()
	for i := 0; i < len(inorderEntries.dataqueue); i++ {
		currentNode := inorderEntries.dataqueue[i]
		if currentNode.left != nil && currentNode.right != nil {
			continue
		}
		if i == 0 {
			if currentNode.data < 2 {
				return true
			}
		} else if i == len(inorderEntries.dataqueue)-1 {
			continue
		} else {
			if (currentNode.data-inorderEntries.dataqueue[i-1].data) < 2 && (inorderEntries.dataqueue[i+1].data-currentNode.data) < 2 {
				return true
			}
		}

	}
	return false
}

func IsPreorderAValidBST(preorder []int) bool {
	isRootFinished, isRightStarted := false, false
	binaryTree := BinarySearchTree{}
	for _, value := range preorder {
		if !isRootFinished {
			binaryTree.root = &BSTNode{data: value}
			isRootFinished = true
			continue
		}
		if value < binaryTree.root.data && isRightStarted {
			return false
		} else {
			if value > binaryTree.root.data {
				isRightStarted = true
			}
			binaryTree.Add(&BSTNode{data: value})
		}
	}

	return true

}
