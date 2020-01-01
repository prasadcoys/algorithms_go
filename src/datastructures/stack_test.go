package datastructures

import (
	"testing"
)

func TestIfEmptyStackIsCreated(t *testing.T) {
	stack := Stack{}
	if stack.Size() != 0 {
		t.FailNow()
	}
}

func TestIfPushToEmptyStackWorks(t *testing.T) {
	stack := Stack{}
	stack.Push(2)
	if stack.Size() != 1 {
		t.FailNow()
	}
}

func TestIfPopFromStackWorks(t *testing.T) {
	stack := Stack{}
	stack.Push(2)
	topData, _ := stack.Pop()
	if topData != 2 {
		t.FailNow()
	}
	topData, _ = stack.Pop()
	if topData != -1 {
		t.FailNow()
	}
}

func TestIfPushAndPopOfMultipleElementsWorks(t *testing.T) {
	stack := Stack{}
	stack.Push(2)
	stack.Push(3)
	stack.Push(5)
	top, _ := stack.Pop()
	if top != 5 {
		t.FailNow()
	}
	top, _ = stack.Pop()
	if top != 3 {
		t.FailNow()
	}
	top, _ = stack.Pop()
	if top != 2 {
		t.FailNow()
	}
}

func TestIsEmpty(t *testing.T) {
	stack := Stack{}
	if !stack.IsEmpty() {
		t.FailNow()
	}
	stack.Push(2)
	if stack.IsEmpty() {
		t.FailNow()
	}
}

func TestPeek(t *testing.T) {
	stack := Stack{}
	stack.Push(2)
	stack.Push(3)
	if stack.Peek() != 3 {
		t.FailNow()
	}
}
