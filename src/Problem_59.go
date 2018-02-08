/*
Spiral Matrix II
Given an integer n, generate a square matrix filled with elements from 1 to n2 in spiral order.

For example,
Given n = 3,
You should return the following matrix:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
*/

package main

// 模拟螺旋的方式
func generateMatrix(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	matrix := make([][]int,n)
	for i := range matrix {
		matrix[i] = make([]int,n)
	}
	towards := 3 // 0 - down 1 - up 2 - left 3 - right
	width, height, i, j := n - 1, n - 1, 0, 0
	start := 1
	for idx := 1;idx <= n * n;idx++ {
		matrix[i][j] = idx
		switch towards {
		case 0:
			if i == height {
				towards = 2
				j--
			} else {
				i++
			}
		case 1:
			if i == start {
				start++
				width--
				height--
				towards = 3
				j++
			} else {
				i--
			}
		case 2:
			if j == start - 1 {
				towards = 1
				i--
			} else {
				j--
			}
		case 3:
			if j == width {
				towards = 0
				i++
			} else {
				j++
			}
		}
	}
	return matrix
}

//  层层推进
func generateMatrix1(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	matrix := make([][]int,n)
	for i := range matrix {
		matrix[i] = make([]int,n)
	}
	start,width,height := 0,n - 1,n - 1
	length := n * n
	current := 1
	for start <= height && start <= width {
		if start == height && start == width {
			matrix[start][start] = current
		}
		for i := start;i < width;i++ {
			matrix[start][i] = current
			current++
		}
		for i := start;i < height;i++ {
			matrix[i][width] = current
			current++
		}
		for i := width;i > start;i-- {
			matrix[height][i] = current
			current++
			if current == length + 1 {
				return matrix
			}
		}
		for i := height;i > start;i-- {
			matrix[i][start] = current
			current++
			if current == length + 1 {
				return matrix
			}
		}
		start++
		width--
		height--
	}
	return matrix
}