/*
Group Anagrams
Given an array of strings, group anagrams together.

For example, given: ["eat", "tea", "tan", "ate", "nat", "bat"],
Return:
[
  ["ate", "eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note: All inputs will be in lower-case.
*/

package main

func groupAnagrams(strs []string) [][]string {
	var result [][]string
	dict := make(map[string]int)
	for _,v := range strs {
		char := make([]byte,26)
		for _,val := range v {
			char[val - 'a']++
		}
		z := string(char)
		if x,ok := dict[z];ok == false {
			dict[z] = len(result)
			result = append(result, []string{v})
		} else {
			result[x] = append(result[x], v)
		}
	}
	return result
}