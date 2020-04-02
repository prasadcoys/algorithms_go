package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfNumberCanBeReversed(t *testing.T) {
	number := 123
	assert.Equal(t, 321, ReverseNumber(number))
}

func TestIfNumberWithLastDigitZeroIsReversed(t *testing.T) {
	number := 210
	assert.Equal(t, 12, ReverseNumber(number))
}

func TestIfReversingOfNegativeNumberWorksCorrectly(t *testing.T) {
	number := -123
	assert.Equal(t, -321, ReverseNumber(number))
}
