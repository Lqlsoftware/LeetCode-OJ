/*
Binary Tree Inorder Traversal
Given a binary tree, return the inorder traversal of its nodes' values.

For example:
Given binary tree [1,null,2,3],
   1
    \
     2
    /
   3
return [1,3,2].
Note: Recursive solution is trivial, could you do it iteratively?
*/

package main

func inorderTraversal(root *TreeNode) []int {
	var result []int
	stack,current := list.New(),root
	for current != nil || stack.Len() > 0 {
		for current != nil {
			stack.PushFront(current)
			current = current.Left
		}
		current = stack.Remove(stack.Front()).(*TreeNode)
		result = append(result, current.Val)
		current = current.Right
	}
	return result
}

// Morris Traversal 对原二叉树进行修改 不推荐使用
func inorderTraversal1(root *TreeNode) []int {
	var result []int
	current,pre := root,root
	for current != nil {
		if current.Left == nil {
			result = append(result, current.Val)
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
	return result
}

// 递归解法
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)
	return append(append(left,root.Val),right...)
}