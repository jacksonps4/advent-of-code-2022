package main

import "testing"

func TestPushAndPop(t *testing.T) {
	stack := NewStack()
	stack.Push("one")
	stack.Push("two")

	two := stack.Pop()
	one := stack.Pop()

	if two != "two" {
		t.Errorf("Expected 'two' but was '%s'", two)
	}
	if one != "one" {
		t.Errorf("Expected 'one' but was '%s'", one)
	}
}

func TestPopEmpty(t *testing.T) {
	stack := NewStack()

	empty := stack.Pop()
	if empty != "" {
		t.Errorf("Expected \"\" but was %s", empty)
	}
}
