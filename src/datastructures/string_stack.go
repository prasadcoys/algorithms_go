package datastructures

type StringStack struct {
	data []string
}

func (s *StringStack) push(text string) {
	s.data = append(s.data, text)
}

func (s *StringStack) pop() string {
	pop := s.data[len(s.data)-1:][0]
	s.data = s.data[0 : len(s.data)-1]
	return pop
}

func (s *StringStack) peek() string {
	pop := s.data[len(s.data)-1:][0]
	return pop
}

func (s *StringStack) isEmpty() bool {
	return len(s.data) == 0
}
