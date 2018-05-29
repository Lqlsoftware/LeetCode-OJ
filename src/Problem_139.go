/*
Word Break
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
determine if s can be segmented into a space-separated sequence of one or more dictionary words.
Note:
The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:
Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".

Example 2:
Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
Note that you are allowed to reuse a dictionary word.

Example 3:
Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
*/

package main

// dp solution
// dp[end] = true when dp[start] && in dict
func wordBreak(s string, wordDict []string) bool {
	dp, dict := make([]bool, len(s)+1), make(map[string]bool, len(wordDict))
	lens, max, min := len(s), 0, math.MaxInt32
	dp[0] = true
	for _, v := range wordDict {
		if len(v) < min {
			min = len(v)
		}
		if len(v) > max {
			max = len(v)
		}
		dict[v] = true
	}
	for start, length := 0, lens-min; start <= length; start++ {
		if max > lens {
			max = lens
		}
		if dp[start] {
			for end := min; end <= max; end++ {
				if _, ok := dict[s[start:end]]; ok {
					dp[end] = true
				}
			}
		}
		min++
		max++
	}
	return dp[len(s)]
}

// regexp is too great!!!
func wordBreakRegexp(s string, wordDict []string) bool {
	res, _ := regexp.MatchString("^("+strings.Join(wordDict, "|")+")*$", s)
	return res
}