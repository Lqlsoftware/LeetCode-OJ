/*
Minimum Window Substring
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

For example,
S = "ADOBECODEBANC"
T = "ABC"
Minimum window is "BANC".

Note:
If there is no such window in S that covers all characters in T, return the empty string "".
If there are multiple such windows, you are guaranteed that there will always be only one unique minimum window in S.
*/

package main

func minWindow(s string, t string) string {
	dict := make([]int,128)
	for _,v := range t {
		dict[v]++
	}
	counter,begin,end,window,head := len(t), 0, 0, math.MaxInt32, 0
	for end < len(s) {
		if dict[s[end]] > 0 {
			counter--
		}
		dict[s[end]]--
		end++
		for counter == 0 { //valid
			if end - begin < window {
				window = end - begin
				head = begin
			}
			dict[s[begin]]++
			if dict[s[begin]] > 0 {
				counter++ //make it invalid
			}
			begin++
		}
	}
    if window == math.MaxInt32 {
		return ""
	}
	return s[head:head + window]
}