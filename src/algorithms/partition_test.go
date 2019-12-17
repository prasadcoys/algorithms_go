package algorithms

import (
	"fmt"
	"testing"
)

func TestIfRightMostElementIsSelectedAsPivot(t *testing.T) {
	numbers := []int{5, 4, 0, 1, 2, 3}
	expectedPivotIndex := 5
	pivotIndex := SelectPivot(numbers)
	if expectedPivotIndex != pivotIndex {
		t.FailNow()
	}
}

func TestIfMovingTheLeftAndRightPointersWorksAsExpected(t *testing.T) {
	numbers := []int{5, 4, 0, 1, 2, 3}
	expectedLeftPointerPosition := 0
	expectedRightPointerPosition := 4
	currentLeft, currentRight := MoveLeftAndRightPointersTillCondition(numbers, SelectPivot(numbers))
	if expectedLeftPointerPosition != currentLeft || expectedRightPointerPosition != currentRight {
		t.FailNow()
	}
}

func TestIfMovingLeftAndRightPointerWorksIfConditionBreaksAfterScenarioOne(t *testing.T) {
	numbers := []int{1, 2, 4, 5, 0, 3}
	expectedLeftPointerPosition := 2
	expectedRightPointerPosition := 4
	currentLeft, currentRight := MoveLeftAndRightPointersTillCondition(numbers, SelectPivot(numbers))
	fmt.Printf("Left : %d ,Right : %d\n", currentLeft, currentRight)
	if expectedLeftPointerPosition != currentLeft || expectedRightPointerPosition != currentRight {
		t.FailNow()
	}
}
