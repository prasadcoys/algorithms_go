package datastructures

import "testing"

func TestIfEmptyQueueIsCreated(t *testing.T) {
	queue := Queue{}
	if !queue.IsEmpty() {
		t.FailNow()
	}
}

func TestIfEnqueueToEmptyQueueWorks(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(1)
	if queue.IsEmpty() {
		t.FailNow()
	}
}

func TestIfEnqueueToOneElementQueueWorks(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	if queue.IsEmpty() {
		t.FailNow()
	}
	if queue.Size() != 2 {
		t.FailNow()
	}
}

func TestIfDequeueWorksOnTwoElementQueues(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	if queue.Size() != 3 {
		t.FailNow()
	}
	num := queue.Dequeue()
	if num != 2 {
		t.FailNow()
	}
	num = queue.Dequeue()
	if num != 3 {
		t.FailNow()
	}
	num = queue.Dequeue()
	if num != 4 {
		t.FailNow()
	}
}

func TestIfPeekWorks(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	if queue.Peek() != 2 {
		t.FailNow()
	}
}
