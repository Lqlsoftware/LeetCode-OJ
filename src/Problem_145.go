/*
Binary Tree Postorder Traversal
Given a binary tree, return the postorder traversal of its nodes' values.

Example:
Input: [1,null,2,3]
   1
    \
     2
    /
   3
Output: [3,2,1]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

package main

// 压伪造节点入栈
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := list.New()
	stack.PushFront(root)
	var res []int
	for stack.Len() != 0 {
		curr := stack.Remove(stack.Front()).(*TreeNode)
		if curr.Left != nil {
			stack.PushFront(&TreeNode{curr.Val, nil, curr.Right})
			stack.PushFront(curr.Left)
		} else if curr.Right != nil {
			stack.PushFront(&TreeNode{curr.Val, nil, nil})
			stack.PushFront(curr.Right)
		} else {
			res = append(res, curr.Val)
		}
	}
	return res
}

// 修改栈顶元素指针 不推荐 会修改原二叉树
func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := list.New()
	stack.PushFront(root)
	var res []int
	for stack.Len() != 0 {
		curr := stack.Front().Value.(*TreeNode)
		if curr.Left != nil {
			stack.PushFront(curr.Left)
			curr.Left = nil
		} else if curr.Right != nil {
			stack.PushFront(curr.Right)
			curr.Right = nil
		} else {
			res = append(res, curr.Val)
			stack.Remove(stack.Front())
		}
	}
	return res
}