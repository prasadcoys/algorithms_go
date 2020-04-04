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

func TestConvertStringToIntegerWorks(t *testing.T) {
	text := "42"
	assert.Equal(t, 42, AtoI(text))
}

func TestConvertStringWithSpaceToIntegerWorks(t *testing.T) {
	text := "  42"
	assert.Equal(t, 42, AtoI(text))
}

func TestIfConversionOfNegativeNumberWorks(t *testing.T) {
	text := "-42"
	assert.Equal(t, -42, AtoI(text))
}

func TestIfAnyAdditionalStringAfterNumberIsIgnored(t *testing.T) {
	text := "-42 and some text"
	assert.Equal(t, -42, AtoI(text))
}

func TestIfTextBeforeNumberReturnsZero(t *testing.T) {
	text := "some txt and number"
	assert.Equal(t, 0, AtoI(text))
}

func TestIfOvershootingOfValuesResultInMaxValue(t *testing.T) {
	text := "-91283472332"
	assert.Equal(t, -2147483648, AtoI(text))
}

func TestIfNegativeNumberStringWithSpaceIsCorrectlyTransformed(t *testing.T) {
	text := "   -42"
	assert.Equal(t, -42, AtoI(text))
}

func TestIfPositiveNumberStringWithPlusSignIsCorrectlyTransformed(t *testing.T) {
	text := "+42"
	assert.Equal(t, 42, AtoI(text))
}

func TestIfStringWithPlusAndMinusSignIsCorrectlyTransformed(t *testing.T) {
	text := "+-42"
	assert.Equal(t, 0, AtoI(text))
}

func TestIfNumberGreaterThanInt32RangeIsCheckedForOverflow(t *testing.T) {
	text := "9223372036854775808"
	assert.Equal(t, 2147483647, AtoI(text))
}

func TestIfExpressionIsIgnoredt(t *testing.T) {
	text := "1-0"
	assert.Equal(t, 0, AtoI(text))
}

func TestIfDoubleSignsAreIgnored(t *testing.T) {
	text := " ++1"
	assert.Equal(t, 0, AtoI(text))
}
