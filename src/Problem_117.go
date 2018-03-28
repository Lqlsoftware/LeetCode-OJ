/*
Populating Next Right Pointers in Each Node II
Follow up for problem "Populating Next Right Pointers in Each Node".
What if the given tree could be any binary tree? Would your previous solution still work?

Note:
You may only use constant extra space.
For example,
Given the following binary tree,
         1
       /  \
      2    3
     / \    \
    4   5    7
After calling your function, the tree should look like:
         1 -> NULL
       /  \
      2 -> 3 -> NULL
     / \    \
    4-> 5 -> 7 -> NULL
*/

package main

// 在访问父母的时候将孩子连接起来 下一层从连接的孩子开始访问
func connect(root *TreeLinkNode) {
	curr := root
	fakeHead := &TreeLinkNode{0,nil,nil,nil}
	for curr != nil {
		last := fakeHead
		for curr != nil {
			if curr.Left != nil {
				last.Next = curr.Left
				last = last.Next
			}
			if curr.Right != nil {
				last.Next = curr.Right
				last = last.Next
			}
			curr = curr.Next
		}
		last.Next,curr = nil,fakeHead.Next
	}
}