package tests

import (
	"proj/models"
	"testing"
)

func TestStack_Push_WhenStackIsFull(t *testing.T) {
	stack := models.NewStack(2)
	stack.Push(1)
	stack.Push(2)
	if stack.Push(3) != false {
		t.Error("Push method is incorrect")
	}
}

func TestStack_Pop_WhenStackIsEmpty(t *testing.T) {
	stack := models.NewStack(2)
	_, ok := stack.Pop()
	if ok != false {
		t.Error("Pop method is incorrect")
	}
}

func TestStack_Push(t *testing.T) {
	s := models.NewStack(2)

	if !s.Push(1) {
		t.Error("Error pushing item to stack")
	}
	if s.GetTop() != 0 {
		t.Errorf("Expected top of stack to be 0, but got %d", s.GetTop())
	}
	if s.Display()[0] != 1 {
		t.Errorf("Expected data at top of stack to be 5, but got %d", s.Display()[0])
	}
	if !s.Push(2) {
		t.Error("Error pushing item to stack")
	}
	if s.Push(3) {
		t.Error("Error: stack should be full")
	}
}

func TestStack_Pop(t *testing.T) {
	s := models.NewStack(10)
	s.Push(1)
	s.Push(2)

	item, ok := s.Pop()
	if !ok || item != 2 {
		t.Error("Error popping item from stack")
	}
	item, ok = s.Pop()
	if !ok || item != 1 {
		t.Error("Error popping item from stack")
	}
	_, ok = s.Pop()
	if ok {
		t.Error("Error: stack should be empty")
	}
}

func TestStack_Top(t *testing.T) {
	s := models.NewStack(10)
	s.Push(1)

	item, ok := s.Top()
	if !ok || item != 1 {
		t.Error("Error getting top item from stack")
	}
	_, ok = s.Pop()
	if !ok {
		t.Error("Error: stack should be empty")
	}
}

func TestStack_Display(t *testing.T) {
	s := models.NewStack(10)
	s.Push(5)
	// Test displaying items in a stack with items
	data := s.Display()
	if len(data) != 1 {
		t.Errorf("Expected data slice to have 1 item, but it has %d", len(data))
	}
	if data[0] != 5 {
		t.Errorf("Expected data at top of stack to be 5, but got %d", data[0])
	}

	// Test displaying items in an empty stack
	s.Pop()
	data = s.Display()
	if data != nil {
		t.Error("Expected data to be nil, but it is not")
	}
}
