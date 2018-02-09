/*
Unique Paths II
Follow up for "Unique Paths":
Now consider if some obstacles are added to the grids. How many unique paths would there be?
An obstacle and empty space is marked as 1 and 0 respectively in the grid.

For example,
There is one obstacle in the middle of a 3x3 grid as illustrated below.
[
  [0,0,0],
  [0,1,0],
  [0,0,0]
]
The total number of unique paths is 2.
Note: m and n will be at most 100.
*/

package main

ffunc uniquePathsWithObstacles(obstacleGrid [][]int) int {
	width,height := len(obstacleGrid[0]),len(obstacleGrid)
	dp := make([]int,height + 1)
	if obstacleGrid[0][0] == 0 {
		dp[1] = 1
	}
	for i := 2;i < width + height;i++ {
		for j := height;j > 0;j-- {
			if x := i - j;x >= 0 && x < width && obstacleGrid[j - 1][x] == 1 {
				dp[j] = 0
			} else if j > 1 {
				dp[j] = dp[j - 1] + dp[j]
			}
		}
	}
	return dp[height]
}

// dp[i] = dp[i - 1] + dp[i] 从左边和上边来
// 按步数分层进行动态规划
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	width,height := len(obstacleGrid[0]),len(obstacleGrid)
	dp := make([]int,height + 1)
	if obstacleGrid[0][0] == 0 {
		dp[1] = 1
	}
	for i := 2;i < width + height;i++ {
		for j := height;j > 0;j-- {
			if x := i - j;x >= 0 && x < width && obstacleGrid[j - 1][x] == 1 {
				dp[j] = 0
			} else if j > 1 {
				dp[j] = dp[j - 1] + dp[j]
			}
		}
	}
	return dp[height]
}

// dp[i] = dp[i] + dp[i - 1] 从左边和上边来
// 按行数分层进行动态规划
func uniquePathsWithObstacles1(obstacleGrid [][]int) int {
	width,height := len(obstacleGrid[0]),len(obstacleGrid)
	dp := make([]int,height + 1)
	dp[1] = 1
	for i := 0;i < width;i++ {
		for j := 0;j < height;j++ {
			if obstacleGrid[j][i] == 1 {
				dp[j + 1] = 0
			} else {
				dp[j + 1] = dp[j + 1] + dp[j]
			}
		}
	}
	return dp[height]
}