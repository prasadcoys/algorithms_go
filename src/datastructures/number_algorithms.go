package datastructures

import (
	"fmt"
	"math"
	"strings"
)

func ReverseNumber(number int) int {
	reversedNumber := 0
	for {
		if number == 0 {
			break
		}
		reversedNumber = (reversedNumber * 10) + (number % 10)
		number = number / 10
	}
	if reversedNumber < -2147483648 || reversedNumber > 2147483647 {
		return 0
	}
	return reversedNumber
}

func AtoI(text string) int {
	text = strings.TrimSpace(text)
	isNegative := false
	isPositive := false
	numberStarted := false
	numbersMap := make(map[int]int)
	numbersMap[45] = -1
	number := 0

	j := 0
	for i := 48; i <= 57; i++ {
		numbersMap[i] = j
		j = j + 1
	}

	for _, c := range text {
		if c == 32 {
			break
		}
		if c == 43 {
			if numberStarted {
				break
			}
			if isNegative || isPositive {
				break
			}
			isPositive = true
			continue
		}

		numberInInteger, ok := numbersMap[int(c)]
		if !ok {
			break
		}
		if numberInInteger < 0 {
			if numberStarted {
				break
			}
			if isPositive || isNegative {
				break
			}
			isNegative = true
			continue
		}
		numberStarted = true

		number = (number * 10) + numberInInteger
		if number*-1 < math.MinInt32 {
			break
		}
	}

	if isNegative {
		if number*-1 < math.MinInt32 {
			return math.MinInt32
		}
		return number * -1
	}
	if number > math.MaxInt32 {
		return math.MaxInt32
	}
	return number
}

func IsNumberPalindrome(num int) bool {
	tNum := num
	if num < 0 {
		return false
	}
	reversedNumber := 0
	for {
		if num == 0 {
			break
		}
		reversedNumber = reversedNumber*10 + num%10
		num = num / 10

	}
	return (tNum - reversedNumber) == 0
}

func ConvertRomanNumberToInteger(roman string) int {
	romanToNumbers := make(map[rune]int, 0)
	romanToNumbers['I'] = 1
	romanToNumbers['V'] = 5
	romanToNumbers['X'] = 10
	romanToNumbers['L'] = 50
	romanToNumbers['C'] = 100
	romanToNumbers['D'] = 500
	romanToNumbers['M'] = 1000
	number := 0
	prevDigit := 0
	for _, digit := range roman {
		currentNumber := romanToNumbers[digit]
		if currentNumber > prevDigit {
			number = number - prevDigit + currentNumber - prevDigit
		} else {
			number = number + currentNumber
		}
		prevDigit = currentNumber
	}
	return number
}

func ConvertIntegerToRomanNumber(number int) string {
	romanNumber := ""
	numbersToRoman := make(map[int]string, 7)
	numbersToRoman[1000] = "M"
	numbersToRoman[500] = "D"
	numbersToRoman[100] = "C"
	numbersToRoman[50] = "L"
	numbersToRoman[10] = "X"
	numbersToRoman[5] = "V"
	numbersToRoman[1] = "I"
	keys := []int{1000, 500, 100, 50, 10, 5, 1}

	for _, key := range keys {
		if number == 0 {
			break
		}
		digit := number / key
		for i := 0; i < digit; i++ {
			romanNumber = romanNumber + numbersToRoman[key]
		}
		number = number % key
		if key > 1 {
			penultimateDigit := number / (key - (key / 10))
			if penultimateDigit == 1 {
				romanNumber = romanNumber + numbersToRoman[key/10] + numbersToRoman[key]
				number = number % (key - (key / 10))
			}
			penultimateDigit = number / (key - (key / 5))
			fmt.Println(penultimateDigit, number, key)
			if penultimateDigit == 1 {
				romanNumber = romanNumber + numbersToRoman[key/5] + numbersToRoman[key]
				number = number % (key - (key / 5))
			}
		}

	}
	return romanNumber

}
