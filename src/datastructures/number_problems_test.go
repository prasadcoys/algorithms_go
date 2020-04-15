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

func TestIf1IsConvertedIntoI(t *testing.T) {
	assert.Equal(t, "I", ConvertIntegerToRomanNumber(1))
}

func TestIf3IsConvertedIntoIII(t *testing.T) {
	assert.Equal(t, "III", ConvertIntegerToRomanNumber(3))
}

func TestIf4IsConvertedToIV(t *testing.T) {
	assert.Equal(t, "IV", ConvertIntegerToRomanNumber(4))
}

func TestIf18IsConvertedToVIII(t *testing.T) {
	assert.Equal(t, "VIII", ConvertIntegerToRomanNumber(8))
}

func TestIf19IsConvertedToXIX(t *testing.T) {
	assert.Equal(t, "XIX", ConvertIntegerToRomanNumber(19))
}

func TestIf3000IsConvertedToMMM(t *testing.T) {
	assert.Equal(t, "MMM", ConvertIntegerToRomanNumber(3000))
}

func TestIf3400IsConvertedToMMMCD(t *testing.T) {
	assert.Equal(t, "MMMCD", ConvertIntegerToRomanNumber(3400))
}

func TestIf3900IsConvertedToMMMD(t *testing.T) {
	assert.Equal(t, "MMMCM", ConvertIntegerToRomanNumber(3900))
}

func TestIf1994IsConvertedToMCMXCIV(t *testing.T) {
	assert.Equal(t, "MCMXCIV", ConvertIntegerToRomanNumber(1994))
}

func TestIfWeCanPrintAllPossibleLetterCombinationsOf23FromThePhonePad(t *testing.T) {
	expectedCombinations := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	assert.Equal(t, expectedCombinations, GetAllPossibleCombinationsOfPhonePadNumberRec("23"))
}

func TestIfWeCanPrintAllPossibleLetterCombinationsOfEmptyFromThePhonePad(t *testing.T) {
	expectedCombinations := []string{}
	assert.Equal(t, expectedCombinations, GetAllPossibleCombinationsOfPhonePadNumberRec(""))
}

func TestIfWeCanPrintAllPossibleLetterCombinationsOf2FromThePhonePad(t *testing.T) {
	expectedCombinations := []string{"a", "b", "c"}
	assert.Equal(t, expectedCombinations, GetAllPossibleCombinationsOfPhonePadNumberRec("2"))
}

func TestIfWeCanPrintAllPossibleLetterCombinationsOf999FromThePhonePad(t *testing.T) {
	expectedCombinations := []string{"www", "wwx", "wwy", "wwz", "wxw", "wxx", "wxy", "wxz", "wyw", "wyx", "wyy", "wyz", "wzw", "wzx", "wzy", "wzz", "xww", "xwx", "xwy", "xwz", "xxw", "xxx", "xxy", "xxz", "xyw", "xyx", "xyy", "xyz", "xzw", "xzx", "xzy", "xzz", "yww", "ywx", "ywy", "ywz", "yxw", "yxx", "yxy", "yxz", "yyw", "yyx", "yyy", "yyz", "yzw", "yzx", "yzy", "yzz", "zww", "zwx", "zwy", "zwz", "zxw", "zxx", "zxy", "zxz", "zyw", "zyx", "zyy", "zyz", "zzw", "zzx", "zzy", "zzz"}
	assert.Equal(t, expectedCombinations, GetAllPossibleCombinationsOfPhonePadNumberRec("999"))
}

func TestIfWeCanMultiplyAnArrayOfStringsWithAnotherArray(t *testing.T) {
	expectedCombinations := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	assert.Equal(t, expectedCombinations, GetAllPossibleCombinationsOfPhonePadNumberRec("23"))
}
