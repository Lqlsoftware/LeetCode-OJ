/*
Sort Colors5
Given an array with n objects colored red, white or blue,
sort them so that objects of the same color are adjacent, with the colors in the order red, white and blue.
Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note:
You are not suppose to use the library's sort function for this problem.

Follow up:
A rather straight forward solution is a two-pass algorithm using counting sort.
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
Could you come up with an one-pass algorithm using only constant space?
*/

package main

// swap
func sortColors(nums []int) {
	i,j := 0, len(nums) - 1
	k,l := i,j
	for i <= j {
		if nums[i] == 0 {
			nums[i],nums[k] = nums[k],nums[i]
			i++
			k++
		} else if nums[i] == 2 {
			nums[i],nums[l] = nums[l],nums[i]
			l--
		} else {
			i++
		}
		if nums[j] == 2 {
			nums[j],nums[l] = nums[l],nums[j]
			j--
			l--
		} else if nums[j] == 0 {
			nums[k],nums[j] = nums[j],nums[k]
			k++
		} else {
			j--
		}
	}
}

// Lomuto partitioning : write 1 and overwrite with 0
func sortColors1(nums []int) {
	k,j := 0,0
	for i := 0;i < len(nums);i++ {
		v := nums[i]
		nums[i] = 2
		if v < 2 {
			nums[j] = 1
			j++
		}
		if v == 0 {
			nums[k] = 0
			k++
		}
	}
}