package datastructures

import (
	"fmt"
	"strconv"
	"strings"
)

func letterCasePermutation(S string) []string {
	permutations := []string{}

	letters := strings.Split(S, "")
	recursivelyAddLetterCasePermutations("", letters, &permutations)

	return permutations

}

func recursivelyAddLetterCasePermutations(currentString string, letters []string, permutations *[]string) {
	if len(letters) == 0 {
		(*permutations) = append(*permutations, currentString)
		return
	}
	character := []rune(letters[0])
	if character[0] >= 48 && character[0] <= 57 {
		recursivelyAddLetterCasePermutations(currentString+letters[0], letters[1:], permutations)
	} else {
		recursivelyAddLetterCasePermutations(currentString+strings.ToUpper(letters[0]), letters[1:], permutations)
		recursivelyAddLetterCasePermutations(currentString+strings.ToLower(letters[0]), letters[1:], permutations)
	}

}

func readBinaryWatch(num int) []string {
	permutations := []string{}
	recursivelyCalculateTimeCombinations("", &permutations, num)
	return permutations
}

func recursivelyCalculateTimeCombinations(currentTime string, permutations *[]string, num int) {
	if num == 0 {
		charsLeft := 10 - len(currentTime)
		currentTime = currentTime + strings.Repeat("0", charsLeft)
		hoursInInt, minutesInInt := getTimeFromBinary(currentTime)
		if hoursInInt > 11 || minutesInInt > 59 {
			return
		}
		var minutesPadded string
		if minutesInInt < 10 {
			minutesPadded = fmt.Sprintf("%s%d", "0", minutesInInt)
		} else {
			minutesPadded = fmt.Sprintf("%d", minutesInInt)
		}
		time := fmt.Sprintf("%d%s%s", hoursInInt, ":", minutesPadded)
		*permutations = append(*permutations, time)
		return
	} else if len(currentTime) == 10 {
		return
	} else {
		recursivelyCalculateTimeCombinations(currentTime+"0", permutations, num)
		recursivelyCalculateTimeCombinations(currentTime+"1", permutations, num-1)
	}
}

func getTimeFromBinary(timeInBinary string) (int64, int64) {
	hours := timeInBinary[0:4]
	minutes := timeInBinary[4:]
	var minutesInInt int64
	minutesInInt, _ = strconv.ParseInt(minutes, 2, 32)

	hoursInInt, _ := strconv.ParseInt(hours, 2, 32)
	//fmt.Println("" + int(hoursInInt) + ":" + int(minutesInInt))

	return hoursInInt, minutesInInt
}

func combinationSum2(candidates []int, target int) [][]int {
	permutations := [][]int{}
	recursivelyCalculateCombinationSum2(candidates, target, &permutations, []int{})
	falseList := candidates
	falseList[0] = 0
	recursivelyCalculateCombinationSum2(falseList, target, &permutations, []int{})
	fmt.Println(permutations)
	return permutations
}

func recursivelyCalculateCombinationSum2(candidates []int, target int, permutations *[][]int, currentPermutation []int) {
	fmt.Println(currentPermutation, target)
	if len(candidates) == 0 {
		return
	}
	currentPermutation = append(currentPermutation, candidates[0])
	target = target - candidates[0]
	if target == 0 {
		*permutations = append(*permutations, currentPermutation)
		return
	}
	if len(candidates) == 1 {
		return
	}
	recursivelyCalculateCombinationSum2(candidates[1:], target, permutations, currentPermutation)
	falseList := candidates[1:]
	falseList[0] = 0
	recursivelyCalculateCombinationSum2(falseList, target, permutations, currentPermutation)
	return
}

func subsets(elements []int) [][]int {
	subsets := [][]int{}
	recursivelyAddSubsetToPermutation(&subsets, elements, []int{})
	fmt.Println(subsets)
	return subsets
}

func recursivelyAddSubsetToPermutation(subsets *[][]int, elements []int, currentList []int) {
	if len(elements) == 0 {
		*subsets = append(*subsets, currentList)
		return
	}
	{

		recursivelyAddSubsetToPermutation(subsets, elements[1:], currentList)
		recursivelyAddSubsetToPermutation(subsets, elements[1:], append(currentList, elements[0]))
	}
	return
}
