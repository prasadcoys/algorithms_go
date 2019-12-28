package datastructures

import (
	"testing"
)

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
	list.AddAt(3, 4)
	if list.Size() != 4 {
		t.FailNow()
	}
	if nodeOne.Next.Data != 2 {
		t.FailNow()
	}
	if nodeOne.Next.Next.Data != 4 {
		t.FailNow()
	}
	if nodeOne.Next.Next.Next.Data != 3 {
		t.FailNow()
	}
}

func TestIfAddAtOutofBoundsIndex_lower_ThrowsError(t *testing.T) {
	nodeOne := CreateNode(1)
	list := LinkedList{
		Head: &nodeOne,
	}
	err := list.AddAt(-1, 2)
	if err == nil {
		t.FailNow()
	}
}

func TestIfAddAtOutofBoundsIndex_upper_ThrowsError(t *testing.T) {
	nodeOne := CreateNode(1)
	list := LinkedList{
		Head: &nodeOne,
	}
	err := list.AddAt(2, 2)
	if err == nil {
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

func TestIfPeekFirstWorks(t *testing.T) {
	node := CreateNode(2)
	list := LinkedList{}
	list.Append(node)
	first, _ := list.PeekFirst()
	if first != 2 {
		t.FailNow()
	}
}
func TestIfPeekFirstWorksWhenHeadIsNotSet(t *testing.T) {
	list := LinkedList{}
	_, err := list.PeekFirst()
	if err == nil {
		t.FailNow()
	}
}

func TestIfRemoveAtSomeNonBoundaryLocationWorks(t *testing.T) {
	list := LinkedList{}
	nodeOne := CreateNode(1)
	nodeTwo := CreateNode(2)
	nodeThree := CreateNode(3)
	nodeFour := CreateNode(4)
	list.Append(nodeOne)
	list.Append(nodeTwo)
	list.Append(nodeThree)
	list.Append(nodeFour)
	list.RemoveAt(3)
	if list.Size() != 3 {
		t.FailNow()
	}
	if list.Head.Next.Next.Data != 4 {
		t.FailNow()
	}
}
func TestIfRemoveAtFirstWorks(t *testing.T) {
	list := LinkedList{}
	nodeOne := CreateNode(1)
	nodeTwo := CreateNode(2)
	nodeThree := CreateNode(3)
	nodeFour := CreateNode(4)
	list.Append(nodeOne)
	list.Append(nodeTwo)
	list.Append(nodeThree)
	list.Append(nodeFour)
	list.RemoveFirst()
	if list.Size() != 3 {
		t.FailNow()
	}
	if list.Head.Data != 2 {
		t.FailNow()
	}
}

func TestIfRemoveAtFirstWorksOnEmptyList(t *testing.T) {
	list := LinkedList{}
	err := list.RemoveFirst()
	if err == nil {
		t.FailNow()
	}
}

func TestIfRemovalAtEndWorks(t *testing.T) {
	list := LinkedList{}
	nodeOne := CreateNode(1)
	nodeTwo := CreateNode(2)
	nodeThree := CreateNode(3)
	nodeFour := CreateNode(4)
	list.Append(nodeOne)
	list.Append(nodeTwo)
	list.Append(nodeThree)
	list.Append(nodeFour)
	list.RemoveLast()
	if list.Size() != 3 {
		t.FailNow()
	}
	if list.Head.Next.Next.Next != nil {
		t.FailNow()
	}
}

func TestIfRemoveAtEndWorksOnEmptyList(t *testing.T) {
	list := LinkedList{}
	err := list.RemoveLast()
	if err == nil {
		t.FailNow()
	}
}
