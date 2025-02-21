package stack

import (
	"errors"
	"fmt"
	"strings"
)

type IStack[T any] interface {
	Push(value T)
	Pop() (T, error)
	Peek() (T, error)
	Clear()
	IsEmpty() bool
	Size() int
	Capacity() int
	fmt.Stringer
}

type stack[T any] struct {
	slice *[]T
}

func New[T any](capacity ...int) stack[T] {
	initialCapacity := 0
	if len(capacity) > 0 {
		initialCapacity = capacity[0]
	}

	slice := make([]T, 0, initialCapacity)
	return stack[T]{&slice}
}

func (s stack[T]) Push(value T) {
	*s.slice = append(*s.slice, value)
}

func (s stack[T]) Pop() (T, error) {
	l := len(*s.slice)

	if l == 0 {
		return *new(T), errors.New("empty stack")
	}

	res := (*s.slice)[l-1]
	*s.slice = (*s.slice)[:l-1]
	return res, nil
}

func (s stack[T]) Peek() (T, error) {
	l := len(*s.slice)

	if l == 0 {
		return *new(T), errors.New("empty stack")
	}

	res := (*s.slice)[l-1]
	return res, nil
}

func (s stack[T]) Clear() {
	*s.slice = (*s.slice)[:0]
}

func (s stack[T]) IsEmpty() bool {
	return len(*s.slice) == 0
}

func (s stack[T]) Size() int {
	return len(*s.slice)
}

func (s stack[T]) Capacity() int {
	return cap(*s.slice)
}

func (s stack[T]) String() string {
	var builder strings.Builder
	builder.WriteString("Stack[")

	for i, v := range *s.slice {
		if i > 0 {
			builder.WriteString(", ")
		}
		
		builder.WriteString(fmt.Sprintf("%v", v))
	}

	builder.WriteString("]")
	return builder.String()
}