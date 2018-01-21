/*
Longest Substring Without Repeating Characters
Given a string, find the length of the longest substring without repeating characters.

Examples:
Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

package main

func lengthOfLongestSubstring(s string) int {
	// 定义字符数组存储每个字符的出现个数
	character := [128]int{}
	max, count, last := 0, 0, 0
	for index, ch := range s {
		// 如果该字符最近出现的位置在last之前 最长字串长度++
		if character[ch] < last {
			count++
		} else {
			// 记录新的起始位置 重复的字母最近出现的位置之后
			last = character[ch]
			if count > max {
				max = count
			}
			count = index + 1 - last
		}
		// 记录该字符最近出现的位置
		character[ch] = index + 1
	}
	if count > max {
		max = count
	}
	return max
}
