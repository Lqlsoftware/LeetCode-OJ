/*
Longest Valid Parentheses
Given a string containing just the characters '(' and ')',
find the length of the longest valid (well-formed) parentheses substring.
For "(()", the longest valid parentheses substring is "()", which has length = 2.
Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.
*/

package main

func longestValidParentheses(s string) int {
	if len(s) < 2 || s == ")(" {
		return 0
	}
	dp := make([]int,len(s) + 1)
	max := 0
	for idx := 1;idx < len(s);idx++ {
		if s[idx] == ')' && s[idx - 1] == '(' {
			dp[idx + 1] = dp[idx - 1] + 2
		} else if s[idx] == ')' && s[idx - 1] == ')' {
			if idx - dp[idx] - 1 >= 0 && s[idx - dp[idx] - 1] == '(' {
				dp[idx + 1] = dp[idx] + dp[idx - dp[idx] - 1] + 2
			}
		}
		if dp[idx + 1] > max {
			max = dp[idx + 1]
		}
	}
	return max
}