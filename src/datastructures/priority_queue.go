package datastructures

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
	p.heap = p.heap[1:]
	return top
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.heap) == 0
}
