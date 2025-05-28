package task32

type stack []int

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(v int) {
	*s = append(*s, v)
}

func (s *stack) pop() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

func (s *stack) peek() (int, bool) {
	if s.isEmpty() {
		return 0, false
	}
	return (*s)[len(*s)-1], true
}
