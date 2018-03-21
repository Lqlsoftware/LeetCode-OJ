/*
Construct Binary Tree from Preorder and Inorder Traversal
Given preorder and inorder traversal of a tree, construct the binary tree.
Note:
You may assume that duplicates do not exist in the tree.

For example, given
preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

package main

// recursively
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	root,i := preorder[0],0
	for i < len(inorder) && inorder[i] != root {
		i++
	}
	return &TreeNode{root,buildTree(preorder[1:i + 1],inorder[:i]),buildTree(preorder[i + 1:],inorder[i + 1:])}
}