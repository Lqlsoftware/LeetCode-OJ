/*
Swap Nodes in Pairs
Given a linked list, swap every two adjacent nodes and return its head.

For example,
Given 1->2->3->4, you should return the list as 2->1->4->3.

Your algorithm should use only constant space.
You may not modify the values in the list, only nodes itself can be changed.
*/

package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	h := new(ListNode)
	h.Next = head
	p,q := h,head.Next
	for q != nil {
		// exchange
		p.Next.Next = q.Next
		q.Next = p.Next
		p.Next = q
		// next 2 nodes
		p = q.Next
		q = p.Next
		if q != nil {
			q = q.Next
		}
	}
	return h.Next
}
