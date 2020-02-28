package datastructures

import "testing"

func TestIfEntryCanBeAddedToQueue(t *testing.T) {
	nodeQueue := NodeQueue{}
	nodeQueue.Enqueue(&BSTNode{data: 1})
	if len(nodeQueue.dataqueue) != 1 {
		t.Fail()
	}
	nodeQueue.Enqueue(&BSTNode{data: 2})
	if len(nodeQueue.dataqueue) != 2 {
		t.Fail()
	}
}

func TestIfDequeueFromQueueWorksCorrectly(t *testing.T) {
	nodeQueue := NodeQueue{}
	nodeQueue.Enqueue(&BSTNode{data: 1})
	if len(nodeQueue.dataqueue) != 1 {
		t.Fail()
	}
	nodeQueue.Enqueue(&BSTNode{data: 2})
	if len(nodeQueue.dataqueue) != 2 {
		t.Fail()
	}
	if nodeQueue.Dequeue().data != 1 {
		t.Fail()
	}
	if nodeQueue.Dequeue().data != 2 {
		t.Fail()
	}
}
