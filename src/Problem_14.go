/*
Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.
*/

package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i,v := range strs[0] {
		for j := 1;j < len(strs);j++ {
			if i == len(strs[j]) || strs[j][i] != uint8(v) {
				return strs[0][0:i]
			}
		}
	}
	return strs[0]
}