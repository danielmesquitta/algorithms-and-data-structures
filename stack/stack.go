package stack

import "fmt"

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty stack")
	}
	return s.elements[len(s.elements)-1], nil
}

func (s *Stack[T]) Pop() (T, error) {
	element, err := s.Peek()
	if err != nil {
		var zero T
		return zero, err
	}
	s.elements = s.elements[:len(s.elements)-1]
	return element, nil
}
