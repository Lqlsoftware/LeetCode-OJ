/*
Maximal Rectangle
Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

For example, given the following matrix:
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0
Return 6.
*/

package main

// trick with problem 84's solution
func maximalRectangle(matrix [][]byte) int {
    if len(matrix) == 0 {
        return 0
    }
	n := len(matrix[0])
	height := make([]int,n)
	maxS := 0
	for i := range matrix {
		for j,v := range matrix[i] {
			if v == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		maxS = max(maxS,largestRectangleArea(height))
	}
	return maxS
}

// dp with left[n] and right[n] 
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	n,maxS := len(matrix[0]),0
	height,left,right := make([]int,n),make([]int,n),make([]int,n)
	for i := range right {
		right[i] = n
	}
	for i := range matrix {
		curl,curr := 0,n
		// height at i,j
		for j,v := range matrix[i] {
			if v == '1' {
				height[j]++
			} else {
				height[j] = 0
			}
		}
		// left start at i,j's height
		for j,v := range matrix[i] {
			if v == '1' {
				left[j] = max(left[j],curl)
			} else {
				left[j] = 0
				curl = j + 1
			}
		}
		// right end at i,j's height
		for j := n - 1;j >= 0;j-- {
			if matrix[i][j] == '1' {
				right[j] = min(right[j],curr)
			} else {
				right[j] = n
				curr = j
			}
		}
		// calculate s
		for j := range matrix[i] {
			x := (right[j] - left[j]) * height[j]
			maxS = max(maxS,x)
		}
	}
	return maxS
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}