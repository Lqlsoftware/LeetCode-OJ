/*
Permutations
Given a collection of distinct numbers, return all possible permutations.

For example,
[1,2,3] have the following permutations:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

package main

// DFS
func permute(nums []int) [][]int {
	var result [][]int
	x := make([]bool, len(nums))
	solvePermute(&result, nums, []int{}, x)
	return result
}

func solvePermute(result *[][]int, nums, current []int, used []bool) {
	if len(current) == len(nums) {
		*result = append(*result, current)
		return
	}
	lenth := len(current)
	for i,v := range used {
		if !v {
			used[i] = true
			solvePermute(result, nums, append(current, nums[i]), used)
			current = current[:lenth]
			used[i] = false
		}
	}
}

// 使用交换完成DFS
func permute1(nums []int) [][]int {
	var result [][]int
	solvePermute1(&result, nums, 0)
	return result
}

func solvePermute1(result *[][]int, nums []int, idx int) {
	nextIdx := idx + 1
	if nextIdx == len(nums) {
		temp := make([]int, len(nums))
		copy(temp,nums)
		*result = append(*result, temp)
		return
	}
	for i := idx;i < len(nums);i++ {
		nums[idx],nums[i] = nums[i],nums[idx]
		solvePermute1(result, nums, nextIdx)
		nums[idx],nums[i] = nums[i],nums[idx]
	}
}