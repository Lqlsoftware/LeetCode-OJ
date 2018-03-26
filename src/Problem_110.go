/*
Balanced Binary Tree
Given a binary tree, determine if it is height-balanced.
For this problem, a height-balanced binary tree is defined as:
a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

Example 1:
Given the following tree [3,9,20,null,null,15,7]:

    3
   / \
  9  20
    /  \
   15   7
Return true.

Example 2:
Given the following tree [1,2,2,3,3,null,null,4,4]:

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
Return false.
*/

package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	result := true
	getDepth(root,&result)
	return result
}

func getDepth(root *TreeNode, result *bool) int {
	if root == nil {
		return 0
	} else if *result == false {
		return -1
	}
	left := getDepth(root.Left, result)
	right := getDepth(root.Right, result)
	if diff := left - right;diff > 1 || diff < -1 {
		*result = false
	}
	return max(left,right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}