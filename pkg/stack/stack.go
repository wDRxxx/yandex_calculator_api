package stack

type Stack[T any] struct {
	sl []T
}

// NewStack creates new stack with optional args: length, capacity
func NewStack[T any](args ...int) *Stack[T] {
	if len(args) == 1 {
		return &Stack[T]{
			sl: make([]T, args[0]),
		}
	}
	if len(args) == 2 {
		return &Stack[T]{
			sl: make([]T, args[0], args[1]),
		}
	}

	return &Stack[T]{}
}

func (s *Stack[T]) Push(v T) {
	s.sl = append(s.sl, v)
}

func (s *Stack[T]) Pop() T {
	if len(s.sl) == 0 {
		panic("stack is empty")
	}

	v := s.sl[len(s.sl)-1]
	s.sl = s.sl[:len(s.sl)-1]

	return v
}

func (s *Stack[T]) Peek() T {
	return s.sl[len(s.sl)-1]
}

func (s *Stack[T]) Len() int {
	return len(s.sl)
}
