/*
Subsets II
Given a collection of integers that might contain duplicates, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.

For example,
If nums = [1,2,2], a solution is:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]
*/

package main

// 如果发现数字和上一个数字相同,则只往上一个数字所增加的result里添加该数
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{{}}
	l := 0
	for i := range nums {
		// l为上个数所创建的result的长度
		if i == 0 || nums[i] != nums[i - 1] {
			// 这样就更新全部result
			l = len(result)
		}
		length := len(result)
		for j := length - l;j < length;j++ {
			temp := make([]int, len(result[j]), len(result[j]) + 1)
			copy(temp, result[j])
			result = append(result, append(temp, nums[i]))
		}
	}
	return result
}

// 思路同 Problem 78 Subsets 相当于如果发现数字和上一个数字相同,则只往上一个数字所增加的result里添加该数
func subsetsWithDup1(nums []int) [][]int {
	sort.Ints(nums)
	return solve(nums)
}

func solve(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	// 得到不包含第一个数字的子集
	count := 0
	for ;count < len(nums);count++ {
		if nums[count] != nums[0] {
			break
		}
	}
	result := solve(nums[count:])
	// 将第一个数字(以及它的2 ~ count个复制)添加到复制后到子集中
	for i := range result {
		temp,templen := []int{nums[0]},len(result[i])
		for j := 1;j <= count;j++ {
			target := make([]int,templen,templen + j)
			copy(target,result[i])
			target = append(target, temp...)
			result = append(result, target)
			temp = append(temp, nums[0])
		}
	}
	return result
}