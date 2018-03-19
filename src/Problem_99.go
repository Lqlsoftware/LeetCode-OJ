/*
Recover Binary Search Tree
Two elements of a binary search tree (BST) are swapped by mistake.
Recover the tree without changing its structure.
Note:
A solution using O(n) space is pretty straight forward. Could you devise a constant space solution?
*/

package main

// 中序遍历 第一个比前面小的错误 m1,m2 = last,curr  第二个错误 m2 = curr
// Morris Traversal using O(1) space
func recoverTree(root *TreeNode)  {
	var pre,last,m1,m2 *TreeNode
	for root != nil {
		if root.Left == nil {
			if last != nil && root.Val < last.Val {
				if m1 == nil {
					m1 = last
				}
				m2 = root
			}
			last,root = root,root.Right
		} else {
			// 查找前驱节点
			for pre = root.Left;pre.Right != nil;pre = pre.Right {
				if pre.Right == root {
					break
				}
			}
			// 前驱节点对右孩子为空则添加当前节点到右孩子 否则恢复前驱节点的右孩子
			if pre.Right == nil {
				pre.Right,root = root,root.Left
			} else {
				pre.Right = nil
				if last != nil && root.Val < last.Val {
					if m1 == nil {
						m1 = last
					}
					m2 = root
				}
				last,root = root,root.Right
			}
		}
	}
	if m1 != nil && m2 != nil {
		m1.Val, m2.Val = m2.Val, m1.Val
	}
}

// Inorder Travel using O(n) space
func recoverTree1(root *TreeNode)  {
	var last,m1,m2 *TreeNode
	curr := root
	stack := list.New()
	for curr != nil || stack.Len() > 0 {
		for curr != nil {
			stack.PushFront(curr)
			curr = curr.Left
		}
		curr = stack.Remove(stack.Front()).(*TreeNode)
		if last != nil && curr.Val <= last.Val {
			if m1 == nil {
				m1 = last
			}
			m2 = curr
		}
		last,curr = curr,curr.Right
	}
	if m1 != nil && m2 != nil {
		m1.Val, m2.Val = m2.Val, m1.Val
	}
}