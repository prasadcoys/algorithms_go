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

func TestIfNodeCanBeAddedToAnNonEmptyBinarySearchTree(t *testing.T) {
	binaryTree := BinarySearchTree{}
	binaryTree.Add(&BSTNode{data: 23})
	binaryTree.Add(&BSTNode{data: 30})
	if binaryTree.root.right.data != 30 {
		t.Fail()
	}
	if binaryTree.root.left != nil {
		t.Fail()
	}
}
