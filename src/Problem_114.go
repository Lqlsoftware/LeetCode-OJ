/*
Flatten Binary Tree to Linked List
Given a binary tree, flatten it to a linked list in-place.

For example,
Given

         1
        / \
       2   5
      / \   \
     3   4   6
The flattened tree should look like:
   1
    \
     2
      \
       3
        \
         4
          \
           5
            \
             6
Hints:
If you notice carefully in the flattened tree, each node's right child points to the next node of a pre-order traversal.
*/

package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := list.New()
	stack.PushFront(root)
	curr,last := root,&TreeNode{0,nil,nil}
	for stack.Len() != 0 {
		curr = stack.Remove(stack.Front()).(*TreeNode)
		if curr.Right != nil {
			stack.PushFront(curr.Right)
		}
		if curr.Left != nil {
			stack.PushFront(curr.Left)
		}
		last.Right = curr
		last = curr
		curr.Left = nil
	}
	last.Right = nil
}