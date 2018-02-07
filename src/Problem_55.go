/*
Jump Game
Given an array of non-negative integers, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Determine if you are able to reach the last index.

For example:
A = [2,3,1,1,4], return true.
A = [3,2,1,0,4], return false.
*/

package main

// 从后往前贪心
func canJump(nums []int) bool {
	current := len(nums) - 1
	for i := current;i >= 0;i-- {
		if i + nums[i] >= current {
			current = i
		}
	}
	return current == 0
}

// 从前往后贪心
func canJump1(nums []int) bool {
	length,current := len(nums) - 1,0
	for i := 0;i <= current;i++ {
		if current >= length {
			return true
		}
		current = Max(current,i + nums[i])
	}
	return false
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}