package main

import (
	"container/list"
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewBinaryTree(val []interface{}) *TreeNode {
	if len(val) == 0 || val[0] == nil {
		return nil
	}
	queue := list.New()
	root := &TreeNode{val[0].(int), nil, nil}
	queue.PushBack(root)
	current := root
	for i := 1; i < len(val); i++ {
		current = queue.Remove(queue.Front()).(*TreeNode)
		if val[i] != nil {
			current.Left = &TreeNode{val[i].(int), nil, nil}
			queue.PushBack(current.Left)
		}
		i++
		if i < len(val) && val[i] != nil {
			current.Right = &TreeNode{val[i].(int), nil, nil}
			queue.PushBack(current.Right)
		}
	}
	return root
}

func (root *TreeNode) Print() *TreeNode {
	iRoot := createInfoTree(root)
	iRoot = calcSpace(iRoot)
	printBinaryTree(iRoot)
	return root
}

type infoTree struct {
	// 用于输出图形
	Val        int
	Level      int
	Left       *infoTree
	Right      *infoTree
	PreSpace   int
	LeftSpace  int
	RightSpace int
	oragin     *TreeNode
}

func createInfoTree(root *TreeNode) *infoTree {
	if root == nil {
		return nil
	}
	queue := list.New()
	iRoot := &infoTree{root.Val, 1, nil, nil, 0, 0, 0, root}
	queue.PushBack(iRoot)
	currentLevel := 1
	for queue.Len() != 0 {
		cur := queue.Remove(queue.Front()).(*infoTree)
		if cur.Level > currentLevel {
			currentLevel = cur.Level
		}
		if cur.oragin.Left != nil {
			left := &infoTree{cur.oragin.Left.Val, currentLevel + 1, nil, nil, 0, 0, 0, cur.oragin.Left}
			cur.Left = left
			queue.PushBack(left)
		} else {
			cur.Left = &infoTree{math.MaxInt32, currentLevel + 1, nil, nil, 0, 0, 0, nil}
		}
		if cur.oragin.Right != nil {
			right := &infoTree{cur.oragin.Right.Val, currentLevel + 1, nil, nil, 0, 0, 0, cur.oragin.Right}
			cur.Right = right
			queue.PushBack(right)
		} else {
			cur.Right = &infoTree{math.MaxInt32, currentLevel + 1, nil, nil, 0, 0, 0, nil}
		}
	}
	return iRoot
}

func calcCharacters(value int) int {
	if value == 0 || value == math.MaxInt32 {
		return 1
	}
	characters := 0
	for value != 0 {
		characters++
		value /= 10
	}
	return characters
}

func calcSpace(root *infoTree) *infoTree {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		root.LeftSpace, root.RightSpace = 0, 0
	} else if root.Left == nil && root.Right != nil {
		root.Right.PreSpace = calcCharacters(root.Val)
		root.LeftSpace = 1
		right := calcSpace(root.Right)
		root.RightSpace = right.LeftSpace + calcCharacters(root.Right.Val) + right.RightSpace
	} else if root.Left != nil && root.Right == nil {
		root.Left.PreSpace = root.PreSpace
		root.RightSpace = 0
		left := calcSpace(root.Left)
		root.LeftSpace = left.LeftSpace + calcCharacters(root.Left.Val) + left.RightSpace
	} else {
		//set pre spaces, from father to son
		root.Left.PreSpace = root.PreSpace
		root.Right.PreSpace = calcCharacters(root.Val)
		//set left spaces
		left := calcSpace(root.Left)
		root.LeftSpace = left.LeftSpace + calcCharacters(root.Left.Val) + left.RightSpace
		//set right spaces
		right := calcSpace(root.Right)
		root.RightSpace = right.LeftSpace + calcCharacters(root.Right.Val) + right.RightSpace
	}
	return root
}

func printTreeNode(root *infoTree) {
	if root == nil {
		return
	}
	leftSpace := root.LeftSpace + root.PreSpace
	for i := 0; i < leftSpace; i++ {
		fmt.Print(" ")
	}
	if root.Val == math.MaxInt32 {
		fmt.Print(" ")
	} else {
		fmt.Print(root.Val)
	}
	for i := 0; i < root.RightSpace; i++ {
		fmt.Print(" ")
	}
}

func printBinaryTree(root *infoTree) {
	if root == nil {
		return
	}
	queue := list.New()
	queue.PushBack(root)
	currentLevel := 1
	for queue.Len() != 0 {
		cur := queue.Remove(queue.Front()).(*infoTree)
		if cur.Level > currentLevel {
			fmt.Println()
			currentLevel = cur.Level
		}
		printTreeNode(cur)
		if cur.Left != nil {
			queue.PushBack(cur.Left)
		}
		if cur.Right != nil {
			queue.PushBack(cur.Right)
		}
	}
	fmt.Println()
}
