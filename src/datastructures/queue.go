package datastructures

type Queue struct {
	data DoublyLinkedList
}

func (q *Queue) IsEmpty() bool {
	return q.data.IsEmpty()
}

func (q *Queue) Enqueue(element int) {
	q.data.Append(element)
}

func (q *Queue) Size() int {
	return q.data.Size()
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic("Empty queue")
	}
	first := q.data.PeekFirst()
	q.data.RemoveFirst()
	return first
}

func (q *Queue) Peek() int {
	return q.data.PeekFirst()
}
