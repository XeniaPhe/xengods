package pqueue

import (
	"testing"
)

func TestPQConstructors(t *testing.T) {
	minpq := NewMinPQueue[int]()

	if minpq.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", minpq.Size())
	}

	if minpq.Capacity() != 0 {
		t.Errorf("Expected capacity 0, got %d instead", minpq.Capacity())
	}

	if !minpq.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	maxpq := NewMaxPQueue[int]()

	if maxpq.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", maxpq.Size())
	}

	if maxpq.Capacity() != 0 {
		t.Errorf("Expected capacity 0, got %d instead", maxpq.Capacity())
	}

	if !maxpq.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	minpq = NewPQueue[int](true)

	if minpq.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", minpq.Size())
	}

	if minpq.Capacity() != 0 {
		t.Errorf("Expected capacity 0, got %d instead", minpq.Capacity())
	}

	if !minpq.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	maxpq = NewPQueue[int](false)

	if maxpq.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", maxpq.Size())
	}

	if maxpq.Capacity() != 0 {
		t.Errorf("Expected capacity 0, got %d instead", maxpq.Capacity())
	}

	if !maxpq.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	minpq = NewMinPQueue[int](100)

	if minpq.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", minpq.Size())
	}

	if minpq.Capacity() != 100 {
		t.Errorf("Expected capacity 100, got %d instead", minpq.Capacity())
	}

	if !minpq.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	maxpq = NewMaxPQueue[int](100)

	if maxpq.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", maxpq.Size())
	}

	if maxpq.Capacity() != 100 {
		t.Errorf("Expected capacity 100, got %d instead", maxpq.Capacity())
	}

	if !maxpq.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	minpq = NewMinPQueue[int](5)
	minpq.Enqueue(5, 5)
	minpq.Enqueue(6, 6)
	minpq.Enqueue(4, 4)

	clone := minpq.Clone()

	if clone.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", clone.Size())
	}

	if clone.Capacity() != 3 {
		t.Errorf("Expected capacity 3, got %d instead", clone.Capacity())
	}

	if !clone.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	maxpq = NewMinPQueue[int](5)
	maxpq.Enqueue(5, 5)
	maxpq.Enqueue(6, 6)
	maxpq.Enqueue(4, 4)

	clone = maxpq.Clone()

	if clone.Size() != 3 {
		t.Errorf("Expected size 3, got %d instead", clone.Size())
	}

	if clone.Capacity() != 3 {
		t.Errorf("Expected capacity 3, got %d instead", clone.Capacity())
	}

	if !clone.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	minpq = PQueue[int]{}

	if minpq.IsInitialized() {
		t.Error("Expected false, got true instead")
	}

	minpq.InitializeIfNot(true)

	if !minpq.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	if minpq.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", minpq.Size())
	}

	if minpq.Capacity() != 0 {
		t.Errorf("Expected capacity 0, got %d instead", minpq.Capacity())
	}

	maxpq = PQueue[int]{}

	if maxpq.IsInitialized() {
		t.Error("Expected false, got true instead")
	}

	maxpq.InitializeIfNot(true)

	if !maxpq.IsInitialized() {
		t.Error("Expected true, got false instead")
	}

	if maxpq.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", minpq.Size())
	}

	if maxpq.Capacity() != 0 {
		t.Errorf("Expected capacity 0, got %d instead", minpq.Capacity())
	}
}

func TestPQEnqueueDequeuePeek(t *testing.T) {
	min := NewMinPQueue[int]()
	min.Enqueue(10, 5)
	min.Enqueue(5, 2)
	min.Enqueue(6, 5)
	min.Enqueue(2, 10)

	item, err := min.Dequeue()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 2 {
		t.Errorf("Expected priority 2, got %d instead", item.Priority)
	}

	if item.Value != 5 {
		t.Errorf("Expected value 5, got %d instead", item.Value)
	}

	item, err = min.Peek()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 5 {
		t.Errorf("Expected priority 5, got %d instead", item.Priority)
	}

	if item.Value != 10 {
		t.Errorf("Expected value 10, got %d instead", item.Value)
	}

	item, err = min.Dequeue()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 5 {
		t.Errorf("Expected priority 5, got %d instead", item.Priority)
	}

	if item.Value != 10 {
		t.Errorf("Expected value 10, got %d instead", item.Value)
	}

	item, err = min.Dequeue()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 5 {
		t.Errorf("Expected priority 5, got %d instead", item.Priority)
	}

	if item.Value != 6 {
		t.Errorf("Expected value 6, got %d instead", item.Value)
	}

	item, err = min.Dequeue()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 10 {
		t.Errorf("Expected priority 10, got %d instead", item.Priority)
	}

	if item.Value != 2 {
		t.Errorf("Expected value 2, got %d instead", item.Value)
	}

	item, err = min.Dequeue()

	if err == nil {
		t.Error("Expected an error, got nothing")
	}

	item, err = min.Peek()

	if err == nil {
		t.Error("Expected an error, got nothing")
	}

	max := NewMaxPQueue[int]()
	max.Enqueue(10, 5)
	max.Enqueue(5, 2)
	max.Enqueue(6, 5)
	max.Enqueue(2, 10)

	item, err = max.Dequeue()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 10 {
		t.Errorf("Expected priority 10, got %d instead", item.Priority)
	}

	if item.Value != 2 {
		t.Errorf("Expected value 2, got %d instead", item.Value)
	}

	item, err = max.Peek()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 5 {
		t.Errorf("Expected priority 5, got %d instead", item.Priority)
	}

	if item.Value != 10 {
		t.Errorf("Expected value 10, got %d instead", item.Value)
	}

	item, err = max.Dequeue()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 5 {
		t.Errorf("Expected priority 5, got %d instead", item.Priority)
	}

	if item.Value != 10 {
		t.Errorf("Expected value 10, got %d instead", item.Value)
	}

	item, err = max.Dequeue()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 5 {
		t.Errorf("Expected priority 5, got %d instead", item.Priority)
	}

	if item.Value != 6 {
		t.Errorf("Expected value 6, got %d instead", item.Value)
	}

	item, err = max.Dequeue()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if item.Priority != 2 {
		t.Errorf("Expected priority 2, got %d instead", item.Priority)
	}

	if item.Value != 5 {
		t.Errorf("Expected value 5, got %d instead", item.Value)
	}

	item, err = max.Dequeue()

	if err == nil {
		t.Error("Expected an error, got nothing")
	}

	item, err = max.Peek()

	if err == nil {
		t.Error("Expected an error, got nothing")
	}
}

func TestPQClear(t *testing.T) {
	pq := NewMinPQueue[int](100)

	for i := range 100 {
		pq.Enqueue(i, 0)
	}

	capacity := pq.Capacity()
	pq.Clear()

	if !pq.IsEmpty() {
		t.Errorf("Expected size 0, got %d instead", pq.Size())
	}

	if pq.Capacity() != capacity {
		t.Errorf("Expected capacity %d, got %d instead", capacity, pq.Capacity())
	}

	pq = NewMaxPQueue[int](100)

	for i := range 100 {
		pq.Enqueue(i, 0)
	}

	capacity = pq.Capacity()
	pq.Clear()

	if !pq.IsEmpty() {
		t.Errorf("Expected size 0, got %d instead", pq.Size())
	}

	if pq.Capacity() != capacity {
		t.Errorf("Expected capacity %d, got %d instead", capacity, pq.Capacity())
	}
}

func TestPQString(t *testing.T) {
	pq := NewMinPQueue[int]()
	pq.Enqueue(1, 2)
	pq.Enqueue(3, 4)
	pq.Enqueue(5, 6)
	pq.Enqueue(7, 8)
	pq.Enqueue(9, 0)

	str := pq.String()

	if str != "MinPQueue[(9:0), (1:2), (3:4), (5:6), (7:8)]" {
		t.Errorf("Expected 'MinPQueue[(9:0), (1:2), (3:4), (5:6), (7:8)]', got '%s' instead", str)
	}

	pq = NewMaxPQueue[int]()
	pq.Enqueue(1, 2)
	pq.Enqueue(3, 4)
	pq.Enqueue(5, 6)
	pq.Enqueue(7, 8)
	pq.Enqueue(9, 0)

	str = pq.String()

	if str != "MaxPQueue[(7:8), (5:6), (3:4), (1:2), (9:0)]" {
		t.Errorf("Expected 'MaxPQueue[(7:8), (5:6), (3:4), (1:2), (9:0)]', got '%s' instead", str)
	}
}