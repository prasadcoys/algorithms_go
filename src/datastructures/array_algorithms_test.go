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
