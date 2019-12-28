package datastructures

import "testing"

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
