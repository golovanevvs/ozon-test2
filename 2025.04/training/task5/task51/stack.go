package task51

type stack []vertex

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(v vertex) {
	*s = append(*s, v)
}

func (s *stack) pop() (vertex, bool) {
	if s.isEmpty() {
		return vertex{}, false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

// func (s *stack) peek() (vertex, bool) {
// 	if s.isEmpty() {
// 		return vertex{}, false
// 	}
// 	return (*s)[len(*s)-1], true
// }
