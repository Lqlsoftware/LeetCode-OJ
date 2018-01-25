/*
Search in Rotated Sorted Array

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
You are given a target value to search.
If found in the array return its index, otherwise return -1.
You may assume no duplicate exists in the array.
*/

package main

func search(nums []int, target int) int {
	if len(nums) < 1 {
		return -1
	}
	left,right,mid := 0, len(nums) - 1,0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		// 一定要跟右边比较，取mid时会取较小的值(删掉小数)作为mid值
		else if nums[mid] > nums[right] {
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}