package main

import (
	"fmt"

	"github.com/kfelter/stacks"
)

type Queue struct {
	Inverted bool //always init as true
	S        *stacks.Stack
}

func main() {
	stringStack := &stacks.Stack{}
	fmt.Println("START STACK TEST")
	stringStack.Push("First")
	stringStack.Push("Second")

	fmt.Println(stringStack.ToString())
	fmt.Println("END STACK TEST")

	myqueue := NewQueuePointer()
	//-----------------------------------
	//First item in queue
	myqueue.Enqueue(map[string]string{
		"first": "Q",
	})
	fmt.Println("Enqueue: ", map[string]string{
		"first": "Q",
	})
	//---------------------------------

	//------------------------------------
	//second item in queue
	myqueue.Enqueue(map[string]string{
		"second": "Q",
	})
	fmt.Println("Enqueue: ", map[string]string{
		"second": "Q",
	})
	//-------------------------------------

	//----------------------------------
	//Dequeue first item
	fmt.Println("Dequeue an element: ")
	fmt.Println(myqueue.Dequeue())
	//---------------------------------

	//enqueue third item
	fmt.Println("Enqueue: ", map[string]string{
		"third": "Q",
	})
	myqueue.Enqueue(map[string]string{
		"third": "Q",
	})
	//dequeue remaining items

	fmt.Println("Dequeue an element: ")
	fmt.Println(myqueue.Dequeue())
	fmt.Println("Dequeue an element: ")
	fmt.Println(myqueue.Dequeue())
}

func (q *Queue) Enqueue(elem interface{}) {
	if q.Inverted {
		q.S.Push(elem)
	} else {
		temp := &stacks.Stack{}
		telem := q.S.Pop()
		for telem != nil {
			temp.Push(telem)
			telem = q.S.Pop()
		}
		temp.Push(elem)
		q.S = temp
		q.Inverted = !q.Inverted
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.Inverted {
		temp := &stacks.Stack{}
		elem := q.S.Pop()
		for elem != nil {
			temp.Push(elem)
			elem = q.S.Pop()
		}
		elem = temp.Pop()
		q.S = temp
		q.Inverted = !q.Inverted
		return elem
	} else {
		return q.S.Pop()
	}
}

func NewQueuePointer() *Queue {
	return &Queue{
		Inverted: true,
		S:        &stacks.Stack{},
	}
}
