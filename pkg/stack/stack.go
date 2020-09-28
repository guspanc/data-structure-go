package stack

// Stack represents a stack struct
type Stack struct {
	items []interface{}
}

// Push will add a value at the end (top)
func (s *Stack) Push(i interface{}) {
	s.items = append(s.items, i)
}

// Pop will remove the value at the end (top)
func (s *Stack) Pop() interface{} {
	lastIndex := len(s.items) - 1
	toRemove := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return toRemove
}
