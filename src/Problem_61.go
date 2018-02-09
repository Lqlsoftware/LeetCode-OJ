/*
Rotate List
Given a list, rotate the list to the right by k places, where k is non-negative.

Example:
Given 1->2->3->4->5->NULL and k = 2,
return 4->5->1->2->3->NULL.
*/

package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre,curr,length := head,head,0
	for ;curr.Next != nil;length++ {
		curr = curr.Next
	}
	k %= length + 1
	if k == 0 {
		return head
	}
	for i := 0;i < length - k;i++ {
		pre = pre.Next
	}
	curr.Next,head,pre.Next = head,pre.Next,nil
	return head
}