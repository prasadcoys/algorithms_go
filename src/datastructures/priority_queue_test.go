package datastructures

import "testing"

func TestIfAddingOneElementToPriorityQueueWorks(t *testing.T) {
	pq := PriorityQueue{}
	pq.Add(2)
	if pq.Poll() != 2 {
<<<<<<< HEAD
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
=======
		t.FailNow()
	}

>>>>>>> f9ac7cd1f4485a7593925df6c09e79513008fa68
}
