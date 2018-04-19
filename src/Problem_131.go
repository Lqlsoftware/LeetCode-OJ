/*
Palindrome Partitioning
Given a string s, partition s such that every substring of the partition is a palindrome.
Return all possible palindrome partitioning of s.

For example, given s = "aab",
Return
[
  ["aa","b"],
  ["a","a","b"]
]
*/

package main

// dfs  (store the break point)
func partition(s string) [][]string {
	var result [][]string
	dfs(&result, s, []int{0})
	return result
}

func dfs(result *[][]string, s string, current []int) {
	length,last := len(current),current[len(current) - 1]
	if last == len(s) {
		res := make([]string, 0,len(current))
		for i := 1;i < len(current);i++ {
			res = append(res, s[current[i - 1]:current[i]])
		}
		*result = append(*result, res)
		return
	}
	for i := last;i < len(s);i++ {
		if check(s, last, i) {
			current = append(current, i + 1)
			dfs(result, s, current)
			current = current[:length]
		}
	}
}

func check(s string, start, end int) bool {
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}