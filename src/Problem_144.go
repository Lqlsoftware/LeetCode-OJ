/*
Binary Tree Preorder Traversal
Given a binary tree, return the preorder traversal of its nodes' values.

Example:
Input: [1,null,2,3]
   1
    \
     2
    /
   3
Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

package main

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := list.New()
	stack.PushFront(root)
	var res []int
	for stack.Len() != 0 {
		curr := stack.Remove(stack.Front()).(*TreeNode)
		res = append(res, curr.Val)
		if curr.Right != nil {
			stack.PushFront(curr.Right)
		}
		if curr.Left != nil {
			stack.PushFront(curr.Left)
		}
	}
	return res
}