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
	character := [128]int{}
	max, count, last := 0, 0, 0
	for index, ch := range s {
		if character[ch] < last {
			count++
		} else {
			last = character[ch]
			if count > max {
				max = count
			}
			count = index + 1 - last
		}
		character[ch] = index + 1
	}
	if count > max {
		max = count
	}
	return max
}
