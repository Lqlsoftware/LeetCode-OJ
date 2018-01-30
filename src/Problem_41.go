/*
First Missing Positive
Given an unsorted integer array, find the first missing positive integer.

For example,
Given [1,2,0] return 3,
and [3,4,-1,1] return 2.

Your algorithm should run in O(n) time and uses constant space.
*/

package main

func firstMissingPositive(nums []int) int {
	lenth,idx := len(nums),0
	for idx < lenth {
		if v := nums[idx];v > 0 && v < lenth && v != nums[v - 1] {
			nums[idx],nums[v - 1] = nums[v - 1],v
		} else {
			idx++
		}
	}
	for idx = 1;idx <= lenth && idx == nums[idx - 1];idx++ {}
	return idx
}