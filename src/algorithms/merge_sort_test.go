package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfMergeSortWorksCorrectlyForTwoElementsNotSorted(t *testing.T) {
	entries := []int{4, 3, 2, 1}
	assert.Equal(t, []int{1, 2, 3, 4}, MergeSort(entries))
}
