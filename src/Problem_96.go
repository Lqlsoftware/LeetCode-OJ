/*
Unique Binary Search Trees
Given n, how many structurally unique BST's (binary search trees) that store values 1...n?

For example,
Given n = 3, there are a total of 5 unique BST's.

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3
*/

package main

// i 划分成两部分 长度 i - 1 和 n - i
// dp(n) = dp(i - 1) * dp(n - i) (1 <= i <= n)
func numTrees(n int) int {
	dp := make([]int,n + 1)
	dp[0],dp[1] = 1,1
	for l := 2;l <= n;l++ {
		half := l >> 1
		for i := 1;i <= half;i++ {
			dp[l] += dp[i - 1] * dp[l - i]
		}
		// 减少重复计算
		dp[l] <<= 1
		dp[l] += l & 1 * dp[half] * dp[half]
	}
	return dp[n]
}