package datastructures

import "testing"
import "github.com/stretchr/testify/assert"

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
