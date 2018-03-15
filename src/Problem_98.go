/*
Validate Binary Search Tree
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:
The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
Example 1:
    2
   / \
  1   3
Binary tree [2,1,3], return true.
Example 2:
    1
   / \
  2   3
Binary tree [1,2,3], return false.
*/

package main

// 中序遍历
func isValidBST(root *TreeNode) bool {
	last := math.MinInt64
	stack,current := list.New(),root
	for current != nil || stack.Len() > 0 {
		for current != nil {
			stack.PushFront(current)
			current = current.Left
		}
		current = stack.Remove(stack.Front()).(*TreeNode)
		if current.Val <= last {
			return false
		}
		last = current.Val
		current = current.Right
	}
	return true
}

// Morris Traversal 中序遍历
func isValidBST1(root *TreeNode) bool {
	last := math.MinInt64
	current,pre := root,root
	for current != nil {
		if current.Left == nil {
			if current.Val <= last {
				return false
			}
			last = current.Val
			current = current.Right
		} else {
			pre = current.Left
			for pre.Right != nil {
				pre = pre.Right
			}
			pre.Right = current
			current,current.Left = current.Left,nil
		}
	}
	return true
}

// 递归处理
func isValidBST2(root *TreeNode) bool {
	return solve(root,math.MinInt64,math.MaxInt64)
}

func solve(root *TreeNode, min, max int64) bool {
	if root == nil {
		return true
	}
	val := int64(root.Val)
	if val <= min || val >= max {
		return false
	}
	return solve(root.Left,min,val) && solve(root.Right,val,max)
}