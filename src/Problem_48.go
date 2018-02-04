/*
Rotate Image
You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).

Note:
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
DO NOT allocate another 2D matrix and do the rotation.

Example 1:
Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],
rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]

Example 2:
Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],
rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/

package main

// 每次旋转当前正方形矩阵的第一行
func rotate(matrix [][]int)  {
	length, temp := len(matrix) - 1,0
	for start := 0;length >= 0;start++ {
		for i := start;i < length;i++ {
			temp = length - i + start
			matrix[start][i], matrix[i][length] = matrix[i][length], matrix[start][i]
			matrix[start][i], matrix[length][temp] = matrix[length][temp], matrix[start][i]
			matrix[start][i], matrix[temp][start] = matrix[temp][start], matrix[start][i]
		}
		length--
	}
}