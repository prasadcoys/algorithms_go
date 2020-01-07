package datastructures

type PriorityQueue struct {
	heap []int
}

func (p *PriorityQueue) Add(element int) {
	p.heap = append(p.heap, element)
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
