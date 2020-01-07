package datastructures

import "testing"

func TestIfAddingOneElementToPriorityQueueWorks(t *testing.T) {
	pq := PriorityQueue{}
	pq.Add(2)
	if pq.Poll() != 2 {
		t.Fail()
	}
	if !pq.IsEmpty() {
		t.Fail()
	}
}

func TestIfAddingTwoElementsReturnsTheElementsInPriorityOrder(t *testing.T) {
	pq := PriorityQueue{}
	pq.Add(2)
	pq.Add(1)
	if pq.Poll() != 1 {
		t.Fail()
	}
	if pq.Poll() != 2 {
		t.Fail()
	}
}
