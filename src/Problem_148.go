/*
Sort List
Sort a linked list in O(n log n) time using constant space complexity.

Example 1:
Input: 4->2->1->3
Output: 1->2->3->4

Example 2:
Input: -1->5->3->4->0
Output: -1->0->3->4->5
*/

package main

// 归并排序 -> 链表光速排序
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, curr := head, head.Next
	for curr != nil && curr.Next != nil {
		pre = pre.Next
		curr = curr.Next.Next
	}
	right := sortList(pre.Next)
	pre.Next = nil
	left := sortList(head)
	return mergeList(left, right)
}

func mergeList(l1, l2 *ListNode) *ListNode {
	fakeHead := &ListNode{0,nil}
	curr := fakeHead
	for l1 != nil || l2 != nil {
		if l1 == nil {
			curr.Next = l2
			l2 = l2.Next
		} else if l2 == nil {
			curr.Next = l1
			l1 = l1.Next
		} else if l1.Val >= l2.Val {
			curr.Next = l2
			l2 = l2.Next
		} else {
			curr.Next = l1
			l1 = l1.Next
		}
		curr = curr.Next
	}
	return fakeHead.Next
}

// 快排 -> 链表低速排序
func sortList1(head *ListNode) *ListNode {
	return qsortList(head, nil)
}
func qsortList(start, end *ListNode) *ListNode {
	if start == end {
		return start
	}
	fakeHead := &ListNode{0, start}
	less, mid, curr := fakeHead, start, start
	for curr.Next != end {
		if curr.Next.Val < mid.Val {
			less.Next, curr.Next.Next, curr.Next, less = curr.Next, less.Next, curr.Next.Next, curr.Next
		} else {
			curr = curr.Next
		}
	}
	fakeHead.Next = qsortList(fakeHead.Next, mid)
	mid.Next = qsortList(mid.Next, end)
	return fakeHead.Next
}

// insert sort -> 爬行排序
func sortList(head *ListNode) *ListNode {
	fakeHead := &ListNode{0,head}
	var curr, pre, next *ListNode
	for curr = fakeHead; curr.Next != nil; curr = next {
		next = curr.Next
		if curr.Val < next.Val {
			continue
		}
		for pre = fakeHead; pre.Next != next;pre = pre.Next {
			if curr.Next.Val < pre.Next.Val {
				next = curr
				pre.Next, curr.Next.Next, curr.Next = curr.Next, pre.Next, curr.Next.Next
				break
			}
		}
	}
	return fakeHead.Next
}