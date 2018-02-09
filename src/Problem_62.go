/*
Unique Paths
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).
The robot can only move either down or right at any point in time. The robot is trying to reach the bottom-right corner of the grid.
(marked 'Finish' in the diagram below)
How many possible unique paths are there?

Above is a 3 x 7 grid. How many possible unique paths are there?
Note: m and n will be at most 100.
*/

package main

// dp!
func uniquePaths(m int, n int) int {
	dp := make([]int,m + 1)
	dp[1] = 1
	for i := 0;i < m + n - 2;i++ {
		for j := m;j > 1;j-- {
			dp[j] = dp[j - 1] + dp[j]
		}
	}
	return dp[m]
}

// math!
func uniquePaths1(m int, n int) int {
	if m < n {
		return uniquePaths1(n,m)
	} else if n == 1 {
		return 1
	}
	res := m + n - 2
	for i := res - 1;i >= m;i-- {
		res *= i
	}
	for i := 2;i < n;i++ {
		res /= i
	}
	return res
}