package datastructures

import (
	"fmt"
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
	fmt.Println(topData)
}
