/*
Edit Distance
Given two words word1 and word2, find the minimum number of steps required to convert word1 to word2.
(each operation is counted as 1 step.)

You have the following 3 operations permitted on a word:
a) Insert a character
b) Delete a character
c) Replace a character
*/

package main

// dp
func minDistance(word1 string, word2 string) int {
	dp := make([][]int,len(word2) + 1)
	for i := range dp {
		dp[i] = make([]int,len(word1) + 1)
		dp[i][0] = i
	}
	for i := range dp[0] {
		dp[0][i] = i
	}
	for i,x := range word2 {
		for j,y := range word1 {
			if x == y {
				dp[i + 1][j + 1] = dp[i][j]
			} else {
				dp[i + 1][j + 1] = min(min(dp[i][j + 1],dp[i + 1][j]), dp[i][j]) + 1
			}
		}
	}
	return dp[len(word2)][len(word1)]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// optimized memory
func minDistance1(word1 string, word2 string) int {
	dp := make([]int,len(word2) + 1)
	for i := range dp {
		dp[i] = i
	}
	for i,x := range word1 {
		upLeft := dp[0]
		dp[0] = i + 1
		for j,y := range word2 {
			temp := dp[j + 1]
			if x == y {
				dp[j + 1] = upLeft
			} else {
				dp[j + 1] = min(min(dp[j],dp[j + 1]), upLeft) + 1
			}
			upLeft = temp
		}
	}
	return dp[len(word2)]
}