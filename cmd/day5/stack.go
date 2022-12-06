package main

type stack struct {
	data []string
}

func NewStack() *stack {
	return &stack{
		data: make([]string, 0),
	}
}

func (s *stack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	pop := s.data[len(s.data)-1]
	s.data = s.data[0 : len(s.data)-1]
	return pop
}
