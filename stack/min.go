package stack

type MinStack struct {
	data, min []int
}

func (s *MinStack) Push(value int) {
	s.data = append(s.data, value)

	if len(s.min) == 0 || s.min[len(s.min)-1] > value {
		s.min = append(s.min, value)
		return
	}

	s.min = append(s.min, s.min[len(s.min)-1])
}

func (s *MinStack) Pop() int {
	if len(s.data) == 0 || len(s.min) == 0 {
		panic("stack is empty")
	}

	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.min = s.min[:len(s.min)-1]

	return value
}

func (s *MinStack) GetMin() int {
	if len(s.data) == 0 || len(s.min) == 0 {
		panic("stack is empty")
	}

	return s.min[len(s.min)-1]
}
