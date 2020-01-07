package datastructures

type PriorityQueue struct {
<<<<<<< HEAD
	heap []int
}

func (p *PriorityQueue) Add(element int) {
	p.heap = append(p.heap, element)
=======
	data []int
}

func (p *PriorityQueue) Add(element int) {
	p.data = append(p.data, element)
>>>>>>> f9ac7cd1f4485a7593925df6c09e79513008fa68
}

func (p *PriorityQueue) Poll() int {
	if p.IsEmpty() {
		panic("Queue is empty")
	}
<<<<<<< HEAD
	top := p.heap[0]
	p.heap = p.heap[1:]
	return top
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.heap) == 0
=======
	return p.data[0]
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.data) == 0
>>>>>>> f9ac7cd1f4485a7593925df6c09e79513008fa68
}
