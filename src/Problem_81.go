/*
Search in Rotated Sorted Array II
Follow up for "Problem 33: Search in Rotated Sorted Array":
What if duplicates are allowed?
Would this affect the run-time complexity? How and why?
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
(i.e., 0 1 2 4 5 6 7 might become 4 5 6 7 0 1 2).
Write a function to determine if a given target is in the array.
The array may contain duplicates.
*/

package main

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	left,right,mid := 0, len(nums) - 1,0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[right] {
			if nums[mid] > target && nums[left] <= target {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[right] {
			if nums[mid] < target && nums[right] >= target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			right--
		}
	}
	return false
}