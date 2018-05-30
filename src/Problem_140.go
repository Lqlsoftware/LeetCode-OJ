/*
Word Break II
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words,
add spaces in s to construct a sentence where each word is a valid dictionary word.
Return all such possible sentences.

Note:
The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.

Example 1:
Input:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
Output:
[
  "cats and dog",
  "cat sand dog"
]

Example 2:
Input:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
Output:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
Explanation: Note that you are allowed to reuse a dictionary word.

Example 3:
Input:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
Output:
[]
*/

package main

// dp + queue + dfs生成
func wordBreak(s string, wordDict []string) []string {
	if len(wordDict) == 0 {
		return []string{}
	}
	dict := make(map[string]bool, len(wordDict))
	min, max := len(wordDict[0]), len(wordDict[0])

	// hash map
	for _, v := range wordDict {
		if len(v) > max {
			max = len(v)
		} else if len(v) < min {
			min = len(v)
		}
		dict[v] = true
	}

	// dp or tree
	dp := make([][]string, len(s) + 1)
	for i := range dp {
		dp[i] = []string{}
	}
	dp[0] = append(dp[0], "")

	// queue
	position := list.New()
	position.PushBack(0)
	for position.Len() != 0 {
		pos := position.Remove(position.Front()).(int)
		start, end := min + pos, max + pos
		if end > len(s) {
			end = len(s)
		}
		for j := start;j <= end;j++ {
			v := s[pos:j]
			if _,ok := dict[v];ok {
				if len(dp[j]) == 0 {
					position.PushBack(j)
				}
				dp[j] = append(dp[j], v)
			}
		}
	}
	return generate(dp, len(s))
}

// dfs generate result from dp
func generate(dp [][]string, curr int) []string {
	if curr == 0 {
		return []string{}
	}
	var res []string
	for _,v := range dp[curr] {
		temp := generate(dp, curr - len(v))
		if len(temp) == 0 {
			temp = append(temp, v)
		} else {
			for i := range temp {
				temp[i] = temp[i] + " " +  v
			}
		}
		res = append(res, temp...)
	}
	return res
}