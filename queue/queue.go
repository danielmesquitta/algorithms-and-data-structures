package queue

import "fmt"

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, fmt.Errorf("empty queue")
	}
	return q.elements[0], nil
}

func (q *Queue[T]) Push(element T) {
	q.elements = append(q.elements, element)
}

func (q *Queue[T]) Pop() (T, error) {
	element, err := q.Peek()
	if err != nil {
		var zero T
		return zero, err
	}
	q.elements = q.elements[1:]
	return element, nil
}
