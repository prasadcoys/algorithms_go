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
