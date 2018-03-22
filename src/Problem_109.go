/*
Convert Sorted List to Binary Search Tree
Given a singly linked list where elements are sorted in ascending order, convert it to a height balanced BST.
For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example:

Given the sorted linked list: [-10,-3,0,5,9],
One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

      0
     / \
   -3   9
   /   /
 -10  5
*/

package main

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var nums []int
	for curr := head;curr != nil;curr = curr.Next {
		nums = append(nums, curr.Val)
	}
	// Problem 108‘s solution
	sortedArrayToBST(nums)
}

func sortedListToBST1(head *ListNode) *TreeNode {
	length := 0
	for curr := head;curr != nil;curr = curr.Next {
		length++
	}
	return solve(head,length)
}

func solve(start *ListNode,length int) *TreeNode {
	if start == nil || length == 0 {
		return nil
	}
	left := solve(start, length / 2)
	for count := 0;count < length / 2;count++ {
		start = start.Next
	}
	return &TreeNode{
		Val:	start.Val,
		Left:	left,
		Right:	solve(start.Next, (length - 1) / 2),
	}
}