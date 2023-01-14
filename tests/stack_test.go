package tests

import (
	"proj/models"
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := models.NewStack(10)

	if !s.Push(1) {
		t.Error("Error pushing item to stack")
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
	if ok {
		t.Error("Error: stack should be empty")
	}
}
