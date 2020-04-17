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

func TestIfWeCanGenerateParanthesesFor1(t *testing.T) {
	expected := []string{"()"}
	assert.Equal(t, expected, GetParanthesesCombination(1))
}

func TestIfWeCanGenerateParanthesesFor2(t *testing.T) {
	expected := []string{"()()", "(())"}
	assert.Equal(t, expected, GetParanthesesCombination(2))
}

func TestIfWeCanGenerateParanthesesFor3(t *testing.T) {
	expected := []string{"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()"}
	assert.ElementsMatch(t, expected, GetParanthesesCombination(3))
}

func TestIfWeCanGenerateParanthesesFor4(t *testing.T) {
	expected := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}
	actual := GetParanthesesCombination(4)
	assert.ElementsMatch(t, expected, actual)
}

func TestIfWeCanBracketEverySingleParanthesis(t *testing.T) {
	expected := []string{"(())()", "()(())"}
	actual := BracketEverySingleparanthesis("()()")
	assert.ElementsMatch(t, expected, actual)
}
