package datastructures

type Stack struct {
	Head *Node
	data LinkedList
	size int
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Push(element int) {
	if s.data.IsEmpty() {
		s.data = LinkedList{}
	}
	s.data.AddAt(1, element)
	s.Head = s.data.Head
	s.size++
}

func (s *Stack) Pop() (int, error) {
	topData := s.Head.Data
	err := s.data.RemoveFirst()
	if err == nil {
		s.size--
	}

	return topData, err
}
