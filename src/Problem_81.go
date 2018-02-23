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

// seems redundant but efficient(maybe not) (8ms)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	num,counter,k := make([]int,len(nums)),false,0
	for i := 1;i < len(nums);i++ {
		if nums[i] == nums[k] {
			if !counter {
				k++
				counter,nums[k] = true,nums[i]
			}
		} else {
			k++
			counter,nums[k] = false,nums[i]
		}
	}
	nums = num
	return k + 1
}

// (12ms)
func removeDuplicates1(nums []int) int {
	k := 0
	for _,v := range nums {
		if k < 2 || v != nums[k - 2] {
			nums[k] = v
			k++
		}
	}
	return k
}