package datastructures

type PriorityQueue struct {
	data []int
}

func (p *PriorityQueue) Add(element int) {
	p.data = append(p.data, element)
}

func (p *PriorityQueue) Poll() int {
	if p.IsEmpty() {
		panic("Queue is empty")
	}
	return p.data[0]
}

func (p *PriorityQueue) IsEmpty() bool {
	return len(p.data) == 0
}
