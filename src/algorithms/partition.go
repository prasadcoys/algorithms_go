package algorithms

// Selects the pivot and returns the pivot element and its position in the array
func SelectPivot(numbers []int) int {
	return len(numbers) - 1
}

func MoveLeftAndRightPointersTillCondition(numbers []int, pivot int) (int, int) {
	leftPointer := 0
	rightPointer := pivot - 1
	for {
		if leftPointer >= rightPointer {
			break
		}
		if numbers[leftPointer] > numbers[pivot] {
			if numbers[rightPointer] < numbers[pivot] {
				return leftPointer, rightPointer
			} else {
				rightPointer--
			}
		} else {
			leftPointer++
		}
	}
	return leftPointer, rightPointer
}
