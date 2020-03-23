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
