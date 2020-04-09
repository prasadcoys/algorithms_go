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

func TestIfDoubleSignsAreIgnored(t *testing.T) {
	text := " ++1"
	assert.Equal(t, 0, AtoI(text))
}

func TestIfNegativeNumberIsNotPalindrome(t *testing.T) {
	num := -121
	assert.Equal(t, false, IsNumberPalindrome(num))
}

func TestIfPositivePalindromNumberReturnsTrue(t *testing.T) {
	num := 121
	assert.Equal(t, true, IsNumberPalindrome(num))
}

func TestIfPositiveNonPalidromeNumberReturnsFalse(t *testing.T) {
	num := 123
	assert.Equal(t, false, IsNumberPalindrome(num))
}

func TestIfRomanLetterICanBeConverted(t *testing.T) {
	romanNumber := "I"
	assert.Equal(t, 1, ConvertRomanNumberToInteger(romanNumber))
}

func TestIfRomanLetterIIICanBeConverted(t *testing.T) {
	romanNumber := "LVIII"
	assert.Equal(t, 58, ConvertRomanNumberToInteger(romanNumber))
}

func TestIfRomanLetterIXCanBeConverted(t *testing.T) {
	romanNumber := "IX"
	assert.Equal(t, 9, ConvertRomanNumberToInteger(romanNumber))
}

func TestIfRomanCMCanBeConverted(t *testing.T) {
	romanNumber := "CM"
	assert.Equal(t, 900, ConvertRomanNumberToInteger(romanNumber))
}

func TestIfMCMXCIVCanBeConverted(t *testing.T) {
	romanNumber := "MCMXCIV"
	assert.Equal(t, 1994, ConvertRomanNumberToInteger(romanNumber))
}
