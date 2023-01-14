package models

type Stack struct {
	size int
	top  int
	data []int
}

func NewStack(size int) *Stack {
	return &Stack{
		size: size,
		top:  -1,
		data: make([]int, size),
	}
}

func (s *Stack) Push(item int) bool {
	if s.top == s.size-1 {
		return false
	}
	s.top++
	s.data[s.top] = item
	return true
}

func (s *Stack) Pop() (int, bool) {
	if s.top == -1 {
		return 0, false
	}
	item := s.data[s.top]
	s.top--
	return item, true
}

func (s *Stack) Top() (int, bool) {
	if s.top == -1 {
		return 0, false
	}
	return s.data[s.top], true
}

func (s *Stack) GetTop() int {
	return s.top
}

func (s *Stack) Display() []int {
	if s.top == -1 {
		return nil
	}
	return s.data[:s.top+1]
}

// new
func (s *Stack) NewStack(size int) {
	s.size = size
	s.top = -1
	s.data = make([]int, size)
}
