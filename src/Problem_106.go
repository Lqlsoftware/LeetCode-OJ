/*
Construct Binary Tree from Inorder and Postorder Traversal
Given inorder and postorder traversal of a tree, construct the binary tree.
Note:
You may assume that duplicates do not exist in the tree.

For example, given
inorder = [9,3,15,20,7]
postorder = [9,15,7,20,3]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
*/

package main

// recursively
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 || len(postorder) != len(inorder) {
		return nil
	}
	last := len(postorder) - 1
	root,i := postorder[last],last
	for i >= 0 && inorder[i] != root {
		i--
	}
	return &TreeNode{root,buildTree(inorder[:i],postorder[:i]),buildTree(inorder[i + 1:],postorder[i:last])}
}