/*
Add Two Numbers
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/
package main

// ListNode list
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c := 0
	p, q := l1, l2
	res := new(ListNode)
	r := res
	r.Next = &ListNode{
		Val:  0,
		Next: res,
	}
	for p != nil || q != nil {
		r = r.Next
		if p == nil {
			r.Val = (q.Val + c) % 10
			c = (q.Val + c) / 10
			q = q.Next
		} else if q == nil {
			r.Val = (p.Val + c) % 10
			c = (p.Val + c) / 10
			p = p.Next
		} else {
			r.Val = (p.Val + q.Val + c) % 10
			c = (p.Val + q.Val + c) / 10
			p = p.Next
			q = q.Next
		}
		r.Next = new(ListNode)
	}
	if c == 0 {
		r.Next = nil
	} else {
		r.Next.Val = c
	}
	return res.Next
}
