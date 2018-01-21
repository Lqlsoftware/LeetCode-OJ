/*
Letter Combinations of a Phone Number
Given a digit string, return all possible letter combinations that the number could represent.
A mapping of digit to letters (just like on the telephone buttons) is given below.

Input:Digit string "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:
Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	dict := make(map[string][]string)
	dict["2"] = []string{"a", "b", "c"}
	dict["3"] = []string{"d", "e", "f"}
	dict["4"] = []string{"g", "h", "i"}
	dict["5"] = []string{"j", "k", "l"}
	dict["6"] = []string{"m", "n", "o"}
	dict["7"] = []string{"p", "q", "r", "s"}
	dict["8"] = []string{"t", "u", "v"}
	dict["9"] = []string{"w", "x", "y", "z"}
	res := []string{""}
	for _, v := range digits {
		res2 := make([]string, 0)
		for _, v2 := range dict[string(v)] {
			for _, v3 := range res {
				res2 = append(res2, v3+v2)
			}
		}
		res = res2
	}
	return res
}
