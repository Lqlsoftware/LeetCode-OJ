/*
Generate Parentheses
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

package main

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	// 第一个字符为 '('
	genParenthesis(n*2, 1, "(", &res)
	return res
}

// 回溯算法
func genParenthesis(n, left int, temp string, res *[]string) {
	// 到达最后一个字符 ')'
	if n == 2 {
		if left == 1 {
			*res = append(*res, temp+")")
		}
		return
	}
	// 当前位置左边的 '(' 个数不足
	if left < n-1 {
		genParenthesis(n-1, left+1, temp+"(", res)
	}
	// 当前位置左边的 '(' 个数足够
	if left > 0 {
		genParenthesis(n-1, left-1, temp+")", res)
	}
}
