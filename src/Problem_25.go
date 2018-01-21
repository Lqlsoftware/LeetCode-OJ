/*
Reverse Nodes in k-Group
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.
k is a positive integer and is less than or equal to the length of the linked list.
If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.
You may not alter the values in the nodes, only nodes itself may be changed.
Only constant memory is allowed.

For example,
Given this linked list: 1->2->3->4->5
For k = 2, you should return: 2->1->4->3->5
For k = 3, you should return: 3->2->1->4->5
*/

package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	h := new(ListNode)
	h.Next = head
	p,q := h,h
	for q != nil {
		// next k nodes
		q = p
		for i:=k;i>0;i-- {
			q = q.Next
			if q == nil {
				return h.Next
			}
		}
		// exchange
		p.Next,p = reverse(p.Next,q)
	}
	return h.Next
}

func reverse(start, end *ListNode) (s *ListNode, e*ListNode) {
	e, s = start, end
	for start != end {
		end.Next,start.Next,start = start,end.Next,start.Next
	}
	return s, e
}