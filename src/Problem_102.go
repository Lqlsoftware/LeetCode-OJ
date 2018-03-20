/*
Binary Tree Level Order Traversal
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/

package main

// recursively
func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	solve(&result,root,0)
	return result
}

func solve(result *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(*result) {
		*result = append(*result, []int{root.Val})
	} else {
		(*result)[level] = append((*result)[level], root.Val)
	}
	level++
	solve(result,root.Left,level)
	solve(result,root.Right,level)
}

// iteratively
func levelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		var temp []int
		size := queue.Len()
		for i := 0;i < size;i++ {
			root = queue.Remove(queue.Front()).(*TreeNode)
			temp = append(temp, root.Val)
			if root.Left != nil {
				queue.PushBack(root.Left)
			}
			if root.Right != nil {
				queue.PushBack(root.Right)
			}
		}
		result = append(result, temp)
	}
	return result
}