package datastructures

import "testing"

func TestIfCreationOfNewDynamicArrayWorks(t *testing.T) {
	numbers := &DynamicArray{}
	if numbers.Size() != 0 {
		t.FailNow()
	}

}

func TestIfIsEmptyWorksIfArrayIsEmpty(t *testing.T) {
	numbers := &DynamicArray{}
	if !numbers.IsEmpty() {
		t.FailNow()
	}
}

func TestIfEmptyWorksIfArrayIsNotEmpty(t *testing.T) {
	numbers := &DynamicArray{}
	numbers.elements = []int{1, 2, 3}
	if numbers.IsEmpty() {
		t.FailNow()
	}
}

func TestIfSizeIsCalculatedCorrectly(t *testing.T) {
	numbers := &DynamicArray{}
	numbers.elements = []int{1, 2, 3}
	if numbers.Size() != 3 {
		t.FailNow()
	}
}

func TestGetAtIndex(t *testing.T) {
	numbers := &DynamicArray{}
	numbers.elements = []int{1, 2, 3}
	if numbers.Get(2) != 3 {
		t.FailNow()
	}
}
