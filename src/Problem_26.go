/*
Remove Duplicates from Sorted Array
Reverse Nodes in k-Group
Given a sorted array,
remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array,
you must do this by modifying the input array in-place with O(1) extra memory.

Example:
Given nums = [1,1,2],
Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.
It doesn't matter what you leave beyond the new length.
*/

package main

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	pos := 0
	for i := 0;i < len(nums);i++ {
		if nums[i] != nums[pos] {
			pos++
			nums[pos] = nums[i]
		}
	}
	return pos + 1
}