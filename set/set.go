package set

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	set map[T]struct{}
}

func New[T comparable](size ...int) Set[T] {
	sizeHint := 0

	if len(size) > 0 {
		sizeHint = size[0]
	}

	return Set[T]{make(map[T]struct{}, sizeHint)}
}

func Of[T comparable](values ...T) Set[T] {
	set := New[T](len(values))

	for _, val := range values {
		set.Add(val)
	}

	return set
}

func FromKeys[K comparable, V any](hashmap map[K]V) Set[K] {
	set := New[K](len(hashmap))

	for key := range hashmap {
		set.set[key] = struct{}{}
	}

	return set
}

func FromValues[K comparable, V comparable](hashmap map[K]V) Set[V] {
	set := New[V](len(hashmap))

	for _, val := range hashmap {
		set.set[val] = struct{}{}
	}

	return set
}

func (s *Set[T]) Clear() {
	s.set = make(map[T]struct{})
}

func (s Set[T]) Clone() Set[T] {
	clone := New[T](len(s.set))

	for val := range s.set {
		clone.set[val] = struct{}{}
	}

	return clone
}

func (s Set[T]) IsInitialized() bool {
	return s.set != nil
}

func (s *Set[T]) InitializeIfNot() {
	if s.set == nil {
		s.set = make(map[T]struct{})
	}
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

func (s Set[T]) ContainsSome(values ...T) bool {
	for _, val := range values {
		if _, found := s.set[val]; found {
			return true
		}
	}

	return false
}

func (s Set[T]) ContainsAll(values ...T) bool {
	for _, val := range values {
		if _, found := s.set[val]; !found {
			return false
		}
	}

	return true
}

func (s Set[T]) Size() int {
	return len(s.set)
}

func (s Set[T]) IsEmpty() bool {
	return len(s.set) == 0
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	smaller, bigger := orderBySize(s, other)
	sizeHint := len(bigger.set) + len(smaller.set) / 2
	union := New[T](sizeHint)

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
	sizeHint := len(smaller.set) / 2
	intersection := New[T](sizeHint)

	for val := range smaller.set {
		if bigger.Contains(val) {
			intersection.set[val] = struct{}{}
		}
	}

	return intersection
}

func (s Set[T]) IntersectWith(other Set[T]) {
	smaller, bigger := orderBySize(s, other)
	var sizeHint int

	if len(s.set) == len(smaller.set) {
		sizeHint = len(smaller.set) / 2
	} else {
		sizeHint = len(bigger.set) - len(smaller.set) / 2
	}

	marked := make([]T, 0, sizeHint)

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
	smaller, bigger := orderBySize(s, other)
	var sizeHint int

	if len(s.set) == len(smaller.set) {
		sizeHint = len(smaller.set) / 2
	} else {
		sizeHint = len(bigger.set) - len(smaller.set) / 2
	}

	except := New[T](sizeHint)

	for val := range s.set {
		if !other.Contains(val) {
			except.set[val] = struct{}{}
		}
	}

	return except
}

func (s Set[T]) ExceptWith(other Set[T]) {
	smaller, _ := orderBySize(s, other)
	sizeHint := len(smaller.set) / 2
	marked := make([]T, 0, sizeHint)

	for val := range s.set {
		if other.Contains(val) {
			marked = append(marked, val)
		}
	}

	for _, val := range marked {
		delete(s.set, val)
	}
}

func (s Set[T]) SymmetricExcept(other Set[T]) Set[T] {
	_, bigger := orderBySize(s, other)
	sizeHint := len(bigger.set)
	symmetricExcept := New[T](sizeHint)

	for val := range s.set {
		if !other.Contains(val) {
			symmetricExcept.set[val] = struct{}{}
		}
	}

	for val := range other.set {
		if !s.Contains(val) {
			symmetricExcept.set[val] = struct{}{}
		}
	}

	return symmetricExcept
}

func (s Set[T]) SymmetricExceptWith(other Set[T]) {
	for val := range other.set {
		if _, exists := s.set[val]; exists {
			delete(s.set, val)
		} else {
			s.set[val] = struct{}{}
		}
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

func (s Set[T]) SetEquals(other Set[T]) bool {
	if len(s.set) != len(other.set) {
		return false
	}

	for val := range s.set {
		if !other.Contains(val) {
			return false
		}
	}

	return true
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

func (s Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(s.set))

	for val := range s.set {
		slice = append(slice, val)
	}

	return slice
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

func orderBySize[T comparable](lhs Set[T], rhs Set[T]) (Set[T], Set[T]) {
	if len(lhs.set) <= len(rhs.set) {
		return lhs, rhs
	}

	return rhs, lhs
}