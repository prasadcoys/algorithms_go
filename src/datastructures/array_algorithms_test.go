package datastructures

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfGivenAnArrayOfSortedNumbersWeCanCheckIfPairWithGivenSumExists(t *testing.T) {
	entries := []int{1, 2, 3}
	assert.True(t, DoesPairMakingGivenSumExistsIn(entries, 5))
}

func TestIfWeCanBalanceAnEvenArrayIfRightSideIsBigger(t *testing.T) {
	entries := []int{1, 2, 1, 2, 1, 3}
	assert.Equal(t, 2, MinimumRequiredForBalancing(entries))
}

func TestIfWeCanBalanceAnEvenArrayIfLeftSideIsBigger(t *testing.T) {
	entries := []int{20, 10}
	assert.Equal(t, 10, MinimumRequiredForBalancing(entries))
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

func TestIfGivenAnUnsortedArrayWeCanFindPairExistsWithGivenSum(t *testing.T) {
	entries := []int{8, 7, 1, 3}
	assert.EqualValues(t, []int{1, 2}, PairWithGivenSumInUnsortedArray(entries, 8))
}

func TestIfArrayCanBeRotated(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	expected_after_rotation := []int{3, 4, 5, 1, 2}
	assert.Equal(t, expected_after_rotation, RotateArray(nums, 2))
}

func TestIfArrayCanBeRotatedForABiggerArray(t *testing.T) {
	nums := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
	expected_after_rotation := []int{8, 10, 12, 14, 16, 18, 2, 4, 6}
	assert.Equal(t, expected_after_rotation, RotateArray(nums, 3))
}

func TestIfSumOfAllPairsOverAnArray(t *testing.T) {
	nums := []int{6, 6, 4, 4}
	assert.Equal(t, -8, CalculateSumOfAllPairs(nums))
	nums = []int{1, 2, 3, 1, 3}
	assert.Equal(t, 4, CalculateSumOfAllPairs(nums))
}

func TestIfThreeSumWithValidEntriesIsCalculated(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(FindTripletsWithSumZero(nums))

}

func TestIfThreeSumWithAllZeroEntriesIsCalculated(t *testing.T) {
	nums := []int{0, 0, 0, 0}

	fmt.Println(FindTripletsWithSumZero(nums))

}

func TestIfArrayWithLargeNumbersWorksCorrectly(t *testing.T) {
	nums := []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}
	fmt.Println(FindTripletsWithSumZero(nums))
}

func TestIf3SumClosestToTargetIsReturnedForSmallArray(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	actual := FindTripletsWithSumClosestTo(nums, 1)

	assert.Equal(t, 2, actual)
}

func TestIf3SumClosestToTargetIsReturnedWhenExactMatchIsFound(t *testing.T) {
	nums := []int{0, 1, 2}
	actual := FindTripletsWithSumClosestTo(nums, 3)

	assert.Equal(t, 3, actual)
}

func TestIf3SumClosestToTargetIsReturnedWhenExactMatchIsFoundWithDuplicatesInEntries(t *testing.T) {
	nums := []int{1, 1, -1, -1, 3}
	actual := FindTripletsWithSumClosestTo(nums, -1)

	assert.Equal(t, -1, actual)
}

func TestIf3SumClosestToTargetIsReturnedWhenNoExactMatchWithAVeryLargeElemetn(t *testing.T) {
	nums := []int{-1, 0, 1, 1, 55}
	actual := FindTripletsWithSumClosestTo(nums, 3)

	assert.Equal(t, 2, actual)
}

func TestIfLeastDiffWorks(t *testing.T) {
	nums := []int{1, 2}
	fmt.Println(GetClosestSum(nums, 2))
}

func TestIfModulusWorksCorrectly(t *testing.T) {
	assert.Equal(t, 1, Modulus(-1))
	assert.Equal(t, 1, Modulus(1))
}
