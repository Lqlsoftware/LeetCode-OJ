/*
Distinct Subsequences
Given a string S and a string T, count the number of distinct subsequences of S which equals T.
A subsequence of a string is a new string which is formed from the original string by deleting some (can be none) of the characters without disturbing the relative positions of the remaining characters.
(ie, "ACE" is a subsequence of "ABCDE" while "AEC" is not).

Here is an example:
S = "rabbbit", T = "rabbit"
Return 3.
*/

package main

/*
s[i] == t[j] -> dp[i + 1][j + 1] = dp[i][j] + dp[i][j + 1]
s[i] != t[j] -> dp[i + 1][j + 1] = dp[i][j + 1]
	 _   r   a   b   b   i   t
    ---------------------------
 _ | 1 | 0 | 0 | 0 | 0 | 0 | 0 |
    ---------------------------
 r | 1 | 1 | 0 | 0 | 0 | 0 | 0 |
    ---------------------------
 a |   | 1 | 1 | 0 | 0 | 0 | 0 |
    ---------------------------
 b |       | 1 | 1 | 0 | 0 | 0 |
    ---------------------------
 b |           | 2 | 1 | 0 | 0 |
    ---------------------------
 b |               | 3 | 0 | 0 |
    ---------------------------
 i |                   | 3 | 0 |
    ---------------------------
 t |                       | 3 |
	---------------------------

simplify :
s[i] == t[j] -> dp[j + 1] = dp[j] + dp[j + 1] (j from tail to head)
s[i] != t[j] -> dp[j + 1] = dp[j + 1]
*/
func numDistinct(s string, t string) int {
	dp := make([]int, len(t) + 1)
	dp[0] = 1
	start := len(t) - len(s)
	for i := 0;i < len(s);i++ {
		for j := len(t);j > 0 && j > start;j-- {
			if s[i] == t[j - 1] {
				dp[j] = dp[j - 1] + dp[j]
			}
		}
		start++
	}
	return dp[len(t)]
}