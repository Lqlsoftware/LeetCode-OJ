/*
Interleaving String
Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.

For example,
Given:
s1 = "aabcc",
s2 = "dbbca",
When s3 = "aadbbcbcac", return true.
When s3 = "aadbbbaccc", return false.
*/

package main

// dp 二维dp化简
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	dp := make([]bool, len(s1) + 1)
	dp[0] = true
	for j := 0;j <= len(s2);j++ {
		for i := range dp {
			dp[i] = i > 0 && dp[i - 1] && s1[i - 1] == s3[i + j - 1] || j > 0 && dp[i] && s2[j - 1] == s3[i + j - 1] || i == 0 && j == 0
		}
	}
	return dp[len(s1)]
}

// 二维dp
func isInterleave1(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
		return false
	}
	dp := make([][]bool, len(s1) + 1)
	for i := range dp {
		dp[i] = make([]bool,len(s2) + 1)
	}
	dp[0][0] = true
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = i > 0 && dp[i - 1][j] && s1[i - 1] == s3[i + j - 1] || j > 0 && dp[i][j - 1] && s2[j - 1] == s3[i + j - 1] || dp[i][j]
		}
	}
	return dp[len(s1)][len(s2)]
}

// 递归处理
func isInterleave2(s1 string, s2 string, s3 string) bool {
	if len(s1) > len(s2) {
		return isInterleave(s2,s1,s3)
	} else if len(s1) == 0 {
		if s2 == s3 {
			return true
		} else {
			return false
		}
	}
	if s1[0] == s3[0] && isInterleave(s1[1:],s2,s3[1:]) {
		return true
	} else if s2[0] == s3[0] && isInterleave(s1,s2[1:],s3[1:]) {
		return true
	}
	return false
}