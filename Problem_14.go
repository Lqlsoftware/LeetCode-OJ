/*
Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.
*/

package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs == nil {
		return ""
	}

	// 寻找最短字符串
	min := 0
	for i, v := range strs {
		if len(v) < len(strs[min]) {
			min = i
		}
	}
	res := []rune(strs[min])

	length := len(res)
	if length == 0 || strs[0] == "" {
		return ""
	}

	// find Common Prefix
	for _, v := range strs {
		runes := []rune(v)
		count := 0
		for i := 0; i < length; i++ {
			v := res[i]
			if v == runes[count] {
				count++
			} else {
				length = count
			}
		}
	}
	return string(res[:length])
}
