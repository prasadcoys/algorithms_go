package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

}
