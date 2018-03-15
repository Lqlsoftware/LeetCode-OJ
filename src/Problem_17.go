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
	dict,res := []string{"abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"},[]string{""}
	for _,v := range digits {
		temp := make([]string,0)
		for _,v1 := range dict[v - '2'] {
			v2 := string(v1)
			for _,v3 := range res {
				temp = append(temp, v3 + v2)
			}
		}
		res = temp
	}
	return res
}