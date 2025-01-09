package collections

type Stack[T any] struct {
	stack []T
}

func (s *Stack[T]) Push(t T) {
	s.stack = append(s.stack, t)
}

func (s *Stack[T]) Pop() T {
	i := len(s.stack) - 1
	t := s.stack[i]
	s.stack = s.stack[:i]
	return t
}

func (s *Stack[T]) Peek() T {
	return s.stack[len(s.stack)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.stack)
}
