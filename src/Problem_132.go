/*
Palindrome Partitioning II
Given a string s, partition s such that every substring of the partition is a palindrome.
Return the minimum cuts needed for a palindrome partitioning of s.

Example:
Input: "aab"
Output: 1
Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.
*/

package main

// dp
// 选定中心点 i
// 查找第i个字符为中心最大的回文串
//  	j 是查找宽度
// 		dp[i + j + 1] = min(dp[i + j + 1], 1 + dp[i - j])
func minCut(s string) int {
	dp := make([]int, len(s) + 1)
	for i := range dp {
		dp[i] = i - 1
	}
	for i := 1;i < len(s);i++ {
		for j := 0;i - j >= 0 && i + j < len(s) && s[i - j] == s[i + j];j++ {
			dp[i + j + 1] = min(dp[i + j + 1], 1 + dp[i - j])
		}
		for j := 0;i - j - 1 >= 0 && i + j < len(s) && s[i - j - 1] == s[i + j];j++ {
			dp[i + j + 1] = min(dp[i + j + 1], 1 + dp[i - j - 1])
		}
	}
	return dp[len(s)]
}

// dp[i]是到第i个字符前所需到最小切割数
// end   :  0 -> n
// start :	end -> 0
// 如果 s[end:start] 是回文 那么就是 dp[start] + 1 次切割
// 		dp[end] = min(dp[end], dp[start] + 1)
func minCut1(s string) int {
	dp := make([]int, len(s) + 1)
	for i := range dp {
		dp[i] = i - 1
	}
	isPalin := make([][]bool, len(s))
	for i := range isPalin {
		isPalin[i] = make([]bool, len(s))
	}
	for end := 1;end < len(s);end++ {
		for start := end;start >= 0;start-- {
			// 两个字符相等 距离小于2(aa || aba) || 字符中间是回文
			if s[start] == s[end] && (end - start < 2 || isPalin[start + 1][end - 1]) {
				isPalin[start][end] = true
				dp[end + 1] = min(dp[end + 1], 1 + dp[start])
			}
		}
	}
	return dp[len(s)]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}