/*
Path Sum II
Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

For example:
Given the below binary tree and sum = 22,
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
return
[
   [5,4,11,2],
   [5,8,4,5]
]
*/

package main

func getPathSum(result *[][]int, current []int, root *TreeNode, sum int) {
	if root == nil {
		return
	} else if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			temp := make([]int,len(current) + 1)
			copy(temp,current)
			temp[len(current)] = root.Val
			*result = append(*result, temp)
		}
		return
	}
	current = append(current, root.Val)
	next := sum - root.Val
	getPathSum(result, current, root.Left, next)
	getPathSum(result, current, root.Right, next)
}

func pathSum(root *TreeNode, sum int) [][]int {
	var result [][]int
	getPathSum(&result, []int{}, root, sum)
	return result
}