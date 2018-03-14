/*
Unique Binary Search Trees II
Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1...n.

For example,
Given n = 3, your program should return all 5 unique BST's shown below.

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

package main

// 递归	取根节点为k, 则1 ~ k - 1都在左子树, k + 1 ~ n都在右子树
// g(i,n) = g(i,k - 1) + g (k + 1,n)  (1 <= k <= n)
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateSubtrees(1, n)
}

func generateSubtrees(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var result []*TreeNode
	for i := start;i <= end;i++ {
		left := generateSubtrees(start,i - 1)
		right := generateSubtrees(i + 1,end)
		for _,l := range left {
			for _,r := range right {
				result = append(result, &TreeNode{i,l,r})
			}
		}
	}
	return result
}