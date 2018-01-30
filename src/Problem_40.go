/*
Combination Sum II
Given a collection of candidate numbers (C) and a target number (T),
find all unique combinations in C where the candidate numbers sums to T.
Each number in C may only be used once in the combination.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.

For example, given candidate set [10, 1, 2, 7, 6, 1, 5] and target 8,
A solution set is:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
*/

package main

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var result [][]int
	solve(&result, candidates, &[]int{}, target)
	return result
}

// 回溯
func solve(result *[][]int, candidates []int, buf *[]int, target int) {
	if target == 0 {
		found := make([]int, len(*buf))
		copy(found, *buf)
		*result = append(*result, found)
		return
	}
	for idx,v := range candidates {
		if idx > 0 && candidates[idx] == candidates[idx - 1]  {
			continue
		} else if diff := target - v;diff >= 0 {
			*buf = append(*buf,v)
			solve(result, candidates[idx + 1:], buf, diff)
			*buf = (*buf)[:len(*buf) - 1]
		}
	}
}