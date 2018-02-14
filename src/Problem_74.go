/*
Search a 2D Matrix
Write an efficient algorithm that searches for a value in an m x n matrix.
This matrix has the following properties:
    Integers in each row are sorted from left to right.
    The first integer of each row is greater than the last integer of the previous row.

For example,
Consider the following matrix:
[
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
Given target = 3, return true.
*/

package main

// BS col and row (avoids m * n multiplication overflow.)
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	start,end := 0,len(matrix) - 1
	for start <= end {
		mid := (start + end + 1) >> 1
		if v := matrix[mid][0];v == target {
			return true
		} else if v > target {
			end = mid - 1
		} else {
			if mid == start {
				break
			}
			start = mid
		}
	}
	y := start
	start,end = 0,len(matrix[0]) - 1
	for start <= end {
		mid := (start + end + 1) / 2
		if v := matrix[y][mid];v == target {
			return true
		} else if v > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}

// treat as an array (bad cache using)
func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n,start,end := len(matrix[0]),0,len(matrix) * len(matrix[0]) - 1
	for start <= end {
		mid := (start + end + 1) / 2
		if v := matrix[mid / n][mid % n];v == target {
			return true
		} else if v > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return false
}