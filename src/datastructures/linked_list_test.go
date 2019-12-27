package datastructures

import "testing"

func TestIfLinkedListNodeIsCreated(t *testing.T) {
	node := CreateNode(2)
	if node.Data != 2 {
		t.FailNow()
	}
}

func TestIfLinkingOneNodeToAnotherIsWorking(t *testing.T) {
	nodeOne := CreateNode(1)
	nodeTwo := CreateNode(2)
	nodeOne.Next = &nodeTwo
	if nodeOne.Next.Data != 2 {
		t.FailNow()
	}
}

func TestIfAddingElementAtEndOfListWorks(t *testing.T) {

	nodeOne := CreateNode(1)
	list := LinkedList{
		Head: &nodeOne,
	}
	if list.Head.Data != 1 {
		t.FailNow()
	}
}

func TestIfAppendToLinkedListWorks(t *testing.T) {
	nodeOne := CreateNode(1)
	list := LinkedList{
		Head: &nodeOne,
	}
	nodeTwo := CreateNode(2)
	list.Append(nodeTwo)
	if nodeOne.Next.Data != 2 {
		t.FailNow()
	}
}

func TestIfAppendToLinkedListWithTwoElementsWorks(t *testing.T) {
	nodeOne := CreateNode(1)
	list := LinkedList{
		Head: &nodeOne,
	}
	nodeTwo := CreateNode(2)
	nodeThree := CreateNode(3)
	list.Append(nodeTwo)
	list.Append(nodeThree)
	if nodeOne.Next.Next.Data != 3 {
		t.FailNow()
	}
}

func TestIfSizeOfListIsCalculatedCorrectly(t *testing.T) {
	nodeOne := CreateNode(1)
	list := LinkedList{
		Head: &nodeOne,
	}
	nodeTwo := CreateNode(2)
	nodeThree := CreateNode(3)
	list.Append(nodeTwo)
	list.Append(nodeThree)
	if list.Size() != 3 {
		t.FailNow()
	}
}

func TestIfCreatingAnEmptyListWorks(t *testing.T) {
	list := LinkedList{}
	if list.Size() != 0 {
		t.FailNow()
	}
}

func TestIfIsEmptyWorksCorrectly(t *testing.T) {
	list := LinkedList{}
	if !list.IsEmpty() {
		t.FailNow()
	}
}

func TestIfClearingAListWorks(t *testing.T) {
	nodeOne := CreateNode(1)
	list := LinkedList{
		Head: &nodeOne,
	}
	nodeTwo := CreateNode(2)
	nodeThree := CreateNode(3)
	list.Append(nodeTwo)
	list.Append(nodeThree)
	if list.Size() != 3 {
		t.FailNow()
	}
	if list.IsEmpty() {
		t.FailNow()
	}
	list.Clear()
	if !list.IsEmpty() {
		t.FailNow()
	}
}

func TestIfAddAtIndexWorksCorrectly(t *testing.T) {
	nodeOne := CreateNode(1)
	list := LinkedList{
		Head: &nodeOne,
	}
	nodeTwo := CreateNode(2)
	nodeThree := CreateNode(3)
	list.Append(nodeTwo)
	list.Append(nodeThree)
	list.AddAt(2, 4)
	if list.Size() != 4 {
		t.FailNow()
	}
	if nodeOne.Next.Data != 4 {
		t.FailNow()
	}
	if nodeOne.Next.Next.Data != 2 {
		t.FailNow()
	}
}

func TestIfAppendToAnEmptyListWorks(t *testing.T) {
	list := LinkedList{}
	list.Append(CreateNode(1))
	if list.Size() != 1 {
		t.FailNow()
	}
}
