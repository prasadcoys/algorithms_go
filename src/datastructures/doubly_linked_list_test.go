package datastructures

import (
	"testing"
)

func TestIfDoubleLinkNodeIsCreated(t *testing.T) {
	node := CreateDoubleLinkNode(1)
	if node.Data != 1 {
		t.FailNow()
	}
	if node.Prev != nil {
		t.FailNow()
	}
	if node.Next != nil {
		t.FailNow()
	}
}

func TestIfEmptyDoublyLinkedListWorks(t *testing.T) {
	list := CreateDoublyLinkedList()
	if list.Head != nil {
		t.FailNow()
	}
	if list.Tail != nil {
		t.FailNow()
	}
	if list.Size() != 0 {
		t.FailNow()
	}
	if !list.IsEmpty() {
		t.FailNow()
	}
}

func TestIfAppendingToListWorks(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	if list.ToString() != "1 2 3" {
		t.FailNow()
	}
	if list.Size() != 3 {
		t.FailNow()
	}
	if list.IsEmpty() {
		t.FailNow()
	}
}

func TestIfClearingTheListWorks(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Clear()
	if !list.IsEmpty() {
		t.FailNow()
	}
	if "" != list.ToString() {
		t.FailNow()
	}

}

func TestIfAddAtIndexWorksCorrectly_dll_first_half(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.AddAt(2, 5)
	if "1 5 2 3 4" != list.ToString() {
		t.FailNow()
	}
	if 5 != list.Size() {
		t.FailNow()
	}
}

func TestIfAddAtIndexWorksCorrectly_dll_second_half(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.Append(8)
	list.Append(9)
	list.AddAt(7, 10)
	if "1 2 3 4 5 6 10 7 8 9" != list.ToString() {
		t.FailNow()
	}
	if 10 != list.Size() {
		t.FailNow()
	}
}

func TestIfPeekFirstWorksCorrectly(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	if 1 != list.PeekFirst() {
		t.FailNow()
	}
}

func TestIfPeekLastWorksCorrectly(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	if 5 != list.PeekLast() {
		t.FailNow()
	}
}

func TestIfRemoveAtNonBoundaryIndexInFirsthalfWorksCorrectly(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.RemoveAt(3)
	if "1 2 4 5 6 7" != list.ToString() {
		t.FailNow()
	}
}

func TestIfRemovalAtInvalidIndexGreaterBoundsThrowsException(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	err := list.RemoveAt(10)
	if err == nil {
		t.FailNow()
	}
}

func TestIfRemoveAtNonBoundaryIndexInSecondhalfWorksCorrectly(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	list.RemoveAt(6)
	if "1 2 3 4 5 7" != list.ToString() {
		t.FailNow()
	}
	if 6 != list.Size() {
		t.FailNow()
	}
}

func TestIfRemovalAtInvalidIndexLowerBoundsThrowsException(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	err := list.RemoveAt(-1)
	if err == nil {
		t.FailNow()
	}
}

func TestIfRemovalAtBeginningWorksCorrectly(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveFirst()
	if "2 3" != list.ToString() {
		t.FailNow()
	}
	if 2 != list.Size() {
		t.FailNow()
	}
}

func TestIfRemovalAtEndWorksCorrectly(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.RemoveLast()
	if "1 2" != list.ToString() {
		t.FailNow()
	}
	if 2 != list.Size() {
		t.FailNow()
	}
}

func TestIfIndexOfElementIsFoundCorrectly(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	if 2 != list.IndexOf(2) {
		t.FailNow()
	}
}

func TestIfIndexOfInvalidElementIsHandledCorrectly(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	if -1 != list.IndexOf(4) {
		t.FailNow()
	}
}

func TestIfRemoveAtEndWorksOnEmptyDoubleLinkList(t *testing.T) {
	list := CreateDoublyLinkedList()
	err := list.RemoveLast()
	if err == nil {
		t.FailNow()
	}
}

func TestIfRemoveAtBiginningWorksOnEmptyDoubleLinkList(t *testing.T) {
	list := CreateDoublyLinkedList()
	err := list.RemoveFirst()
	if err == nil {
		t.FailNow()
	}
}

func TestIfIndexOfInvalidElementInEmptyListIsHandledCorrectly(t *testing.T) {
	list := CreateDoublyLinkedList()
	if -1 != list.IndexOf(4) {
		t.FailNow()
	}
}

func TestIfContainsForTrueWorks(t *testing.T) {
	list := CreateDoublyLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	if !list.Contains(2) {
		t.FailNow()
	}
}
