package algorithms

// SelectPivot returns the pivot and returns the pivot element and its position in the array
func SelectPivot(numbers []int) int {
	return len(numbers) - 1
}

//MoveLeftAndRightPointersTillCondition returns the left and right pointer position after finishing the traversal
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
			}
			rightPointer--
		} else {
			leftPointer++
		}
	}
	return leftPointer, rightPointer
}
