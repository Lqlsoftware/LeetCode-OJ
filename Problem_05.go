/*
Longest Palindromic Substring
Given a string s, find the longest palindromic substring in s.
You may assume that the maximum length of s is 1000.

Example:
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example:
Input: "cbbd"
Output: "bb"
*/

package main

func longestPalindrome(s string) string {
	var ress, rese, i, j, lenth, max int
	for index := range s {
		if index > 1 && s[index] == s[index-2] {
			lenth = 3
			i, j = index-2, index
			for i >= 0 && j < len(s) && s[j] == s[i] {
				lenth += 2
				j++
				i--
			}
			if lenth > max {
				max, ress, rese = lenth, i+1, j-1
			}
		}
		if index > 0 && s[index] == s[index-1] {
			lenth = 2
			i, j = index-1, index
			for i >= 0 && j < len(s) && s[j] == s[i] {
				lenth += 2
				j++
				i--
			}
			if lenth > max {
				max, ress, rese = lenth, i+1, j-1
			}
		}
	}
	if max == 0 {
		return s[0:1]
	}
	return s[ress : rese+1]
}
