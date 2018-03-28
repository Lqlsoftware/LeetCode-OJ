/*
Triangle
Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

Note:
Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

*/

package main

// from bottom to top
func minimumTotal(triangle [][]int) int {
	dp := make([]int,len(triangle))
	copy(dp, triangle[len(triangle) - 1])
	for i := len(triangle) - 2;i >= 0;i-- {
		for j := 0;j <= i;j++ {
			dp[j] = min(dp[j], dp[j + 1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}