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
		fmt.Println(number)
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
