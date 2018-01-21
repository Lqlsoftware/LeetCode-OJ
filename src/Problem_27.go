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

func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	pos,i := len(nums) - 1,0
	for i <= pos {
		if nums[i] == val {
			nums[i] = nums[pos]
			pos--
		} else {
			i++
		}
	}
	return i
}

/*
func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) < 1 {
		return 0
	}
	pos,i := len(nums) - 1,0
	for i <= pos {
		if nums[i] == val {
			for nums[pos] == val {
				pos--
				if pos < 0 {
					return 0
				} else if pos < i {
					return i
				}
			}
			nums[i] = nums[pos]
			pos--
		}
		i++
	}
	return i
}
*/