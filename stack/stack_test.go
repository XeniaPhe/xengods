package stack

import (
	"testing"
)

func TestStackConstructors(t *testing.T) {
	s := New[int]()

	if s == nil {
		t.Fatal("Stack is nil")
	}

	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", s.Size())
	}

	if s.Capacity() != 0 {
		t.Errorf("Expected capacity 0, got %d instead", s.Capacity())
	}

	st := New[string](100)

	if st.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", st.Size())
	}

	if st.Capacity() != 100 {
		t.Errorf("Expected capacity 100, got %d instead", st.Capacity())
	}

	sta := Of[uint](1, 2, 3, 4, 5)

	if sta == nil {
		t.Fatal("Stack is nil")
	}

	if sta.Size() != 5 {
		t.Errorf("Expected size 5, got %d instead", sta.Size())
	}

	if sta.Capacity() != 5 {
		t.Errorf("Expected capacity 5, got %d instead", sta.Capacity())
	}

	if sta.String() != "Stack[1, 2, 3, 4, 5]" {
		t.Errorf("Expected '%s', got '%s' instead", "Stack[1, 2, 3, 4, 5]", sta.String())
	}
}

func TestStackPushPopPeek(t *testing.T) {
	s := New[int]()

	_, err := s.Peek()

	if err == nil {
		t.Errorf("Expected an error, got nothing")
	}

	_, err = s.Pop()

	if err == nil {
		t.Errorf("Expected an error, got nothing")
	}
	
	s.Push(1)

	if s.Size() != 1 {
		t.Fatalf("Expected size 1, got %d instead", s.Size())
	}

	val, err := s.Peek()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if val != 1 {
		t.Errorf("Expected value 1, got %d instead", val)
	}

	if s.Size() != 1 {
		t.Fatalf("Expected size 1, got %d instead", s.Size())
	}

	val, err = s.Pop()

	if err != nil {
		t.Errorf("Expected no error, got '%s' instead", err.Error())
	}

	if val != 1 {
		t.Errorf("Expected value 1, got %d instead", val)
	}

	if s.Size() != 0 {
		t.Fatalf("Expected size 0, got %d instead", s.Size())
	}

	for i := range 100 {
		s.Push(i)
	}

	if s.Size() != 100 {
		t.Fatalf("Expected size 100, got %d instead", s.Size())
	}

	for range 100 {
		s.Pop()
	}

	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", s.Size())
	}
}

func TestStackClear(t *testing.T) {
	s := New[int]()

	for i := range 100 {
		s.Push(i)
	}

	if s.Size() != 100 {
		t.Fatalf("Expected size 100, got %d instead", s.Size())
	}

	s.Clear()

	if s.Size() != 0 {
		t.Errorf("Expected size 0, got %d instead", s.Size())
	}
}

func TestStackString(t *testing.T) {
	s := New[int]()

	if s.String() != "Stack[]" {
		t.Errorf("Expected 'Stack[]', got '%s' instead", s.String())
	}

	s.Push(1)

	if s.String() != "Stack[1]" {
		t.Errorf("Expected 'Stack[1]', got %s instead", s.String())
	}

	s.Push(2)

	if s.String() != "Stack[1, 2]" {
		t.Errorf("Expected 'Stack[1, 2]', got %s instead", s.String())
	}
}

func TestStackIsEmpty(t *testing.T) {
	s := New[int]()

	if !s.IsEmpty() {
		t.Errorf("Expected true, got false instead")
	}

	s.Push(1)

	if s.IsEmpty() {
		t.Errorf("Expected false, got true instead")
	}
}