/*
Partition List
Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
You should preserve the original relative order of the nodes in each of the two partitions.

For example,
Given 1->4->3->2->5->2 and x = 3,
return 1->2->2->4->3->5.
*/

package main

// 将小于x的node插入到新链尾部，再将剩余接入尾部
func partition(head *ListNode, x int) *ListNode {
	fakeHead,result := &ListNode{0,head},&ListNode{0,nil}
	r,q := result,fakeHead
	for p := head;p != nil;p = q.Next {
		if p.Val < x {
			q.Next,r.Next,p.Next,r = p.Next,p,nil,p
		} else {
			q = q.Next
		}
	}
	r.Next = fakeHead.Next
	return result.Next
}