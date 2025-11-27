package structures

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{data: []string{}}
}

func (s *Stack) Push(val string) {
	s.data = append(s.data, val)
}

func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	index := len(s.data) - 1
	val := s.data[index]
	s.data = s.data[:index]
	return val
}
