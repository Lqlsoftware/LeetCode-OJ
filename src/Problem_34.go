/*
Search for a Range

Given an array of integers sorted in ascending order,
find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).
If the target is not found in the array, return [-1, -1].

For example,
Given [5, 7, 7, 8, 8, 10] and target value 8,
return [3, 4].
*/

package main

// binary search
func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1,-1}
	}

	left,right := 0,len(nums) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			start, end := mid,mid
			for start > left {
				center := (left + start) / 2
				if nums[center] == target {
					start = center
				} else {
					left = center + 1
				}
			}
			for end < right {
				center := (end + right + 1) / 2
				if nums[center] == target {
					end = center
				} else {
					right = center - 1
				}
			}
			return []int{start,end}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return []int{-1,-1}
}