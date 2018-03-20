/*
Maximum Depth of Binary Tree
Given a binary tree, find its maximum depth.
The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

For example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

package main

// dfs
func maxDepth(root *TreeNode) int {
	if root == nil {
		 return 0
	}
	return max(maxDepth(root.Left),maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}