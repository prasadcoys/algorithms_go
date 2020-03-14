package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfWeCanGetCountOfAllElementsLessThanACertainNumber(t *testing.T) {
	higherLimit := 8
	binaryTree := CreateBigBinaryTree()
	assert.ElementsMatch(t, []int{1, 3, 5, 6}, binaryTree.FetchAllElementsLessThan(higherLimit))
}

func TestIfWeCanGetCountOfAllElementsGreaterThanACertainNumber(t *testing.T) {
	lowerLimit := 8
	binaryTree := CreateBigBinaryTree()
	assert.ElementsMatch(t, []int{11, 12, 13, 14, 15, 17, 19}, binaryTree.FetchAllElementsGreaterThan(lowerLimit))
}

func TestIfWeCanPrintTheElementsInRange(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	lowerLimit := 8
	upperLimit := 17
	expectedMatches := []int{11, 12, 13, 14, 15}
	assert.ElementsMatch(t, expectedMatches, binaryTree.FindElementsInRange(lowerLimit, upperLimit))
}

func TestIfWeCanConvertLevelOrderTraversalToBST(t *testing.T) {
	levelOrder := []int{7, 4, 12}
	binaryTree := ConvertLevelOrderTraversalToBinarySearchTree(levelOrder)
	preorderEntries := []int{4, 12, 7}
	assert.Equal(t, preorderEntries, binaryTree.TraversePostOrder())
}

func TestIfWeCanConvertABiggerLevelOrderTraversalToBST(t *testing.T) {
	levelOrder := []int{7, 4, 12, 3, 6, 8, 1, 5, 10}
	binaryTree := ConvertLevelOrderTraversalToBinarySearchTree(levelOrder)
	preorderEntries := []int{1, 3, 5, 6, 4, 10, 8, 12, 7}
	assert.Equal(t, preorderEntries, binaryTree.TraversePostOrder())
}

func TestIFWeCanCalculateTheSumOfAllLeafNodesInABST(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	assert.Equal(t, 59, binaryTree.GetSumOfAllLeafNodes())

}

func TestIFWeCanCalculateTheSumOfAllLeafNodesInABSTIfRootIsZero(t *testing.T) {
	binaryTree := BinarySearchTree{}
	assert.Equal(t, 0, binaryTree.GetSumOfAllLeafNodes())

}

func TestIfWeCanBalanceABInarySearchTreeWithThreeElementsInAscendingOrder(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 1})
	binaryTree.Add(&BSTNode{data: 2})
	binaryTree.Add(&BSTNode{data: 3})
	assert.Equal(t, 3, binaryTree.Height())
	binaryTree.BalanaceBinaryTree()
	assert.Equal(t, 2, binaryTree.Height())

}

func TestIfWeCanBalanceABInarySearchTreeWithThreeElementsInDescendingOrder(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.Add(&BSTNode{data: 2})
	binaryTree.Add(&BSTNode{data: 1})
	assert.Equal(t, 3, binaryTree.Height())
	binaryTree.BalanaceBinaryTree()
	assert.Equal(t, 2, binaryTree.Height())

}

func TestIfWeCanBalanceABigTree(t *testing.T) {
	//TODO: balancing a bst doesn't work. the handling thus far is very rudimentary and will fail in bigger cases
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 1})
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.Add(&BSTNode{data: 5})
	binaryTree.Add(&BSTNode{data: 6})
	binaryTree.Add(&BSTNode{data: 8})
	binaryTree.Add(&BSTNode{data: 11})
	binaryTree.Add(&BSTNode{data: 12})
	binaryTree.Add(&BSTNode{data: 13})
	binaryTree.Add(&BSTNode{data: 14})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 17})
	binaryTree.Add(&BSTNode{data: 19})
	//Balancing a binarys search tree will be handled after finishing the BBST chapter in William Fiset's series
}

func TestIfMinusOneIsReturnedAsSmallestElementOfAnEmptyBST(t *testing.T) {
	binaryTree := BinarySearchTree{}
	assert.Equal(t, -1, binaryTree.GetSmallestElement())
}

func TestIfTheLoneValueIsReturnedIfTreeHasOnlyOneElement(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 2})
	assert.Equal(t, 2, binaryTree.GetSmallestElement())
}

func TestIfWeGetLowestValueIfWeHaveMoreThanOneElement(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 2})
	binaryTree.Add(&BSTNode{data: 1})
	binaryTree.Add(&BSTNode{data: 3})
	assert.Equal(t, 1, binaryTree.GetSmallestElement())
}

func TestIfWeGetTheLowestValueIfWeHaveToTraverseMultipleLevels(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	assert.Equal(t, 1, binaryTree.GetSmallestElement())
}

func TestIfWeCanGetTheLowestCommonAncestorWhereBothElementsAreInTheSameSide(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	assert.Equal(t, 6, binaryTree.CalculateLowestCommonAncestor(3, 8))
}

func TestIfWeCanGetTheLowestCommonAncestorWhenTheElementsAreInDifferentSides(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 5})
	binaryTree.Add(&BSTNode{data: 4})
	binaryTree.Add(&BSTNode{data: 6})
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.Add(&BSTNode{data: 7})
	binaryTree.Add(&BSTNode{data: 8})
	assert.Equal(t, 7, binaryTree.CalculateLowestCommonAncestor(7, 8))
}

func TestIfWeCanDeleteNodesGreaterThanACertainNumberIfNumberIsOnRightTreeCompletely(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 20})
	binaryTree.Add(&BSTNode{data: 8})
	binaryTree.Add(&BSTNode{data: 22})
	binaryTree.Add(&BSTNode{data: 4})
	binaryTree.Add(&BSTNode{data: 12})
	binaryTree.Add(&BSTNode{data: 10})
	binaryTree.Add(&BSTNode{data: 14})
	binaryTree.DeleteNodesGreaterThan(20)
	inorderTraversals := []int{4, 8, 10, 12, 14, 20}
	assert.Equal(t, inorderTraversals, binaryTree.TraverseInOrder())
}

func TestIfWeCanDeleteNodesGreaterThanACertainNumberIfNumberIsOnLeftTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 20})
	binaryTree.Add(&BSTNode{data: 8})
	binaryTree.Add(&BSTNode{data: 22})
	binaryTree.Add(&BSTNode{data: 4})
	binaryTree.Add(&BSTNode{data: 12})
	binaryTree.Add(&BSTNode{data: 10})
	binaryTree.Add(&BSTNode{data: 14})
	binaryTree.DeleteNodesGreaterThan(8)
	inorderTraversals := []int{4, 8}
	assert.Equal(t, inorderTraversals, binaryTree.TraverseInOrder())
}

func TestIfWeCanDeleteNodesGreaterThanACertainNumberIfNumberIsNotPresentInTheTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 20})
	binaryTree.Add(&BSTNode{data: 8})
	binaryTree.Add(&BSTNode{data: 22})
	binaryTree.Add(&BSTNode{data: 4})
	binaryTree.Add(&BSTNode{data: 12})
	binaryTree.Add(&BSTNode{data: 10})
	binaryTree.Add(&BSTNode{data: 14})
	binaryTree.DeleteNodesGreaterThan(9)
	inorderTraversals := []int{4, 8}
	assert.Equal(t, inorderTraversals, binaryTree.TraverseInOrder())
}

func TestIfWeCanAddAllGreaterValuesToEveryNodeInBST(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 2})
	binaryTree.Add(&BSTNode{data: 1})
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.AddAllGreaterValuesToEveryNode()
	inorderTraversals := []int{6, 5, 3}
	assert.Equal(t, inorderTraversals, binaryTree.TraverseInOrder())
}

func TestIfWeCanAddAllGreaterValuesToEveryNodeInABiggerBST(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 50})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 70})
	binaryTree.Add(&BSTNode{data: 20})
	binaryTree.Add(&BSTNode{data: 40})
	binaryTree.Add(&BSTNode{data: 60})
	binaryTree.Add(&BSTNode{data: 80})
	binaryTree.AddAllGreaterValuesToEveryNode()
	inorderTraversals := []int{350, 330, 300, 260, 210, 150, 80}
	assert.Equal(t, inorderTraversals, binaryTree.TraverseInOrder())
	levelOrderTraversals := []int{260, 330, 150, 350, 300, 210, 80}
	assert.Equal(t, levelOrderTraversals, binaryTree.TraverseLevelOrder())
}

func TestIfBSTFailsForTreeNotSatisfyingTheBSTProperty(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	binaryTree.root.right.left.data = 10
	assert.False(t, binaryTree.IsBST())
}

func TestIfBSTSucceedsForTreeSatisfyingTheBSTProperty(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	assert.True(t, binaryTree.IsBST())
}
