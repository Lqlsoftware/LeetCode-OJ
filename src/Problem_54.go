/*
Spiral Matrix
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

For example,
Given the following matrix:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
You should return [1,2,3,6,9,8,7,4,5].
*/

package main

// 模拟螺旋的方式
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	result := make([]int,len(matrix) * len(matrix[0]))
	towards := 3 // 0 - down 1 - up 2 - left 3 - right
	width, height, i, j := len(matrix[0]) - 1, len(matrix) - 1, 0, 0
	start := 1
 	for idx := range result {
		result[idx] = matrix[i][j]
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
	return result
}

//  层层推进
func spiralOrder1(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	start,width,height := 0,len(matrix[0]) - 1,len(matrix) - 1
	length := len(matrix) * len(matrix[0])
	result,current := make([]int, length),0
	for start <= height && start <= width {
		if start == height && start == width {
			result[current] = matrix[start][start]
		}
		for i := start;i < width;i++ {
			result[current] = matrix[start][i]
			current++
		}
		for i := start;i < height;i++ {
			result[current] = matrix[i][width]
			current++
		}
		for i := width;i > start;i-- {
			result[current] = matrix[height][i]
			current++
			if current == length {
				return result
			}
		}
		for i := height;i > start;i-- {
			result[current] = matrix[i][start]
			current++
			if current == length {
				return result
			}
		}
		start++
		width--
		height--
	}
	return result
}