package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfWeCanGetCountOfAllElementsLessThanACertainNumber(t *testing.T) {
	higherLimit := 8
	binaryTree := CreateBigBinaryTree()
	assert.Equal(t, []int{1, 3, 5, 6}, binaryTree.FetchAllElementsLessThan(higherLimit))
}

func TestIfWeCanGetCountOfAllElementsGreaterThanACertainNumber(t *testing.T) {
	lowerLimit := 8
	binaryTree := CreateBigBinaryTree()
	assert.Equal(t, []int{11, 12, 13, 14, 15, 17, 19}, binaryTree.FetchAllElementsGreaterThan(lowerLimit))
}

func TestIfWeCanPrintTheElementsInRange(t *testing.T) {
	binaryTree := CreateBigBinaryTree()
	lowerLimit := 8
	upperLimit := 17
	expectedMatches := []int{11, 12, 13, 14, 15}
	assert.Equal(t, expectedMatches, binaryTree.FindElementsInRange(lowerLimit, upperLimit))
}
