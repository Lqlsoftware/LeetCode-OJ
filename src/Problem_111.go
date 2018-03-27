/*
Minimum Depth of Binary Tree
Given a binary tree, find its minimum depth.
The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
*/

package main

// BFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	level := 1
	for queue.Len() != 0 {
		size := queue.Len()
		for i := 0;i < size;i++ {
			curr := queue.Remove(queue.Front()).(*TreeNode)
			if curr.Left == nil && curr.Right == nil {
				return level
			}
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
		level++
	}
	return level
}