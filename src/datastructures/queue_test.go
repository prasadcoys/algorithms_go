package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestIfWeCanReverseFirstThreeElementsOfAQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	ReverseFirstKElements(&queue, 3)
	assert.Equal(t, "3 2 1 4 5", queue.data.ToString())

}

func TestIfWeCanReverseAllElementsOfAQueue(t *testing.T) {
	queue := Queue{}
	queue.Enqueue(4)
	queue.Enqueue(3)
	queue.Enqueue(2)
	queue.Enqueue(1)

	ReverseFirstKElements(&queue, 4)
	assert.Equal(t, "1 2 3 4", queue.data.ToString())

}
