/*
Sum Root to Leaf Numbers
Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path 1->2->3 which represents the number 123.
Find the total sum of all root-to-leaf numbers.
For example,
    1
   / \
  2   3
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Return the sum = 12 + 13 = 25.
*/

package main

// eazy dfs
func sumNumbers(root *TreeNode) int {
	res := 0
	dfs(root, &res, 0)
	return res
}

func dfs(root *TreeNode, sum *int, current int) {
	if root == nil {
		return
	}
	current = current * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		*sum += current
	} else {
		dfs(root.Left, sum, current)
		dfs(root.Right, sum, current)
	}
}
