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
