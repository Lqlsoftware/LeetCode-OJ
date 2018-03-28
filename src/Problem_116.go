/*
Populating Next Right Pointers in Each Node
Given a binary tree
    struct TreeLinkNode {
      TreeLinkNode *left;
      TreeLinkNode *right;
      TreeLinkNode *next;
    }
Populate each next pointer to point to its next right node. If there is no next right node, the next pointer should be set to NULL.
Initially, all next pointers are set to NULL.

Note:
You may only use constant extra space.
You may assume that it is a perfect binary tree (ie, all leaves are at the same level, and every parent has two children).
For example,
Given the following perfect binary tree,
         1
       /  \
      2    3
     / \  / \
    4  5  6  7
After calling your function, the tree should look like:
         1 -> NULL
       /  \
      2 -> 3 -> NULL
     / \  / \
    4->5->6->7 -> NULL
*/

package main

type TreeLinkNode struct {
	Val 	int
	Left	*TreeLinkNode
	Right	*TreeLinkNode
	Next	*TreeLinkNode
}

// 后序遍历 翻译成c++后AC
func connect(root *TreeLinkNode) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	fakeHead := &TreeLinkNode{0,nil,nil,nil}
	for queue.Len() != 0 {
		size,last := queue.Len(),fakeHead
		for i := 0;i < size;i++ {
			curr := queue.Remove(queue.Front()).(*TreeLinkNode)
			last.Next,last = curr,curr
			if curr.Left != nil {
				queue.PushBack(curr.Left)
			}
			if curr.Right != nil {
				queue.PushBack(curr.Right)
			}
		}
		last.Next = nil
	}
}