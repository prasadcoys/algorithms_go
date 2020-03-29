package datastructures

func DoesPairMakingGivenSumExistsIn(entries []int, sum int) bool {
	i := 0
	j := len(entries) - 1
	for {
		if i >= j {
			break
		}
		currentSum := entries[i] + entries[j]
		if currentSum == sum {
			return true
		} else if currentSum > sum {
			j--
		} else if currentSum < sum {
			i++
		}
	}
	return false
}

func MergeTwoSortedArrays(array_1 []int, array_2 []int) []int {
	mergedArray := []int{}
	first := 0
	second := 0
	for {
		if first >= len(array_1) || second >= len(array_2) {
			break
		}
		if array_1[first] <= array_2[second] {
			mergedArray = append(mergedArray, array_1[first])
			first = first + 1
		} else {
			mergedArray = append(mergedArray, array_2[first])
			second = second + 1
		}
	}
	mergedArray = append(mergedArray, array_1[first:]...)
	mergedArray = append(mergedArray, array_2[second:]...)

	return mergedArray
}

func PairWithGivenSumInUnsortedArray(entries []int, target int) []int {
	pair := []int{-1, -1}
	differentials := make(map[int]int)
	for index, entry := range entries {
		if value, isEntryPresent := differentials[entry]; isEntryPresent {
			return []int{value, index}
		}
		differenceWithTarget := target - entry
		differentials[differenceWithTarget] = index
	}

	return pair
}

func RotateArray(nums []int, numberOfTimes int) []int {
	nums = append(nums, nums[0:numberOfTimes]...)
	nums = nums[numberOfTimes:]
	return nums
}

func MinimumRequiredForBalancing(numbers []int) int {
	differential := 0
	start := 0
	end := len(numbers) - 1
	for {
		if start > end {
			break
		}
		differential = differential + numbers[end] - numbers[start]
		start = start + 1
		end = end - 1

	}
	if differential < 0 {
		return differential * -1
	}
	return differential
}
