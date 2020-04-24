package datastructures

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfLetterCaseCombinationForGeneralCaseWorksCorrectly(t *testing.T) {
	expected := []string{"a1b2", "a1B2", "A1b2", "A1B2"}
	assert.ElementsMatch(t, expected, letterCasePermutation("a1b2"))
}

func TestIfLetterCaseCombinationForCaseWithOnlyOneLetterWorksCorrectly(t *testing.T) {
	expected := []string{"3z4", "3Z4"}
	assert.ElementsMatch(t, expected, letterCasePermutation("3z4"))
}

func TestIfLetterCaseCombinationForCaseWithOnlyLettersWorksCorrectly(t *testing.T) {
	expected := []string{"12345"}
	assert.ElementsMatch(t, expected, letterCasePermutation("12345"))
}

func TestIfLetterCombinationForCaseWithOnlyLettersWorksCorrectly(t *testing.T) {
	expected := []string{"abc", "abC", "aBc", "aBC", "Abc", "AbC", "ABc", "ABC"}
	assert.ElementsMatch(t, expected, letterCasePermutation("abc"))
	assert.ElementsMatch(t, expected, letterCasePermutation("ABC"))
}

func TestIfBinaryClocksTimeCanBePredictedForOne(t *testing.T) {
	expected := []string{"1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"}
	fmt.Println(strconv.ParseInt("0100", 2, 32))
	assert.ElementsMatch(t, expected, readBinaryWatch(1))
}

func TestIfBinaryClockTimeCanBePredictedForTwo(t *testing.T) {
	expected := []string{"0:03", "0:05", "0:06", "0:09", "0:10", "0:12", "0:17", "0:18", "0:20", "0:24", "0:33", "0:34", "0:36", "0:40", "0:48", "1:01", "1:02", "1:04", "1:08", "1:16", "1:32", "2:01", "2:02", "2:04", "2:08", "2:16", "2:32", "3:00", "4:01", "4:02", "4:04", "4:08", "4:16", "4:32", "5:00", "6:00", "8:01", "8:02", "8:04", "8:08", "8:16", "8:32", "9:00", "10:00"}
	assert.ElementsMatch(t, expected, readBinaryWatch(2))
}

func TestIfWeCanSolveCombinationSumWithOutRepitition(t *testing.T) {
	expected := [][]int{
		{1, 7},
		{1, 2, 5},
		{2, 6},
		{1, 1, 6},
	}
	assert.ElementsMatch(t, expected, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func TestIfWeCanListAllSubSets(t *testing.T) {
	expected := [][]int{{1, 2, 3}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {}}
	assert.ElementsMatch(t, expected, subsets([]int{1, 2, 3}))
}

func TestIfWeCanListAllSubSetsForBiggerSet(t *testing.T) {
	expected := [][]int{{}, {9}, {0}, {0, 9}, {3}, {3, 9}, {0, 3}, {0, 3, 9}, {5}, {5, 9}, {0, 5}, {0, 5, 9}, {3, 5}, {3, 5, 9}, {0, 3, 5}, {0, 3, 5, 9}, {7}, {7, 9}, {0, 7}, {0, 7, 9}, {3, 7}, {3, 7, 9}, {0, 3, 7}, {0, 3, 7, 9}, {5, 7}, {5, 7, 9}, {0, 5, 7}, {0, 5, 7, 9}, {3, 5, 7}, {3, 5, 7, 9}, {0, 3, 5, 7}, {0, 3, 5, 7, 9}}
	assert.ElementsMatch(t, expected, subsets([]int{9, 0, 3, 5, 7}))
}
