package datastructures

import (
	"errors"
)

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
	if s.size == 0 {
		return -1, errors.New("Empty stack.Nothing left to pop")
	}

	topData := s.Head.Data
	err := s.data.RemoveFirst()
	if s.data.Size() == 0 {
		s.Head = nil
	}
	if err == nil {
		s.Head = s.data.Head
		s.size--
	}

	return topData, err
}

func (s *Stack) IsEmpty() bool {
	return s.data.IsEmpty()
}

func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.Head.Data
}
