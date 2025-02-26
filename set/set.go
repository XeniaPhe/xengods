package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	set map[T]struct{}
}

func New[T comparable]() Set[T] {
	return Set[T]{make(map[T]struct{})}
}

func Of[T comparable](values ...T) Set[T] {
	set := Set[T]{make(map[T]struct{})}

	for _, val := range values {
		set.Add(val)
	}

	return set
}

func (s Set[T]) IsInitialized() bool {
	return s.set != nil
}

func (s *Set[T]) InitializeIfNot() {
	if s.set == nil {
		s.set = make(map[T]struct{})
	}
}

func (s *Set[T]) Clear() {
	s.set = make(map[T]struct{})
}

func (s Set[T]) Clone() Set[T] {
	clone := Set[T]{make(map[T]struct{})}

	for val := range s.set {
		clone.set[val] = struct{}{}
	}

	return clone
}

func orderBySize[T comparable](lhs Set[T], rhs Set[T]) (Set[T], Set[T]) {
	if len(lhs.set) <= len(rhs.set) {
		return lhs, rhs
	}

	return rhs, lhs
}

func (s Set[T]) GetRawSet() map[T]struct{} {
	return s.set
}

func (s Set[T]) Add(value T) {
	s.set[value] = struct{}{}
}

func (s Set[T]) Remove(value T) {
	delete(s.set, value)
}

func (s Set[T]) PopOne() T {
	for val := range s.set {
		delete(s.set, val)
		return val
	}

	var zero T
	return zero
}

func (s Set[T]) Contains(value T) bool {
	_, found := s.set[value]
	return found
}

func (s Set[T]) Size() int {
	return len(s.set)
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	union := Set[T]{make(map[T]struct{})}

	for val := range s.set {
		union.set[val] = struct{}{}
	}

	for val := range other.set {
		union.set[val] = struct{}{}
	}

	return union
}

func (s Set[T]) UnionWith(other Set[T]) {
	for val := range other.set {
		s.set[val] = struct{}{}
	}
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	smaller, bigger := orderBySize(s, other)
	intersection := Set[T]{make(map[T]struct{})}

	for val := range smaller.set {
		if bigger.Contains(val) {
			intersection.set[val] = struct{}{}
		}
	}

	return intersection
}

func (s Set[T]) IntersectWith(other Set[T]) {
	marked := make([]T, 0)

	for val := range s.set {
		if !other.Contains(val) {
			marked = append(marked, val)
		}
	}

	for _, val := range marked {
		delete(s.set, val)
	}
}

func (s Set[T]) Except(other Set[T]) Set[T] {
	except := Set[T]{make(map[T]struct{})}

	for val := range s.set {
		if !other.Contains(val) {
			except.set[val] = struct{}{}
		}
	}

	return except
}

func (s Set[T]) ExceptWith(other Set[T]) {
	marked := make([]T, 0)

	for val := range s.set {
		if other.Contains(val) {
			marked = append(marked, val)
		}
	}

	for _, val := range marked {
		delete(s.set, val)
	}
}

func (s Set[T]) Overlaps(other Set[T]) bool {
	smaller, bigger := orderBySize(s, other)

	for val := range smaller.set {
		if bigger.Contains(val) {
			return true
		}
	}

	return false
}

func (s Set[T]) IsSubsetOf(other Set[T]) bool {
	for val := range s.set {
		if !other.Contains(val) {
			return false
		}
	}

	return true
}

func (s Set[T]) IsProperSubsetOf(other Set[T]) bool {
	return s.IsSubsetOf(other) && len(s.set) < len(other.set)
}

func (s Set[T]) IsSupersetOf(other Set[T]) bool {
	return other.IsSubsetOf(s)
}

func (s Set[T]) IsProperSupersetOf(other Set[T]) bool {
	return other.IsProperSubsetOf(s)
}

func (s Set[T]) String() string {
	var builder strings.Builder
	builder.WriteString("Set{")
	first := true

	for val := range s.set {
		if first {
			first = false
			builder.WriteString(", ")
		}
		
		builder.WriteString(fmt.Sprintf("%v", val))
	}

	builder.WriteString("}")
	return builder.String()
}