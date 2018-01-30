/*
Combination Sum

Given a set of candidate numbers (C) (without duplicates) and a target number (T),
find all unique combinations in C where the candidate numbers sums to T.
The same repeated number may be chosen from C unlimited number of times.

Note:
All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.

For example, given candidate set [2, 3, 6, 7] and target 7,
A solution set is:
[
  [7],
  [2, 2, 3]
]
*/

package main

// 回溯法
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	currents := make([]int, 0)
	sort.Ints(candidates)
	solveSum(&result, candidates,currents,target)
	return result
}

func solveSum(result *[][]int,candidates, currents []int, target int) {
	if target < 0 {
		return
	} else if target == 0 {
		temp := make([]int, len(currents))
		copy(temp, currents)
		*result = append(*result, temp)
		return
	}
	for i := 0; i < len(candidates) && target >= candidates[i]; i++ {
		l := len(currents)
		currents = append(currents, candidates[i])
		solveSum(result, candidates[i:], currents, target - candidates[i])
		currents = currents[:l]
	}
}

// 动态规划
func combinationSum1(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	if candidates[0] > target {
		return [][]int{}
	}
	dp := make([][][]int,target + 1)
	for i := 0;i < len(candidates) && candidates[i] <= target;i++ {
		dp[candidates[i]] = [][]int{{candidates[i]}}
	}
	for i := candidates[0] + 1;i<=target;i++ {
		for _,v := range candidates {
			if i < v {
				continue
			}
			for _,val := range dp[i - v] {
				temp := append(val,v)
				sort.Ints(temp)
				if !contain(dp[i],temp) {
					tt := make([]int, len(temp))
					copy(tt, temp)
					dp[i] = append(dp[i], tt)
				}
			}
		}
	}
	return dp[target]
}

func equal(s1, s2 []int) bool {
	if len(s1)!=len(s2) {
		return false
	}
	for i,x := range s1 {
		if x != s2[i] {
			return false
		}
	}
	return true
}

func contain(src [][]int, target []int) bool {
	if len(src) == 0 {
		return false
	}
	for _,v := range src {
		if equal(v,target) {
			return true
		}
	}
	return false
}

// 递归
func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	for idx,v := range candidates {
		if diff := target - v;diff > 0 {
			for _,x := range combinationSum2(candidates[idx:],diff) {
				temp := make([]int,len(x) + 1)
				copy(temp,x)
				temp[len(x)] = v
				res = append(res, temp)
			}
		} else if diff == 0 {
			res = append(res, []int{v})
		}
	}
	return res
}