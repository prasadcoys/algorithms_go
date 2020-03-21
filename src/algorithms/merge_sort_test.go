package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfMergeSortWorksCorrectlyForTwoElementsNotSorted(t *testing.T) {
	entries := []int{4, 3, 2, 1}
	assert.Equal(t, []int{1, 2, 3, 4}, MergeSort(entries))
}

func TestIfMergeSortWorksCorrectlyForTwoElementsNotSortedForPartiallySortedArray(t *testing.T) {
	entries := []int{20, 30, 25, 50, 10, 60, 40}
	assert.Equal(t, []int{10, 20, 25, 30, 40, 50, 60}, MergeSort(entries))
}
