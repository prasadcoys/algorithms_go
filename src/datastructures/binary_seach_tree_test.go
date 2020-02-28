package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateBigBinaryTree() BinarySearchTree {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 11})
	binaryTree.Add(&BSTNode{data: 6})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.Add(&BSTNode{data: 8})
	binaryTree.Add(&BSTNode{data: 13})
	binaryTree.Add(&BSTNode{data: 17})
	binaryTree.Add(&BSTNode{data: 1})
	binaryTree.Add(&BSTNode{data: 5})
	binaryTree.Add(&BSTNode{data: 12})
	binaryTree.Add(&BSTNode{data: 14})
	binaryTree.Add(&BSTNode{data: 19})
	return binaryTree
}

func TestIfBinarySearchTreeNodeCanBeCreated(t *testing.T) {
	node := BSTNode{data: 23}
	assert.Equal(t, 23, node.data)
	if node.left != nil {
		t.Fail()
	}
	if node.right != nil {
		t.Fail()
	}
}

func TestIfBinarySearchTreeCanBeCreated(t *testing.T) {
	binaryTree := BinarySearchTree{}
	if binaryTree.root != nil {
		t.Fail()
	}
}

func TestIfNodeCanBeAddedToAnEmptyBinarySearchTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	if binaryTree.root != nil {
		t.Fail()
	}
	binaryTree.Add(&BSTNode{data: 23})
	if binaryTree.root.data != 23 {
		t.Fail()
	}
}

func TestIfNodeCanBeAddedToRootAnNonEmptyBinarySearchTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 15})

	if binaryTree.root.right.data != 30 {
		t.Fail()
	}
	if binaryTree.root.left.data != 15 {
		t.Fail()
	}
	if binaryTree.root.data != 23 {
		t.Fail()
	}
}

func TestIfNodeCanBeAddedToTheLeftOfALeftSubTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 12})
	if binaryTree.root.left.left.data != 12 {
		t.Fail()
	}
}

func TestIfNodeCanBeAddedToTheRightOfARightSubTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 35})
	if binaryTree.root.right.right.data != 35 {
		t.Fail()
	}
}

func TestIfNodeCanBeAddedToTheRightOfALeftSubTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 18})
	if binaryTree.root.left.right.data != 18 {
		t.Fail()
	}
}

func TestIfNodeCanBeAddedToTheLeftOfARightSubTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 25})
	if binaryTree.root.right.left.data != 25 {
		t.Fail()
	}
}

func TestIfTreeContainsNodeWorksForAPresentValue(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 25})
	if binaryTree.Contains(15) != true {
		t.Fail()
	}
}

func TestIfTreeContainsNodeWorksForAnAbsentValue(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 25})
	if binaryTree.Contains(16) != false {
		t.Fail()
	}
}

func TestIfTreeContainsNodeWorksForAValuePresentInRightSubtree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 25})
	if binaryTree.Contains(30) != true {
		t.Fail()
	}
}

func TestIfTreeContainsNodeWorksForAValueAbsentInRightSubtree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 25})
	if binaryTree.Contains(50) != false {
		t.Fail()
	}
}

func TestIfTreeContainsReturnsFalseOnEmptyTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	if binaryTree.Contains(25) != false {
		t.Fail()
	}
}

func TestIfSizeOfEmptyBinarySearchTreeReturnsZero(t *testing.T) {
	binaryTree := BinarySearchTree{}
	if binaryTree.Size() != 0 {
		t.Fail()
	}
}

func TestIfSizeOfNonEmptyBinarySearchTreeReturnsCorrectSize(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 5})
	binaryTree.Add(&BSTNode{data: 2})
	binaryTree.Add(&BSTNode{data: 7})
	if binaryTree.Size() != 3 {
		t.Fail()
	}
}

func TestIfHeightOfABinaryTreeWithOneNodeIsCalculatedCorrectly(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 5})
	if binaryTree.Height() != 1 {
		t.Fail()
	}
}

func TestIfHeightOfATreeWithOneRootAndALeftLeafIsCalculatedCorrectly(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 5})
	binaryTree.Add(&BSTNode{data: 2})
	if binaryTree.Height() != 2 {
		t.Fail()
	}
}

func TestIfHeightOfABInaryTreeWithRootAndOneLeftAndOneRightLeaf(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 5})
	binaryTree.Add(&BSTNode{data: 9})
	binaryTree.Add(&BSTNode{data: 2})
	if binaryTree.Height() != 2 {
		t.Fail()
	}
}

func TestIfHeightOfATreeWithOneRootAndARightLeafIsCalculatedCorrectly(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 5})
	binaryTree.Add(&BSTNode{data: 9})
	if binaryTree.Height() != 2 {
		t.Fail()
	}
}

func TestIfHeightOFABalancedTreeIsCalculatedCorrectly(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	if binaryTree.Height() != 4 {
		t.Fail()
	}

}

func TestIfHeighOfATreeWithALeftHeavySubtreeIsCalculatedCorrectly(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 11})
	binaryTree.Add(&BSTNode{data: 6})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.Add(&BSTNode{data: 8})
	if binaryTree.Height() != 3 {
		t.Fail()
	}
}

func TestIfHeighOfATreeWithARightHeavySubtreeIsCalculatedCorrectly(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 11})
	binaryTree.Add(&BSTNode{data: 6})
	binaryTree.Add(&BSTNode{data: 15})
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.Add(&BSTNode{data: 8})
	binaryTree.Add(&BSTNode{data: 13})
	binaryTree.Add(&BSTNode{data: 17})
	binaryTree.Add(&BSTNode{data: 12})
	binaryTree.Add(&BSTNode{data: 14})
	binaryTree.Add(&BSTNode{data: 19})
	if binaryTree.Height() != 4 {
		t.Fail()
	}
}

func TestIfPreOrderTraversalOfARootWorksCorrectly(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 1})
	traversalEntries := []int{1}
	assert.Equal(t, traversalEntries, binaryTree.TraversePreOrder())
}

func TestIfInOrderTraversalOfARootWithOneLevelCHildrenWorksCorrectly(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.Add(&BSTNode{data: 1})
	binaryTree.Add(&BSTNode{data: 5})
	traversalEntries := []int{3, 1, 5}
	assert.Equal(t, traversalEntries, binaryTree.TraversePreOrder())
}

func TestIfPreOrderTraversalOfARootWithMultipleLevelsOfChildreWorksCorrectly(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	traversalEntries := []int{11, 6, 3, 1, 5, 8, 15, 13, 12, 14, 17, 19}
	assert.Equal(t, traversalEntries, binaryTree.TraversePreOrder())
}

func TestIfPreOrderTraversalForAnEmptyTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	traversalEntries := []int{}
	assert.Equal(t, traversalEntries, binaryTree.TraversePreOrder())
}

func TestIfInorderTraversalOfASingleNodeTreeWorks(t *testing.T) {
	binaryTree := BinarySearchTree{}
	traversalEntries := []int{}
	assert.Equal(t, traversalEntries, binaryTree.TraverseInOrder())
}

func TestIfPreOrderTraversalOfARootWithOneLevelCHildrenWorksCorrectly(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.Add(&BSTNode{data: 1})
	binaryTree.Add(&BSTNode{data: 5})
	traversalEntries := []int{1, 3, 5}
	assert.Equal(t, traversalEntries, binaryTree.TraverseInOrder())
}

func TestIfInOrderTraversalOfABigTreeWorksCorrectly(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	traversalEntries := []int{1, 3, 5, 6, 8, 11, 12, 13, 14, 15, 17, 19}
	assert.Equal(t, traversalEntries, binaryTree.TraverseInOrder())
}

func TestIfPostOrderTraversalOfASingleRootTreeWorksCorrectly(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 3})
	binaryTree.Add(&BSTNode{data: 1})
	binaryTree.Add(&BSTNode{data: 5})
	traversalEntries := []int{1, 5, 3}
	assert.Equal(t, traversalEntries, binaryTree.TraversePostOrder())
}

func TestIfPostOrderTraversalOfABigTreeWorksCorrectly(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	traversalEntries := []int{1, 5, 3, 8, 6, 12, 14, 13, 19, 17, 15, 11}
	assert.Equal(t, traversalEntries, binaryTree.TraversePostOrder())
}

func TestIfLevelOrderTraversalOfANormalTreeWorksCorrectly(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	traversalEntries := []int{11, 6, 15, 3, 8, 13, 17, 1, 5, 12, 14, 19}
	assert.Equal(t, traversalEntries, binaryTree.TraverseLevelOrder())
}
