/*
Reverse Linked List II
Reverse a linked list from position m to n. Do it in-place and in one-pass.

For example:
Given 1->2->3->4->5->NULL, m = 2 and n = 4,
return 1->4->3->2->5->NULL.

Note:
Given m, n satisfy the following condition:
1 ≤ m ≤ n ≤ length of list.
*/

package main

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	fakeHead := ListNode{0,head}
	pre := &fakeHead
	for i := 1;i < m;i++ {
		pre = pre.Next
	}
	start := pre.Next
	then := start.Next
	for i := 0;i < n - m;i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}
	return fakeHead.Next
}