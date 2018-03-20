/*
Binary Tree Zigzag Level Order Traversal
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its zigzag level order traversal as:
[
  [3],
  [20,9],
  [15,7]
]
*/

package main

// recursively 8ms
func zigzagLevelOrder(root *TreeNode) [][]int {
	var result [][]int
	solve(&result,root,0)
	return result
}

func solve(result *[][]int, root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level == len(*result) {
		*result = append(*result, []int{root.Val})
	} else if level & 1 == 0 {
		(*result)[level] = append((*result)[level], root.Val)
	} else {
		(*result)[level] = append([]int{root.Val},(*result)[level]...)
	}
	level++
	solve(result,root.Left,level)
	solve(result,root.Right,level)
}

// 定义双头顺序队列来加快进队出队的速度 0ms
const MaxSize = 2 << 8
type Queue struct {
	arr 	[]interface{}
	top		int
	end		int
}

func (queue *Queue)Clear() {
	if queue.arr == nil {
		queue.arr = make([]interface{}, MaxSize)
	}
	queue.top = MaxSize >> 1
	queue.end = queue.top
}

func (queue *Queue)Len() int {
	return (queue.top + MaxSize - queue.end) % MaxSize
}

func (queue *Queue)NoEmpty() bool {
	return queue.top != queue.end
}

func (queue *Queue)PushBack(v interface{}) {
	queue.end = (queue.end + MaxSize - 1) % MaxSize
	queue.arr[queue.end] = v
}

func (queue *Queue)PushFront(v interface{}) {
	queue.arr[queue.top] = v
	queue.top = (queue.top + 1) % MaxSize
}

func (queue *Queue)PopFront() interface{} {
	queue.top = (queue.top + MaxSize - 1) % MaxSize
	return queue.arr[queue.top]
}

func (queue *Queue)PopBack() interface{} {
	v := queue.arr[queue.end]
	queue.end = (queue.end + 1) % MaxSize
	return v
}

var queue Queue

func zigzagLevelOrder1(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	queue.Clear()
	queue.PushBack(root)
	flag := 0
	for queue.NoEmpty() {
		var temp []int
		size := queue.Len()
		if flag == 0 {
			for i := 0; i < size; i++ {
				root = queue.PopFront().(*TreeNode)
				temp = append(temp, root.Val)
				if root.Left != nil {
					queue.PushBack(root.Left)
				}
				if root.Right != nil {
					queue.PushBack(root.Right)
				}
			}
		} else {
			for i := 0; i < size; i++ {
				root = queue.PopBack().(*TreeNode)
				temp = append(temp, root.Val)
				if root.Right != nil {
					queue.PushFront(root.Right)
				}
				if root.Left != nil {
					queue.PushFront(root.Left)
				}
			}
		}
		flag ^= 1
		result = append(result, temp)
	}
	return result
}

// 双栈切换方法 4ms
type S struct {
	arr 	[]interface{}
	top		int
}

func (stack *S)Clear() {
	if stack.arr == nil {
		stack.arr = make([]interface{}, MaxSize)
	}
	stack.top = -1
}

func (stack *S)Len() int {
	return stack.top + 1
}

func (stack *S)NoEmpty() bool {
	return stack.top != -1
}

func (stack *S)Push(v interface{}) {
	stack.top++
	stack.arr[stack.top] = v
}

func (stack *S)Pop() interface{} {
	v := stack.arr[stack.top]
	stack.top--
	return v
}

var stack1,stack2 S

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	stack1.Clear()
	stack2.Clear()
	stack1.Push(root)
	flag := 0
	for stack1.NoEmpty() {
		var temp []int
		size := stack1.Len()
		for i := 0; i < size; i++ {
			root = stack1.Pop().(*TreeNode)
			temp = append(temp, root.Val)
			if flag == 0 {
				if root.Left != nil {
					stack2.Push(root.Left)
				}
				if root.Right != nil {
					stack2.Push(root.Right)
				}
			} else {
				if root.Right != nil {
					stack2.Push(root.Right)
				}
				if root.Left != nil {
					stack2.Push(root.Left)
				}
			}
		}
		flag ^= 1
		stack1,stack2 = stack2, stack1
		result = append(result, temp)
	}
	return result
}
