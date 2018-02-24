/*
Remove Duplicates from Sorted List
Given a sorted linked list, delete all duplicates such that each element appear only once.

For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.
*/

package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	for p := head;p.Next != nil; {
		if p.Next.Val == p.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}