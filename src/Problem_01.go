/*
Two Sum
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:
Given nums = [2, 7, 11, 15], target = 9,
Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

package main

func twoSum(nums []int, target int) []int {
	// 定义map存储target - num的值
	hashmap := make(map[int]int, len(nums))
	for i, val := range nums {
		if res, err := hashmap[val]; err {
			return []int{res, i}
		}
		// 未找到合适的 将其加入map
		hashmap[target-val] = i
	}
	return nil
}
