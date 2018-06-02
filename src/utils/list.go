package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (listNode *ListNode) Print() *ListNode {
	if listNode == nil {
		return nil
	}
	for ptr := listNode; ptr != nil; ptr = ptr.Next {
		fmt.Print(ptr.Val)
		if ptr.Next != nil {
			fmt.Print("->")
		}
	}
	fmt.Print("\n")
	return listNode
}

func NewList(val []int) *ListNode {
	if len(val) == 0 {
		return nil
	}
	head := new(ListNode)
	ptr, last := head, head
	for _, v := range val {
		ptr = &ListNode{v, nil}
		last.Next = ptr
		last = ptr
	}
	return head.Next
}
