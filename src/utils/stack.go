package main

import (
	"errors"
)

type Stack struct {
	Head *ListNode
	Len  int
}

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (stack *Stack) Pop() interface{} {
	if stack.Len != 0 {
		x := stack.Head.Val
		stack.Head = stack.Head.Next
		stack.Len--
		return x
	}
	return nil
}

func (stack *Stack) Push(val interface{}) error {
	element := new(ListNode)
	if element != nil {
		element.Val = val.(int)
		element.Next = stack.Head
		stack.Head = element
		stack.Len++
		return nil
	}
	return errors.New("stack over flow")
}

func (stack *Stack) Peak() interface{} {
	if stack.Len != 0 {
		return stack.Head.Val
	}
	return nil
}
