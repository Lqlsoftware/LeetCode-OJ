/*
Symmetric Tree
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
But the following [1,2,2,null,3,null,3] is not:
    1
   / \
  2   2
   \   \
   3    3
Note:
Bonus points if you could solve it both recursively and iteratively.
*/

package main

// recursively
func isSymmetric(root *TreeNode) bool {
    if root == nil {
		return true
	}
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}

// iteratively Morris Traversal
func isSymmetric1(root *TreeNode) bool {
	var pre1,pre2 *TreeNode
	curr1,curr2 := root,root
	var val1,val2 int
	for curr1 != nil || curr2 != nil {
		if curr1 == nil || curr2 == nil {
			return false
		}
		if curr1.Left == nil {
			val1 = curr1.Val
			curr1 = curr1.Right
		} else {
			for pre1 = curr1.Left;pre1.Right != nil;pre1 = pre1.Right {
				if pre1.Right == curr1 {
					break
				}
			}
			if pre1.Right == nil {
				pre1.Right,curr1 = curr1,curr1.Left
			} else {
				pre1.Right = nil
				val1 = curr1.Val
				curr1 = curr1.Right
			}
		}
		if curr2.Right == nil {
			val2 = curr2.Val
			curr2 = curr2.Left
		} else {
			for pre2 = curr2.Right;pre2.Left != nil;pre2 = pre2.Left {
				if pre2.Left == curr2 {
					break
				}
			}
			if pre2.Left == nil {
				pre2.Left,curr2 = curr2,curr2.Right
			} else {
				pre2.Left = nil
				val2 = curr2.Val
				curr2 = curr2.Left
			}
		}
		if val1 != val2 {
			return false
		}
	}
	return true
}