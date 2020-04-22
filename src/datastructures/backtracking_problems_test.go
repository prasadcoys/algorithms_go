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
