package datastructures

import (
	"fmt"
	"testing"
)

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

func TestIfWeCanUseQueueAsAStack(t *testing.T) {
	nodeStack := NodeQueue{}
	nodeStack.Push(&BSTNode{data: 1})
	nodeStack.Push(&BSTNode{data: 2})
	nodeStack.Push(&BSTNode{data: 3})
	if nodeStack.Pop().data != 3 {
		t.Fail()
	}
	data := nodeStack.Pop().data
	if data != 2 {
		fmt.Println(data)
		t.Fail()
	}
	if nodeStack.Pop().data != 1 {
		t.Fail()
	}
}
