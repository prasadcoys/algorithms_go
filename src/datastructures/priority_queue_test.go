package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	expected := []int{1, 2}
	assert.Equal(t, expected, pq.heap)

}

func TestIfAddingElementThatGoesDirectlyToTheTopWorks(t *testing.T) {
	pq := PriorityQueue{}
	pq.heap = []int{6, 7, 12, 10, 15, 17}
	pq.Add(5)
	expected := []int{5, 7, 6, 10, 15, 17, 12}
	assert.Equal(t, expected, pq.heap)
}

func TestIfRemoveElementFromHeapWorks(t *testing.T) {
	pq := PriorityQueue{}
	pq.heap = []int{1, 2, 4, 5, 3, 6}
	top := pq.Poll()
	assert.Equal(t, 1, top)
	assert.Equal(t, []int{2, 3, 4, 5, 6}, pq.heap)
}
