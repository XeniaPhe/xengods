package stack

import (
	"errors"
	"fmt"
	"strings"
)

var emptyStackError error

func init() {
	emptyStackError = errors.New("empty stack")
}

type Stack[T any] []T

func New[T any](capacity ...int) Stack[T] {
	initialCapacity := 0

	if len(capacity) > 0 {
		initialCapacity = capacity[0]
	}

	return make(Stack[T], 0, initialCapacity)
}

func Of[T any](values ...T) Stack[T] {
	length := len(values)
	stack := make(Stack[T], 0, length)

	for _, val := range values {
		stack.Push(val)
	}

	return stack
}

func (s *Stack[T]) Push(value T) {
	*s = append(*s, value)
}

func (s *Stack[T]) Pop() (T, error) {
	l := len(*s)

	if l == 0 {
		return *new(T), emptyStackError
	}

	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res, nil
}

func (s Stack[T]) Peek() (T, error) {
	l := len(s)

	if l == 0 {
		return *new(T), emptyStackError
	}

	res := s[l-1]
	return res, nil
}

func (s *Stack[T]) Clear() {
	*s = (*s)[:0]
}

func (s Stack[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack[T]) Size() int {
	return len(s)
}

func (s Stack[T]) Capacity() int {
	return cap(s)
}

func (s Stack[T]) String() string {
	var builder strings.Builder
	builder.WriteString("Stack[")

	for i, val := range s {
		if i > 0 {
			builder.WriteString(", ")
		}
		
		builder.WriteString(fmt.Sprintf("%v", val))
	}

	builder.WriteString("]")
	return builder.String()
}