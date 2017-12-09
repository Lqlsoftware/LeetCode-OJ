/*
Merge Two Sorted Lists
Merge two sorted linked lists and return it as a new list.
The new list should be made by splicing together the nodes of the first two lists.

Example:
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val > l2.Val {
		return mergeTwoLists(l2, l1)
	}
	ptr1, ptr2 := l1, l2
	for ptr2 != nil {
		for ptr1.Next != nil && ptr1.Next.Val < ptr2.Val {
			ptr1 = ptr1.Next
		}
		if ptr1 == nil {
			ptr1.Next = ptr2
			return l1
		} else {
			ptr2.Next, ptr1.Next, ptr2 = ptr1.Next, ptr2, ptr2.Next
		}
	}
	return l1
}
