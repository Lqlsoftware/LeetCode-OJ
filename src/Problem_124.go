/*
Binary Tree Maximum Path Sum
Given a binary tree, find the maximum path sum.
For this problem, a path is defined as any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The path must contain at least one node and does not need to go through the root.

For example:
Given the below binary tree,

       1
      / \
     2   3
Return 6.
*/

package main

// an elegant solution
func maxPathSum(root *TreeNode) int {
	result := math.MinInt32
	maxPath(&result, root)
	return result
}

func maxPath(result *int,root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := max(0,maxPath(result, root.Left))
	right := max(0,maxPath(result, root.Right))
	// 如果大于最大值就更新,此时最大值已经包括左右两孩子内部的最大值(上面递归过了)
	*result = max(*result, left + right + root.Val)
	return max(left, right) + root.Val
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}