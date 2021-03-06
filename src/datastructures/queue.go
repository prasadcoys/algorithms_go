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

func ReverseFirstKElements(queue *Queue, num int) {
	reversestack := Stack{}
	for i := 0; i < num; i++ {
		reversestack.Push((*queue).Dequeue())
	}
	for i := 0; i < num; i++ {
		element, _ := reversestack.Pop()
		(*queue).Enqueue(element)
	}
	for i := num; i < (*queue).Size(); i++ {
		(*queue).Enqueue((*queue).Dequeue())
	}
}

func ReverseQueue(queue *Queue) {
	stack := Stack{}
	for {
		if (*queue).IsEmpty() {
			break
		}
		num := (*queue).Dequeue()
		stack.Push(num)

	}
	for {
		if stack.IsEmpty() {
			break
		}
		num, _ := stack.Pop()
		(*queue).Enqueue(num)
	}
}
