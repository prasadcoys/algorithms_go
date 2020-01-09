package datastructures

import "fmt"

type PriorityQueue struct {
	heap []int
}

func (p *PriorityQueue) Add(element int) {
	currentIndex := len(p.heap)
	p.heap = append(p.heap, element)
	for {
		var parentIndex int
		parentIndex = (currentIndex - 1) / 2
		if p.heap[currentIndex] >= p.heap[parentIndex] {
			break
		}
		temp := p.heap[currentIndex]
		p.heap[currentIndex] = p.heap[parentIndex]
		p.heap[parentIndex] = temp
		currentIndex = parentIndex
	}
}

func (p *PriorityQueue) Poll() int {
	if p.IsEmpty() {
		panic("Queue is empty")
	}
	top := p.heap[0]
	//first swap with last element
	p.heap[0] = p.heap[len(p.heap)-1]
	if len(p.heap) == 1 {
		p.heap = []int{}
		return top
	}
	p.heap = p.heap[0 : len(p.heap)-1]

	currentIndex := 0
	for {
		fmt.Println(p.heap)
		if currentIndex >= len(p.heap)-1 {
			break
		}
		leftChildIndex := 2*currentIndex + 1
		rightChildIndex := 2*currentIndex + 2
		if p.heap[leftChildIndex] < p.heap[rightChildIndex] {
			if p.heap[leftChildIndex] < p.heap[currentIndex] {
				temp := p.heap[currentIndex]
				p.heap[currentIndex] = p.heap[leftChildIndex]
				p.heap[leftChildIndex] = temp
				currentIndex = leftChildIndex
				continue
			} else {
				break
			}
		} else {
			if p.heap[rightChildIndex] < p.heap[currentIndex] {
				temp := p.heap[currentIndex]
				p.heap[currentIndex] = p.heap[rightChildIndex]
				p.heap[rightChildIndex] = temp
				currentIndex = rightChildIndex
				continue
			} else {
				break
			}
		}

	}

	return top
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.heap) == 0
}
