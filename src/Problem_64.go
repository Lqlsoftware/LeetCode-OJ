/*
Minimum Path Sum
Given a m x n grid filled with non-negative numbers,
find a path from top left to bottom right which minimizes the sum of all numbers along its path.
Note: You can only move either down or right at any point in time.

Example 1:
[[1,3,1],
 [1,5,1],
 [4,2,1]]
Given the above grid map, return 7. Because the path 1→3→1→1→1 minimizes the sum.
*/

package main

func minPathSum(grid [][]int) int {
	width,height := len(grid[0]),len(grid)
	dp := make([]int,width + 1)
	for i := 0;i < width;i++ {
		dp[i + 1] = dp[i] + grid[0][i]
	}
	for i := 1;i < height;i++ {
		dp[1] += grid[i][0]
		for j := 2;j <= width;j++ {
			dp[j] = min(dp[j],dp[j - 1]) + grid[i][j - 1]
		}
	}
	return dp[width]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}