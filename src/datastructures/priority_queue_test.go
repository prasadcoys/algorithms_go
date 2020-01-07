package datastructures

import "testing"

func TestIfAddingOneElementToPriorityQueueWorks(t *testing.T) {
	pq := PriorityQueue{}
	pq.Add(2)
	if pq.Poll() != 2 {
		t.FailNow()
	}

}
