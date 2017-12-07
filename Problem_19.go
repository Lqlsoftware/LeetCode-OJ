/*
Remove Nth Node From End of List
Given a linked list, remove the nth node from the end of list and return its head.
For example,
   Given linked list: 1->2->3->4->5, and n = 2.
   After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:
Given n will always be valid.
Try to do this in one pass.
*/

package main

import "fmt"

func (listNode ListNode) print() {
	ptr := &listNode
	for ptr != nil {
		fmt.Print(ptr.Val)
		if ptr.Next != nil {
			fmt.Print("->")
		}
		ptr = ptr.Next
	}
}

func newList(val []int) *ListNode {
	head := new(ListNode)
	ptr, last := head, head
	for _, v := range val {
		last = ptr
		last.Val = v
		ptr = new(ListNode)
		last.Next = ptr
	}
	last.Next, ptr = nil, nil
	return head
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr, last := head, head
	// get last = ptr - n - 2
	for ptr.Next != nil && n != 0 {
		ptr = ptr.Next
		n--
	}
	// here last + n + 1 = tail
	for ptr.Next != nil {
		last = last.Next
		ptr = ptr.Next
	}
	if last.Next == nil || n > 0 {
		return last.Next
	}
	// delete last n node
	last.Next.Val = 0
	last.Next, last.Next.Next = last.Next.Next, nil
	return head
}
