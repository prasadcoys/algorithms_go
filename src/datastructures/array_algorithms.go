package datastructures

import (
	"fmt"
	"math"
	"sort"
)

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

func GetPairMakingGivenSumExistsIn(entries []int, sum int) [][]int {
	i := 0
	pairs := [][]int{}
	j := len(entries) - 1
	for {

		if i >= j {
			break
		}
		if i > 0 && entries[i] == entries[i-1] {
			i = i + 1
			continue

		}
		currentSum := entries[i] + entries[j]
		if currentSum == sum {
			pairs = append(pairs, []int{entries[i], entries[j]})
			i++
		} else if currentSum > sum {
			j--
		} else if currentSum < sum {
			i++
		}
	}
	return pairs
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

func CalculateSumOfAllPairs(nums []int) int {
	sum := 0
	for i := len(nums) - 1; i > -1; i-- {
		for j := i - 1; j >= 0; j-- {
			difference := nums[i] - nums[j]
			if (difference < 0 && difference < -1) || (difference > 0 && difference > 1) {
				sum = sum + difference
			}
		}
	}
	return sum
}

func FindTripletsWithSumZero(nums []int) [][]int {
	triplets := [][]int{}
	sort.Ints(nums)
	counter := 0
	for _, num := range nums {
		if counter > 0 && nums[counter] == nums[counter-1] {
			fmt.Println("skipping ", nums[counter])
			counter = counter + 1
			continue
		}
		diff := 0 - num
		pairs := GetPairMakingGivenSumExistsIn(nums[counter+1:], diff)

		for _, pair := range pairs {
			if len(pair) > 0 {
				pair = append(pair, num)
				triplets = append(triplets, pair)
			}
		}
		counter = counter + 1
	}
	return triplets
}

func FindTripletsWithSumClosestTo(nums []int, target int) int {
	count := len(nums)
	closestSum := math.MaxInt32
	closestDiff := math.MaxInt32

	sort.Ints(nums)
	for i := 0; i < count-2; i++ {
		diffToTarget := target - nums[i]
		currentSum := nums[i] + diffToTarget - GetClosestSum(nums[i+1:], diffToTarget)
		// if int(math.Abs(float64(target-currentSum))) < closestDiff {
		// 	closestDiff = int(math.Abs(float64(target - currentSum)))
		// 	closestSum = currentSum
		// }
		absoluteDifference := Modulus(target - currentSum)
		if absoluteDifference < closestDiff {
			closestDiff = absoluteDifference
			closestSum = currentSum
		}

	}
	return closestSum
}

func GetClosestSum(entries []int, sum int) int {
	i := 0
	j := len(entries) - 1
	currentSum := 0
	leastDiff := math.MaxInt32
	for {
		if i >= j {
			break
		}
		currentSum = entries[i] + entries[j]
		// if math.Abs(float64(leastDiff)) > math.Abs(float64(sum-currentSum)) {
		// 	leastDiff = sum - currentSum
		// }
		if Modulus(leastDiff) > Modulus(sum-currentSum) {
			leastDiff = sum - currentSum
		}

		// leastDiff = int(math.Min(math.Abs(float64(sum-currentSum)), math.Abs(float64(leastDiff))))
		if currentSum == sum {
			return 0
		} else if currentSum > sum {
			j--
		} else if currentSum < sum {
			i++
		}
	}

	return leastDiff
}

func Modulus(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}
