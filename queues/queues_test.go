package queues

import (
	"testing"
)

func TestQueue(t *testing.T) {
	ex1 := map[string]int{"number": 1}
	ex2 := map[string]int{"number": 2}
	ex3 := map[string]int{"number": 3}

	myqueue := NewQueuePointer()

	myqueue.Enqueue(ex1)
	myqueue.Enqueue(ex2)

	//test 1
	val := myqueue.Dequeue().(map[string]int)
	if val["number"] != ex1["number"] {
		t.Errorf("Val was incorrect, got: %v, want: %v.", val, ex1)
	}

	//test 2
	val = myqueue.Dequeue().(map[string]int)
	if val["number"] != ex2["number"] {
		t.Errorf("Val was incorrect, got: %v, want: %v.", val, ex2)
	}

	myqueue.Enqueue(ex3)
	//test 3
	val = myqueue.Dequeue().(map[string]int)
	if val["number"] != ex3["number"] {
		t.Errorf("Val was incorrect, got: %v, want: %v.", val, ex3)
	}
	//dequeue nothing should be nil
	if myqueue.Dequeue() != nil {
		t.Errorf("Val was incorrect, got: %v, want: %v.", myqueue.Dequeue(), "nil")
	}

}
