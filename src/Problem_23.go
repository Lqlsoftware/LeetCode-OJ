/*
Merge k Sorted Lists
Merge k sorted linked lists and return it as one sorted list.
Analyze and describe its complexity.
*/

package main

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	target := lists[0]
	index := make(map[int]*ListNode)
	for i := 1;i < len(lists);i++ {
		target,index = mergeTwoLists(target, lists[i], index)
	}
	return target
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode, index map[int]*ListNode) (*ListNode,map[int]*ListNode) {
	if l1 == nil {
		return l2, index
	} else if l2 == nil {
		return l1, index
	} else if l1.Val > l2.Val {
		return mergeTwoLists(l2, l1, index)
	}
	ptr1, ptr2 := l1, l2
	for ptr2 != nil {
		if i,err := index[ptr2.Val]; err == true {
			ptr1 = i
		} else {
			for ptr1.Next != nil && ptr1.Next.Val <= ptr2.Val {
				ptr1 = ptr1.Next
			}
		}
		if ptr1 == nil {
			ptr1.Next = ptr2
			index[ptr1.Val] = ptr1
			return l1, index
		} else {
			index[ptr1.Val] = ptr2
			ptr2.Next, ptr1.Next, ptr2 = ptr1.Next, ptr2, ptr2.Next
		}
	}
	return l1, index
}
