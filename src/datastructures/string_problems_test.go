package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfStringWithThreeCharactersCanBeBrokenIntoThreeRows(t *testing.T) {
	text := "ABC"
	texts := [][]string{{"A"}, {"B"}, {"C"}}
	assert.Equal(t, texts, buildZigZagArray(text, 3))
}

func TestIfStringWithMultipleCharacters(t *testing.T) {
	text := "PAYPALISHIRING"
	assert.Equal(t, "PAHNAPLSIIGYIR", ZigZagText(text, 3))
}

func TestIfStringWithMultipleCharactersForFourRows(t *testing.T) {
	text := "PAYPALISHIRING"
	assert.Equal(t, "PINALSIGYAHRPI", ZigZagText(text, 4))
}

func TestIfWeCanFindLongestCommonPrefix(t *testing.T) {
	texts := []string{"flower", "flow", "flight"}
	assert.Equal(t, "fl", FindLongestCommonPrefix(texts))
}

func TestIfWeCanFindLongestCommonPrefixWhenNonExists(t *testing.T) {
	texts := []string{"car", "flow", "animal"}
	assert.Equal(t, "", FindLongestCommonPrefix(texts))
}

func TestIFWeReturnFullStringWhnAllStringsAreEqual(t *testing.T) {
	texts := []string{"car", "car", "car"}
	assert.Equal(t, "car", FindLongestCommonPrefix(texts))
}

func TestIfWeReturnEmptyWhenNoArrayIsEmpty(t *testing.T) {
	texts := []string{}
	assert.Equal(t, "", FindLongestCommonPrefix(texts))
}

func TestIFWeReturnFullStringWhenFirstStringIsEmpty(t *testing.T) {
	texts := []string{"", "car", "car"}
	assert.Equal(t, "", FindLongestCommonPrefix(texts))
}

func TestIfValidParanthesesReturnsTrue(t *testing.T) {
	assert.Equal(t, true, IsValidParantheses("()"))
	assert.Equal(t, true, IsValidParantheses("()[]{}"))
	assert.Equal(t, true, IsValidParantheses("{[]}"))
}

func TestIfInvalidParanthesesReturnsFalse(t *testing.T) {
	assert.Equal(t, false, IsValidParantheses(")("))
	assert.Equal(t, false, IsValidParantheses("(]"))
	assert.Equal(t, false, IsValidParantheses("([)]"))
}
