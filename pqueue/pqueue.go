package pqueue

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

var emptyPQError error

func init() {
	emptyPQError = errors.New("empty priority queue")
}

func minCompare(lhs int32, rhs int32) bool { return lhs < rhs }
func maxCompare(lhs int32, rhs int32) bool { return lhs > rhs }

type Item[T any] struct {
	Value T
	Priority int32
}

type PQueue[T any] struct {
	heap []Item[T]
	compare func(int32, int32) bool
}

func NewMinPQueue[T any](capacity ...int) PQueue[T] {
	initialCapacity := 0

	if len(capacity) > 0 {
		initialCapacity = capacity[0]
	}

	return PQueue[T]{make([]Item[T], 0, initialCapacity), minCompare}
}

func NewMaxPQueue[T any](capacity ...int) PQueue[T] {
	initialCapacity := 0

	if len(capacity) > 0 {
		initialCapacity = capacity[0]
	}

	return PQueue[T]{make([]Item[T], 0, initialCapacity), maxCompare}
}

func NewPQueue[T any](min bool, capacity ...int) PQueue[T] {
	if min {
		return NewMinPQueue[T](capacity...)
	}

	return NewMaxPQueue[T](capacity...)
}

func (p PQueue[T]) IsInitialized() bool {
	return p.compare != nil
}

func (p *PQueue[T]) InitializeIfNot(min bool) {
	p.heap = make([]Item[T], 0)

	if min {
		p.compare = minCompare
	} else {
		p.compare = maxCompare
	}
}

func (p PQueue[T]) Clone() PQueue[T] {
	return PQueue[T]{slices.Clone(p.heap), p.compare}
}

func (p *PQueue[T]) Enqueue(value T, priority int32) {
	p.heap = append(p.heap, Item[T]{value, priority})
	p.heapifyUp()
}

func (p *PQueue[T]) Dequeue() (Item[T], error) {
	size := len(p.heap)
	
	if size == 0 {
		return Item[T]{}, emptyPQError
	}

	var res Item[T]
	res, p.heap[0] = p.heap[0], p.heap[size-1]
	p.heap = p.heap[:size-1]
	p.heapifyDown()
	return res, nil
}

func (p PQueue[T]) Peek() (Item[T], error) {
	if len(p.heap) == 0 {
		return Item[T]{}, emptyPQError
	}

	return p.heap[0], nil
}

func (p PQueue[T]) GetSlice() []Item[T] {
	return p.heap
}

func (p *PQueue[T]) GetSlicePtr() *[]Item[T] {
	return &p.heap
}

func (p *PQueue[T]) Clear() {
	p.heap = p.heap[:0]
}

func (p PQueue[T]) IsEmpty() bool {
	return len(p.heap) == 0
}

func (p PQueue[T]) Size() int {
	return len(p.heap)
}

func (p PQueue[T]) Capacity() int {
	return cap(p.heap)
}

func (p PQueue[T]) String() string {
	var builder strings.Builder
	if p.compare(0, 1) {
		builder.WriteString("Min")
	} else {
		builder.WriteString("Max")
	}

	builder.WriteString("PQueue[")

	clone := p.Clone()
	first := true

	for !clone.IsEmpty() {
		if !first {
			builder.WriteString(", ")
		}

		item, err := clone.Dequeue()
		
		if err != nil {
			panic(err)
		}
		
		builder.WriteString(fmt.Sprintf("(%v:%d)", item.Value, item.Priority))
		first = false
	}

	builder.WriteString("]")
	return builder.String()
}

func (p PQueue[T]) heapifyUp() {
	curr := len(p.heap) - 1

	for {
		if curr <= 0 {
			return
		}

		parent := (curr - 1) >> 1
		currPriority := p.heap[curr].Priority
		parentPriority := p.heap[parent].Priority

		if p.compare(parentPriority, currPriority) ||  parentPriority == currPriority {
			break
		}

		p.heap[curr], p.heap[parent] = p.heap[parent], p.heap[curr]
		curr = parent
	}
}

func (p PQueue[T]) heapifyDown() {
	size := len(p.heap)
	curr := 0
	opt := curr

	for {
		left := (curr << 1) + 1
		right := left + 1

		if left < size && p.compare(p.heap[left].Priority, p.heap[opt].Priority) {
			opt = left
		}

		if right < size && p.compare(p.heap[right].Priority, p.heap[opt].Priority) {
			opt = right
		}

		if curr == opt {
			break
		}

		p.heap[curr], p.heap[opt] = p.heap[opt], p.heap[curr]
		curr = opt
	}
}