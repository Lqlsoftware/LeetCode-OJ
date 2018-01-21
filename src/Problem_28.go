/*
Implement strStr()
Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:
Input: haystack = "hello", needle = "ll"
Output: 2

Example 2:
Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/

package main

func strStr(haystack string, needle string) int {
	lenH,lenN := len(haystack),len(needle)
	if lenN == 0 {
		return 0
	}
	for i,j := 0,0;i <= lenH - lenN;i++ {
		j = 0
		for haystack[i + j] == needle[j] {
			j++
			if j == lenN {
				return i
			}
		}
	}
	return -1
}