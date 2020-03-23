package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfGivenAnArrayOfSortedNumbersWeCanCheckIfPairWithGivenSumExists(t *testing.T) {
	entries := []int{1, 2, 3}
	assert.True(t, DoesPairMakingGivenSumExistsIn(entries, 5))
}

func TestIfGivenAnArrayOfSortedNumbersWeCanCheckIfPairWithGivenSumDoesNotExist(t *testing.T) {
	entries := []int{1, 3, 4, 5, 6}
	assert.False(t, DoesPairMakingGivenSumExistsIn(entries, 12))
}

func TestIfMergeOfTwoSortedArraysWorksCorrectlyWithDuplicateElements(t *testing.T) {
	entries_1 := []int{2, 3, 4, 5, 6}
	entries_2 := []int{1, 2, 3, 6, 7}
	expected_entries := []int{1, 2, 2, 3, 3, 4, 5, 6, 6, 7}
	assert.Equal(t, expected_entries, MergeTwoSortedArrays(entries_1, entries_2))
}
