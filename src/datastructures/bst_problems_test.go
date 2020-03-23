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

func TestIfBSTHasDeadEnd(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 8})
	binaryTree.Add(&BSTNode{data: 7})
	binaryTree.Add(&BSTNode{data: 10})
	binaryTree.Add(&BSTNode{data: 2})
	binaryTree.Add(&BSTNode{data: 9})
	binaryTree.Add(&BSTNode{data: 13})
	assert.True(t, binaryTree.HasDeadEnd())
}

func TestIfBSTHasNoDeadEnd(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 8})
	binaryTree.Add(&BSTNode{data: 7})
	binaryTree.Add(&BSTNode{data: 9})
	assert.False(t, binaryTree.HasDeadEnd())
}

func TestIfPreorderTraversalIsABSTForValidBST(t *testing.T) {
	preorderEntries := []int{40, 30, 35, 80, 100}
	assert.True(t, IsPreorderAValidBST(preorderEntries))
}

func TestIfPreorderTraversalIsABSTForInvalidBST(t *testing.T) {
	preorderEntries := []int{7, 9, 6, 1, 4, 2, 3, 40}
	assert.False(t, IsPreorderAValidBST(preorderEntries))
}

func TestIfPreorderTraversalIsABSTForValidLargerBST(t *testing.T) {
	preorderEntries := []int{40, 30, 32, 35, 80, 90, 100, 120}
	assert.True(t, IsPreorderAValidBST(preorderEntries))
}

func TestIfPreorderTraversalIsABSTForInValidLargerBST(t *testing.T) {
	preorderEntries := []int{40, 30, 32, 35, 80, 90, 100, 120, 20}
	assert.False(t, IsPreorderAValidBST(preorderEntries))
}

func TestIfInorderSuccessorForInvalidValueIsMinusOne(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	assert.Equal(t, -1, binaryTree.GetInorderSuccesor(100))
}

func TestIfInorderSuccessorForValidValueIsFetchedCorrectly(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	assert.Equal(t, 13, binaryTree.GetInorderSuccesor(12))
}

func TestInorderSuccessorForALargeTree(t *testing.T) {
	values := []int{887, 778, 916, 794, 336, 387, 493, 650, 422, 363, 28, 691, 60, 764, 927, 541, 427, 173, 737, 212, 369, 568, 430, 783, 531, 863, 124, 68, 136, 930, 803, 23, 59, 70, 168, 394, 457, 12, 43, 230, 374, 422, 920, 785, 538, 199, 325, 316, 371, 414, 527, 92, 981, 957, 874, 863, 171, 997, 282, 306, 926, 85, 328, 337, 506, 847, 730, 314, 858, 125, 896, 583, 546, 815, 368, 435, 365, 44, 751, 88, 809, 277, 179, 789}
	binaryTree := createTreeWithSlice(values)
	assert.Equal(t, 927, binaryTree.GetInorderSuccesor(926))
}

func TestIfWeCanCorrectlyCalculatePairsViolatingTheBSTProperty(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.root = &BSTNode{data: 50}
	binaryTree.root.left = &BSTNode{data: 30}
	binaryTree.root.right = &BSTNode{data: 60}
	binaryTree.root.left.left = &BSTNode{data: 20}
	binaryTree.root.left.right = &BSTNode{data: 25}
	binaryTree.root.right.left = &BSTNode{data: 10}
	binaryTree.root.right.right = &BSTNode{data: 40}
	assert.Equal(t, 7, binaryTree.FindNumberOfOffendingPairs())
}

func TestIfNumberOfInversionsInMergeSortIsCalculatedCorrectly(t *testing.T) {
	entries := []int{20, 30, 25, 50, 10, 60, 40}
	assert.Equal(t, 7, findNumberOfInversions(entries))
}

func TestIfBinarySearchTreeCanContainAPairWithGivenSum(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	assert.Equal(t, true, DoesPairMakingGivenSumExistsInTree(binaryTree, 27))
}

func TestIfBinarySearchTreeDoesNotContainAPairWithGivenSum(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	assert.Equal(t, false, DoesPairMakingGivenSumExistsInTree(binaryTree, 35))
}
func createTreeWithSlice(values []int) BinarySearchTree {
	binaryTree := BinarySearchTree{}
	for _, value := range values {
		binaryTree.Add(&BSTNode{data: value})
	}
	return binaryTree
}
