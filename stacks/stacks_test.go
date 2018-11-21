package stacks

import (
	"testing"
)

func TestStack(t *testing.T) {
	stringStack := &Stack{}
	stringStack.Push("First")
	stringStack.Push("Second")
	ex_v := stringStack.Pop()
	if ex_v != "Second" {
		t.Errorf("Sum was incorrect, got: %v, want: %v.", ex_v, "Second")
	}
	stringStack.Push("Third")
	ex_v = stringStack.Pop()
	if ex_v != "Third" {
		t.Errorf("Item was incorrect, got: %v, want: %v.", ex_v, "Third")
	}
	ex_v = stringStack.Pop()
	if ex_v != "First" {
		t.Errorf("Item was incorrect, got: %v, want: %v.", ex_v, "First")
	}

	should_be_nil := stringStack.Pop()
	if should_be_nil != nil {
		t.Errorf("Item was incorrect, got: %v, want: %v.", should_be_nil, "nil")
	}
}
