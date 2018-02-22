/*

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