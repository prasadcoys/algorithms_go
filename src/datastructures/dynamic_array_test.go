package datastructures

import (
	"testing"
)

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

func TestSetAtIndex(t *testing.T) {
	numbers := &DynamicArray{}
	numbers.elements = []int{1, 2, 3, 4, 5}
	numbers.SetAt(4, 10)
	if numbers.Get(4) != 10 {
		t.FailNow()
	}
}

func TestIfAddOfElementWorks(t *testing.T) {
	numbers := &DynamicArray{}
	numbers.elements = []int{1, 2, 3, 4, 5}
	numbers.Add(10)
	if numbers.Get(5) != 10 {
		t.FailNow()
	}
	if numbers.Size() != 6 {
		t.FailNow()
	}
}

func TestIfRemovalAtIndexWorks(t *testing.T) {
	numbers := &DynamicArray{}
	numbers.elements = []int{1, 2, 3, 4, 5}
	numbers.RemoveAt(3)
	if numbers.Get(3) != 5 {
		t.FailNow()
	}
	if numbers.Size() != 4 {
		t.FailNow()
	}
}

func TestIfRemoveOfElementWorks(t *testing.T) {
	numbers := &DynamicArray{}
	numbers.elements = []int{1, 2, 3, 4, 5}
	numbers.Remove(4)
	if numbers.Get(3) != 5 {
		t.FailNow()
	}
	if numbers.Size() != 4 {
		t.FailNow()
	}
}
