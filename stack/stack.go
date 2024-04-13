package stack

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(item T) {
	s.elements = append(s.elements, item)
}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, fmt.Errorf("empty stack")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.elements) == 0 {
		var zero T
		return zero, fmt.Errorf("empty stack")
	}
	item := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return item, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}
