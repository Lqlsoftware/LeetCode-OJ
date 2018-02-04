/*
Permutations II
Given a collection of numbers that might contain duplicates, return all possible unique permutations.

For example,
[1,1,2] have the following unique permutations:
[
  [1,1,2],
  [1,2,1],
  [2,1,1]
]
*/

package main

// DFS
func permuteUnique(nums []int) [][]int {
	var result [][]int
	x := make([]bool, len(nums))
	sort.Ints(nums)
	solvePermute(&result, nums, []int{}, x)
	return result
}

func solvePermute(result *[][]int, nums, current []int, used []bool) {
	if len(current) == len(nums) {
		temp := make([]int,len(current))
		copy(temp,current)
		*result = append(*result, temp)
		return
	}
	length,last := len(current),math.MaxInt32
	for i,v := range nums {
		if !used[i] && v != last {
			last = v
			used[i] = true
			solvePermute(result, nums, append(current, v), used)
			current = current[:length]
			used[i] = false
		}
	}
}