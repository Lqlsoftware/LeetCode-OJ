/*
Substring with Concatenation of All Words

You are given a string, s, and a list of words, words, that are all of the same length.
Find all starting indices of substring(s) in s that is a concatenation of each word in words exactly once and without any intervening characters.

For example, given:
s: "barfoothefoobarman"
words: ["foo", "bar"]

You should return the indices: [0,9].
(order does not matter).
*/

package main

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(s) == 0 || len(s) < len(words[0]) * len(words) {
		return []int{}
	}
	lenWord := len(words[0])
	lenWords := len(words)
	length := lenWord * lenWords
	total := make(map[string]int)
	res := make([]int,0)
	for _,word := range words {
		total[word]++
	}
	for idx := 0;idx < len(s) - length + 1;idx++ {
		current := make(map[string]int)
		j := 0
		for j = 0;j < length;j+=lenWord {
			substr := s[idx + j:idx + j + lenWord]
			if _,ok := total[substr];ok {
				current[substr]++
				if current[substr] > total[substr] {
					break
				}
			} else {
				break
			}
		}
		if j == length {
			res = append(res, idx)
		}
	}
	return res
}