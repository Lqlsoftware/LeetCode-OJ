/*
Decode Ways
A message containing letters from A-Z is being encoded to numbers using the following mapping:
'A' -> 1
'B' -> 2
...
'Z' -> 26
Given an encoded message containing digits, determine the total number of ways to decode it.

For example,
Given encoded message "12", it could be decoded as "AB" (1 2) or "L" (12).
The number of ways decoding "12" is 2.
*/

package main

// 动态规划 a,b,c = dp[i],dp[i + 1],dp[i + 2]
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	a,b,c := 0, 0, 1
	if s[len(s) - 1] != '0' {
		a,b = 1,1
	}
	for i := len(s) - 2;i >= 0;i--  {
		if s[i] < '1' || s[i] > '9' || (s[i + 1] == '0' && s[i] > '2') {
			a = 0
		} else if s[i] > '2' || (s[i] == '2' && s[i + 1] > '6') {
			a = b
		} else {
			a = b + c
		}
		b, c = a, b
	}
	return a
}