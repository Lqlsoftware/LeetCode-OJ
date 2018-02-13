/*
Set Matrix Zeroes
Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in place.

Follow up:
Did you use extra space?
A straight forward solution using O(mn) space is probably a bad idea.
A simple improvement uses O(m + n) space, but still not the best solution.
Could you devise a constant space solution?
*/

package main

// store in col[0] and row[0]
func setZeroes(matrix [][]int)  {
	m,n,x := len(matrix),len(matrix[0]),false
	for i := range matrix {
		if matrix[i][0] == 0 {
			x = true
		}
		for j := 1;j < n;j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := m - 1;i >= 0;i-- {
		for j := n - 1;j >= 1;j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if x {
			matrix[i][0] = 0
		}
	}
}