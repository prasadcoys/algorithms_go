package datastructures

import (
	"math"
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

	isNegative := false
	numbersMap := make(map[int]int)
	numbersMap[45] = -1
	number := 0

	j := 0
	for i := 48; i <= 57; i++ {
		numbersMap[i] = j
		j = j + 1
	}

	for _, c := range text {
		if c == 32 || c == 43 {
			continue
		}
		numberInInteger, ok := numbersMap[int(c)]
		if !ok {
			break
		}
		if numberInInteger < 0 {
			isNegative = true
			continue
		}
		number = (number * 10) + numberInInteger
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
